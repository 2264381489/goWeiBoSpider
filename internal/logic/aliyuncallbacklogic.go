package logic

import (
	"context"
	"fmt"
	"goSpider/errors"

	"goSpider/internal/svc"
	"goSpider/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AliyunCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

const (
	UserAccessTokenKey = "user_access_token:%s"
)

func NewAliyunCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AliyunCallbackLogic {
	return &AliyunCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AliyunCallbackLogic) AliyunCallback(req *types.AliYunCallbackReq) (resp *types.AliYunCallbackResp, err error) {
	if req.AccessToken == "" {
		return nil, errors.ErrInvalidAccessToken
	}
	resp = new(types.AliYunCallbackResp)
	redisKey := fmt.Sprintf(UserAccessTokenKey, req.AccessToken)
	err = l.svcCtx.BizRedis.Set(redisKey, req.AccessToken)
	if err != nil {
		return nil, err
	}
	resp.Code = 200
	resp.Message = "success"
	return
}
