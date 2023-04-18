package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/1-aquila-1/sdk-go/v1_0/common"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadS3() {

	path := "C:\\Workspace\\Pessoal\\Projetos\\controle-acesso\\banner5.png"

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()

	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	extensao := strings.ReplaceAll(fileType, "image/", "")

	nome := "banner" + "-" + common.Uuid10() + "." + extensao

	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "sa-east-1"
		return nil
	})
	if err != nil {
		log.Fatalln("error:", err)
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String("aquila7"),
		Key:           aws.String(nome),
		Body:          fileBytes,
		ContentLength: *aws.Int64(size),
		ContentType:   aws.String(fileType),
	})

	fmt.Println(result.Location)
}
