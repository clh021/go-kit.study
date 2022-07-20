package main

type UploadService interface {
	// 上传好的文件列表
	list() []string
	// 上传文件
	// 下载文件
}
