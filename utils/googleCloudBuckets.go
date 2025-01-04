package utils

import (
	"context"
	"encoding/json"
	"io"

	"cloud.google.com/go/storage"
)

const crvhub_bucket = "crvhub_cloudbuild"

func WriteBucketFile(path string, data any) error {
	ctx := context.Background()
	cli, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer cli.Close()

	bkt := cli.Bucket(crvhub_bucket)
	fileWriter := bkt.Object(path).NewWriter(ctx)
	msg, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := fileWriter.Write([]byte(msg)); err != nil {
		return err
	}
	if err := fileWriter.Close(); err != nil {
		return err
	}

	return nil
}

func ReadBucketFile(path string) ([]byte, error) {
	ctx := context.Background()
	cli, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	fileReader, err := cli.Bucket(crvhub_bucket).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	return io.ReadAll(fileReader)
}
