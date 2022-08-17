package job

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysJobLogic {
	return &UpdateSysJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysJobLogic) UpdateSysJob(req *types.UpdateSysJobReq) error {
	sysJob, err := l.svcCtx.SysJobModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = copier.Copy(sysJob, req)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = l.svcCtx.SysJobModel.Update(l.ctx, sysJob)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}