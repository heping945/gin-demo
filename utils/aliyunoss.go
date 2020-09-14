package utils

import (
	"gin-demo/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func AliyunOssObject() *oss.Bucket {
	aliyunoss := global.GVA_CONFIG.AliYunOSS
	client, err := oss.New(aliyunoss.Endpoint, aliyunoss.AccessKeyID, aliyunoss.AccessKeySecret)
	if err != nil {
		global.GVA_LOG.Panic("aliyunoss init failer", err.Error())
	}
	// 获取存储空间。
	bucket, err := client.Bucket(aliyunoss.BucketName)
	if err != nil {
		global.GVA_LOG.Panic("get budget failer", err.Error())
	}
	global.GVA_LOG.Info("生成bucket成功")
	return bucket
}
