package useCase

import (
	"architecture_go/pkg/type/context"
	"architecture_go/services/api/internal/domain/admin"
)

type Request interface {
	Create(c context.Context, requests ...*admin.Request) ([]*admin.Request, error)
}
