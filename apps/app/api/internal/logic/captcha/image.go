package captcha

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"github.com/afocus/captcha"
	uuid "github.com/satori/go.uuid"
	"image/png"
	"tinkdance/apps/app/api/internal/svc"
	"tinkdance/pkg/codex"
)

type ImageRequest struct{}

type ImageResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data ImageResponseData `json:"data"`
}

type ImageResponseData struct {
	CaptchaID string `json:"captcha_id"`
	Base64    string `json:"base64"`
}

func Image(ctx context.Context, req *ImageRequest) (*ImageResponse, error) {
	_, reqCtx, err := svc.RequestWithContext(ctx)
	if err != nil {
		return nil, err
	}

	img, code := reqCtx.SvcCtx.ImageCaptcha.Create(6, captcha.ALL)

	captchaId := uuid.NewV5(uuid.NewV1(), "captcha").String()

	buf := bytes.NewBuffer(nil)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	ok := reqCtx.SvcCtx.Captcha.Set(ctx, captchaId, code)
	if !ok {
		return nil, errors.New("验证码生成失败")
	}

	return &ImageResponse{
		Code: codex.Success,
		Msg:  codex.Msg(codex.Success),
		Data: ImageResponseData{
			CaptchaID: captchaId,
			Base64:    "data:image/jpeg;base64," + base64Str,
		},
	}, nil
}
