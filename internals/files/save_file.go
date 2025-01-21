package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func SaveFile(dile *multipart.FileHeader, path string) error {
	fullPath := "./srt-data/" + path
	// if file exists, overwrite it
	dst, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	defer dst.Close()
	src, err := dile.Open()
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer src.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	return nil
}
