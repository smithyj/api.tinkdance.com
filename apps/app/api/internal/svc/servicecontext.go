package svc

import (
	"context"
	"errors"
	"image/color"
	"tinkdance/pkg/captchax"
	"tinkdance/pkg/redisx"

	afocuscaptcha "github.com/afocus/captcha"

	"tinkdance/apps/app/api/internal/assets"
	"tinkdance/apps/app/api/internal/config"
)

const ServiceKey = "service-context"

type ServiceContext struct {
	Config  *config.Config
	Redis        redisx.Redis
	Captcha      captchax.Captcha
	ImageCaptcha *afocuscaptcha.Captcha
}

func NewServiceContext(config *config.Config) (svcCtx *ServiceContext, err error) {
	redis, err := redisx.New(config.Redis)
	if err != nil {
		return nil, err
	}

	imgCaptcha, err := newImgCaptcha()
	if err != nil {
		return nil, err
	}

	svcCtx = &ServiceContext{
		Config:       config,
		Redis:        redis,
		Captcha:      captchax.New(redis),
		ImageCaptcha: imgCaptcha,
	}
	return
}

func newImgCaptcha() (*afocuscaptcha.Captcha, error) {
	captcha := afocuscaptcha.New()
	buf, err := assets.FS.ReadFile("fonts/comic.ttf")
	if err != nil {
		return nil, err
	}
	if err := captcha.AddFontFromBytes(buf); err != nil {
		return nil, err
	}

	captcha.SetSize(200, 100)
	captcha.SetDisturbance(afocuscaptcha.HIGH)
	captcha.SetFrontColor(color.RGBA{255, 255, 255, 255})
	captcha.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	return captcha, nil
}

// ServiceWithContext 从上下文中获取 ServiceContext
func ServiceWithContext(ctx context.Context) (context.Context, *ServiceContext, error) {
	v, ok := ctx.Value(ServiceKey).(*ServiceContext)
	if !ok {
		return ctx, nil, errors.New("SvcCtx is not initialize")
	}
	return ctx, v, nil
}
