package logic

import (
	"context"
	"fmt"

	"ginweb02/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type HelloWordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) HelloWordLogic {
	return HelloWordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloWordLogic) HelloWord() error {
	// todo: add your logic here and delete this line
	fmt.Println("hello")
	return nil
}
