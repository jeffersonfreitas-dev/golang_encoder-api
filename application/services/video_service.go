package services

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/jeffersonfreitas-dev/encoder-api/application/repositories"
	"github.com/jeffersonfreitas-dev/encoder-api/domain"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(v.Video.FilePath)

	reader, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}

	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	file, err := os.Create(v.Video.ID + ".mp4")
	if err != nil {
		return err
	}

	_, err = file.Write(body)
	if err != nil {
		return err
	}

	defer file.Close()
	log.Printf("video %v has been storaged", v.Video.ID)

	return nil
}
