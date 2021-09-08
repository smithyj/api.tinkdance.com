package svc

import (
	"image/color"
	"tinkdance/pkg/captcha"
	"tinkdance/pkg/redis"

	gocaptcha "github.com/afocus/captcha"

	"tinkdance/apps/app/api/internal/assets"
	"tinkdance/apps/app/api/internal/config"
)

type ServiceContext struct {
	Config       *config.Config
	Redis        redis.Redis
	Captcha      captcha.Captcha
	ImageCaptcha *gocaptcha.Captcha
}

func NewServiceContext(config *config.Config) (svcCtx *ServiceContext, err error) {
	redisConn, err := redis.New(config.Redis)
	if err != nil {
		return nil, err
	}

	imgCaptcha, err := newImgCaptcha()
	if err != nil {
		return nil, err
	}

	svcCtx = &ServiceContext{
		Config:       config,
		Redis:        redisConn,
		Captcha:      captcha.New(redisConn),
		ImageCaptcha: imgCaptcha,
	}
	return
}

func newImgCaptcha() (*gocaptcha.Captcha, error) {
	c := gocaptcha.New()
	buf, err := assets.FS.ReadFile("fonts/comic.ttf")
	if err != nil {
		return nil, err
	}
	if err := c.AddFontFromBytes(buf); err != nil {
		return nil, err
	}

	c.SetSize(200, 100)
	c.SetDisturbance(gocaptcha.HIGH)
	c.SetFrontColor(color.RGBA{255, 255, 255, 255})
	c.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	return c, nil
}
