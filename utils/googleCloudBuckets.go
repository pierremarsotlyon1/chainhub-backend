package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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
	return readLatestFile(crvhub_bucket, path)
}

// generateFileURL ajoute un timestamp à l'URL pour éviter la mise en cache
func generateFileURL(bucketName, objectName string) string {
	// Get current timestamp
	timestamp := time.Now().Unix()
	// Add one day to be sure to have the last file
	timestamp += int64(time.Now().Day())
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s?ts=%d", bucketName, objectName, timestamp)
}

// readFileFromURL télécharge le fichier en contournant le cache
func readFileFromURL(fileURL string) ([]byte, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch file, status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}

func readLatestFile(bucketName, objectName string) ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}
	defer client.Close()

	// Liste les attributs du fichier pour récupérer la dernière génération
	obj := client.Bucket(bucketName).Object(objectName)
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get attributes: %w", err)
	}

	// Lecture du fichier avec l’ID de génération
	reader, err := obj.Generation(attrs.Generation).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader: %w", err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}

	return data, nil
}
