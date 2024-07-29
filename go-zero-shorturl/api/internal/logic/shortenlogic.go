package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"shortenurl/api/internal/svc"
	"shortenurl/api/internal/types"
	"shortenurl/rpc/transform/transform"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (resp *types.ShortenResq, err error) {
	rpcResp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transform.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return nil, err
	}

	return &types.ShortenResq{
		Shorten: rpcResp.Shorten,
	}, nil
}
