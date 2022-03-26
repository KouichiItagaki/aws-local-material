package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/api/v1/cats", getCats)

	e.Logger.Fatal(e.Start(":80"))
}

func getCats(c echo.Context) error {
	bucketName := "local-bucket"
	path := "dir1/"
	endPoint := "http://minio:9999"
	region := "ap-northeast-1"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = region
		o.UsePathStyle = true
		o.EndpointResolver = s3.EndpointResolverFromURL(endPoint)
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	objects, err := getObjectsFromS3(context.TODO(), client, bucketName, path)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, objects)
}

type S3ListObjectsAPI interface {
	ListObjectsV2(ctx context.Context,
		params *s3.ListObjectsV2Input,
		optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

func getObjectsFromS3(ctx context.Context, api S3ListObjectsAPI, bucket, prefix string) ([]string, error) {
	var objects []string
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}

	resp, err := api.ListObjectsV2(context.TODO(), input)
	if err != nil {
		return objects, fmt.Errorf("ERROR: api.ListObjectsV2, %v", err)
	}

	for _, item := range resp.Contents {
		objects = append(objects, *item.Key)
	}

	return objects, nil
}
