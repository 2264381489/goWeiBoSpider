package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/queue"
	"github.com/zeromicro/go-zero/core/logx"
	customerErr "goSpider/errors"
	"goSpider/model"
	"goSpider/spider/types"
	"goSpider/svc"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Spider struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSpider(ctx context.Context, svcCtx *svc.ServiceContext) *Spider {
	return &Spider{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (s *Spider) Run(uid int) {

	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	q, _ := queue.New(
		1, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c.OnRequest(func(request *colly.Request) {
		logx.Infof("calling %s", request.URL.String())
		request.Headers.Add("Referer", "https://m.weibo.cn/u/5680343342")
		request.Headers.Add("Accept", "application/json, text/plain, */*")
		request.Headers.Add("Dnt", fmt.Sprintf("%d", 1))
		request.Headers.Add("Mweibo-Pwa", fmt.Sprintf("%d", 1))
		request.Headers.Add("X-Requested-With", "XMLHttpRequest")
		request.Ctx.Put("uid", uid)
		request.Ctx.Put("containerId", genContainerId(uid, WeiBo))
	})

	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*https.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	c.OnError(func(response *colly.Response, err error) {
		if err != nil {
			panic(err)
		}
		marshal, err := json.Marshal(response)

		if err != nil {
			panic(err)
		}
		fmt.Println(string(marshal))
	})

	c.OnResponse(func(response *colly.Response) {
		resp := &types.GetIndexResp{}
		err := json.Unmarshal(response.Body, resp)
		if err != nil {
			panic(err)
		}

		if err != nil && err != os.ErrExist {
			panic(err)
		}

		err, sinceId := NewSpider(s.ctx, s.svcCtx).handleRespData(*resp)
		if err != nil {
			panic(err)
		}
		mainUrl := fmt.Sprintf(userWeiBoList, uid, genContainerId(uid, WeiBo))
		err = q.AddURL(fmt.Sprintf("%s&since_id=%d", mainUrl, sinceId))
		if err != nil {
			panic(err)
		}
	})

	err = q.AddURL(fmt.Sprintf(userWeiBoList, uid, genContainerId(uid, WeiBo)))
	if err != nil {
		panic(err)
	}
	err = q.Run(c)
	if err != nil {
		panic(err)
	}
}

func genContainerId(uid int, containerType ContainerType) string {
	switch containerType {
	case MainTheme:
		return fmt.Sprintf("%d%d", MainThemeCId, uid)
	case WeiBo:
		return fmt.Sprintf("%d%d", WeiBoCId, uid)
	}

	return ""
}

func (s *Spider) handleRespData(resp types.GetIndexResp) (error, int64) {
	// 获取照片
	if resp.Ok != 1 {
		return customerErr.ErrInvalidRespStatus, 0
	}
	sinceId := resp.Data.CardlistInfo.SinceID

	for _, card := range resp.Data.Cards {
		userId := card.Mblog.User.ID
		if card.CardType == Mblog.ToInt() {
			err := s.SaveMblog(card, userId)
			if err != nil {
				return err, 0
			}
			if card.Mblog.PageInfo.Type == VideoType.String() {
				err = s.SaveVideo(card)
				if err != nil {
					return err, 0
				}
			}
			err = s.SavePictures(card, err)
			if err != nil {
				return err, 0
			}
		}
	}
	return nil, sinceId
}

func (s *Spider) SavePictures(card types.Cards, err error) error {
	data := []*model.Pic{}
	for _, picInfo := range card.Mblog.Pics {
		picUrl := picInfo.URL
		pidLUrl := picInfo.Large.URL
		pid := picInfo.Pid
		pic := &model.Pic{
			PicId:    pid,
			Url:      picUrl,
			LargeUrl: pidLUrl,
		}
		data = append(data, pic)
	}

	if len(data) == 0 {
		logx.Infof(" card scheme :%s  no pic", card.Scheme)
		return nil
	}

	err = s.svcCtx.PicModel.InsertMany(s.ctx, data)
	if err != nil {
		return err
	}

	for _, pic := range data {
		time.Sleep(300)
		_, err := os.Stat("./script/download.sh")
		if err != nil {
			if !os.IsNotExist(err) {
				logx.Errorf("os.Stat(./switchkernel.sh):%v", err)
				return err
			}
		} else {
			if err := Exec("./script/download.sh", pic.LargeUrl, fmt.Sprintf("%s.jpg", pic.PicId)); err != nil {
				logx.Errorf("Exec(./script/download.sh):%v", err)
				return err
			}
		}
	}
	return nil
}

func (s *Spider) SaveVideo(card types.Cards) error {
	title := card.Mblog.PageInfo.Title
	videoURL := card.Mblog.PageInfo.MediaInfo.StreamURL     // 普通视频地址
	videoHDURL := card.Mblog.PageInfo.MediaInfo.StreamURLHd // 高清视频地址
	video720URL := card.Mblog.PageInfo.Urls.Mp4720PMp4      //720p 视频地址
	duration := card.Mblog.PageInfo.MediaInfo.Duration      // 视频时长
	err := s.svcCtx.VideoModel.Insert(s.ctx, &model.Video{
		Duration:    duration,
		VideoUrl:    videoURL,
		VideoHDUrl:  videoHDURL,
		Video720URL: video720URL,
		Title:       title,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Spider) SaveMblog(card types.Cards, userId int64) error {
	text := card.Mblog.Text
	mblogId := card.Mblog.ID
	mblogUrl := card.Scheme
	attitudesCount := card.Mblog.AttitudesCount // 点赞
	comentsCount := card.Mblog.CommentsCount    // 回复
	repostsCount := card.Mblog.RepostsCount     // 转发
	err := s.svcCtx.MblogModel.Insert(s.ctx, &model.Mblog{
		MblogId:        mblogId,
		UserId:         userId,
		Text:           text,
		AttitudesCount: attitudesCount,
		CommentsCount:  comentsCount,
		RepostsCount:   repostsCount,
		Url:            mblogUrl,
	})
	return err
}

func Exec(scriptPath string, args ...string) error {
	cmd := fmt.Sprintf("/bin/bash %s", scriptPath)
	cmdStr := strings.Replace(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), "/bin/bash -c ", "", -1)

	scriptFile, err := genTempScript([]byte(cmdStr))
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return
		}
		os.Remove(scriptFile)
	}()
	command := exec.Command("/bin/bash", scriptFile)
	output, err := command.CombinedOutput()

	if err != nil {
		errStr := string(output)
		if exitError, ok := err.(*exec.ExitError); ok {
			errStr = fmt.Sprintf("%s,exit %d", errStr, exitError.ExitCode())
			logx.Errorf(errStr)
		}
		logx.Errorf(err.Error())
	}
	return nil
}

func genTempScript(content []byte) (scriptFile string, err error) {
	err = os.Mkdir("./tmp", 0o744)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return "", err
	}
	scriptFile = filepath.Join("./tmp", fmt.Sprintf("%d-%d.sh", time.Now().UnixNano(), rand.Intn(10000)))
	if err = ioutil.WriteFile(scriptFile, content, 0o744); err != nil {
		return "", err
	}
	return scriptFile, nil
}
