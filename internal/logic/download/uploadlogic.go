package download

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10 MB

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(r *http.Request, svcCtx *svc.ServiceContext) UploadLogic {
	return UploadLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload() (resp *types.UploadResponse, err error) {
	l.r.ParseMultipartForm(maxFileSize)
	// 在这里添加一句：files := l.r.MultipartForm.File["myFile"]
	files := l.r.MultipartForm.File["myFile"]
	// 用一个for循环来遍历files切片
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		logx.Infof("upload file: %+v, file size: %d, MIME header: %+v",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		tempFile, err := os.Create(path.Join(l.svcCtx.Config.Path, fileHeader.Filename))
		if err != nil {
			return nil, err
		}
		defer tempFile.Close()
		io.Copy(tempFile, file)
	}

	return &types.UploadResponse{
		Code: 0,
	}, nil
}
