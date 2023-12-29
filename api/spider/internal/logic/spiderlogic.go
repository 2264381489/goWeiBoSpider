package logic

import (
	"context"

	"goSpider/api/spider/internal/svc"
	"goSpider/api/spider/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SpiderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSpiderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SpiderLogic {
	return &SpiderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SpiderLogic) Spider(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
