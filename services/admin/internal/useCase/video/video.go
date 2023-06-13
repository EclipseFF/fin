package video

import (
	"architecture_go/pkg/type/context"
	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/admin/internal/domain/video"
	"github.com/opentracing/opentracing-go"
)

func (uc *UseCase) Delete(ctx context.Context, ID uint) error {
	return uc.adapterStorage.DeleteVideo(ctx, ID)
}

func (uc *UseCase) List(c context.Context, parameter queryParameter.QueryParameter) ([]*video.Video, error) {

	span, ctx := opentracing.StartSpanFromContext(c, "List")
	defer span.Finish()

	return uc.adapterStorage.ListVideos(context.New(ctx), parameter)
}
