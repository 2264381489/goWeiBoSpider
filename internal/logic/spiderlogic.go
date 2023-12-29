package logic

import (
	"context"
	"goSpider/internal/svc"
	"goSpider/internal/types"

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

func (l *SpiderLogic) Spider(req *types.SpiderRequest) (resp *types.SpiderResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
