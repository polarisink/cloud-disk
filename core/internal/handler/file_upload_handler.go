package handler

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"crypto/md5"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"path"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// why write logic code here
		file, fileHeader, err := r.FormFile("file")
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash=?", hash).Get(rp)
		if err != nil {
			return
		}
		if has {
			//如果已经存在,就直接返回
			httpx.OkJson(w, types.FileUploadReply{Identity: rp.Identity})
			return
		}
		fileName := fileHeader.Filename
		key := helper.GetFileKey(fileName)
		//往oss存储文件
		upload, err := helper.UploadFromByte(key, b)
		if err != nil {
			return
		}
		req.Name = fileName
		req.Ext = path.Ext(fileName)
		fmt.Printf(path.Dir(fileName))
		req.Size = fileHeader.Size
		req.Path = upload
		req.Hash = hash

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
