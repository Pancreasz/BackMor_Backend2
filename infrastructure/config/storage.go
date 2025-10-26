package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func UploadToAzureBlob(fileName string, fileData []byte) (string, error) {
	fmt.Println(os.Getenv("AZURE_STORAGE_CONNECTION_STRING"))
	fmt.Println(os.Getenv("AZURE_BLOB_CONTAINER"), "asdfasdfasdfasdf")

	connStr := os.Getenv("AZURE_STORAGE_CONNECTION_STRING")
	containerName := os.Getenv("AZURE_BLOB_CONTAINER")

	// Create client
	client, err := azblob.NewClientFromConnectionString(connStr, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create Azure Blob client: %v", err)
	}

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Upload blob
	_, err = client.UploadBuffer(ctx, containerName, fileName, fileData, nil)
	if err != nil {
		return "", fmt.Errorf("failed to upload blob: %v", err)
	}

	// Public URL
	publicURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s",
		os.Getenv("AZURE_STORAGE_ACCOUNT"), containerName, fileName)

	return publicURL, nil
}
