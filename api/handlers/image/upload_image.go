package image

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	o "gamersnote.com/v1/restapi/operations/image"
	img "gamersnote.com/v1/utils/image"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

// NewUploadImageHandler UploadImageHandlerのコンストラクタ
func NewUploadImageHandler(s img.Service) *UploadImageHandler {
	return &UploadImageHandler{
		imgSvc: s,
	}
}

// UploadImageHandler 画像のアップロード
type UploadImageHandler struct {
	imgSvc img.Service
}

// Handle 画像をアップロードします。
func (h UploadImageHandler) Handle(params o.UploadImageParams) middleware.Responder {
	// FormDataをデコード
	img, f, err := image.Decode(params.Image)
	if err != nil {
		return o.NewUploadImageDefault(500)
	}
	// ファイル形式に合わせてエンコード
	var buf bytes.Buffer
	if f == "png" {
		err = png.Encode(&buf, img)
	}
	if f == "jpeg" {
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 30})
	}
	if err != nil {
		return o.NewUploadImageDefault(500)
	}
	// アップロード
	filename := uuid.New().String() + "." + f
	err = h.imgSvc.Upload(filename, bytes.NewReader(buf.Bytes()))
	if err != nil {
		return o.NewUploadImageDefault(500)
	}
	// 参照先urlを返す
	url := os.Getenv("S3_BASEURL") + "/" + filename
	return o.NewUploadImageCreated().WithPayload(&o.UploadImageCreatedBody{
		URL: &url,
	})
}
