package article

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func InitMyCloudinary() (*cloudinary.Cloudinary, error) {
	name := os.Getenv("CLOUDINARY_NAME")
	apikey := os.Getenv("CLOUDINARY_API_KEY")
	apisecreet := os.Getenv("CLOUDINARY_API_SECRET")
	cld, err := cloudinary.NewFromParams(name, apikey, apisecreet)

	if err != nil {
		return nil, err
	}

	return cld, nil
}

func UploadToCloudinary(file io.Reader, filename string) (string, error) {
	// Inisialisasi Cloudinary
	cld, err := InitMyCloudinary()
	if err != nil {
		return "Error init cloudinary", err
	}

	// Upload image to Cloudinary
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder:   "thumbnail", // Folder in Cloudinary where the image will be stored
		PublicID: filename,    // File name in Cloudinary
	})
	if err != nil {
		return "", err
	}

	// Return public URL of the uploaded image
	return uploadResult.SecureURL, nil
}

func DeleteFromCloudinary(filename string) error {
	// Inisialisasi Cloudinary
	cld, err := InitMyCloudinary()
	if err != nil {
		return err
	}

	publicID, _ := extractPublicID(filename)

	// Delete image from Cloudinary
	_, err = cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return err
	}

	return nil
}

func extractPublicID(fileURL string) (string, error) {
	parts := strings.Split(fileURL, "/")
	for i, part := range parts {
		if part == "upload" {
			// Join all parts after "upload" except the file extension
			publicID := strings.Join(parts[i+1:], "/")
			// Remove the file extension
			publicID = publicID[:strings.LastIndex(publicID, ".")]
			return publicID, nil
		}
	}
	return "", fmt.Errorf("invalid Cloudinary URL")
}
