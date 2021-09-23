package captcha

import "context"

type ImageRequest struct{}

type ImageResponse struct{}

func Image(ctx context.Context, req *ImageRequest) *ImageResponse {
	var s int16 = 12
	panic(&s)
	return &ImageResponse{}
}
