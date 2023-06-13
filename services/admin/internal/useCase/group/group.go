package group

import (
	"architecture_go/pkg/type/context"
	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/admin/internal/domain/group"
)

func (uc *UseCase) Create(ctx context.Context, groupCreate *group.Group) (*group.Group, error) {
	return uc.adapterStorage.CreateGroup(ctx, groupCreate)
}

func (uc *UseCase) Delete(ctx context.Context, ID uint) error {
	return uc.adapterStorage.DeleteGroup(ctx, ID)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	return uc.adapterStorage.ListGroup(ctx, parameter)
}
