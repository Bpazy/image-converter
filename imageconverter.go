package image_converter

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Serve(addr string) {
	router := gin.Default()
	router.MaxMultipartMemory = 100 << 20 // 100 MiB
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Printf("Upload file: %s\n", file.Filename)

		targetType, ok := c.GetPostForm("type")
		if !ok {
			panic("请选择目标类型")
		}

		path := saveFile(c, file)
		img, err := Decode(path)
		PanicNonNil(err)

		err = Encode(img, "./output/"+RemovePathExt(filepath.Base(path))+"."+targetType, targetType)
		PanicNonNil(err)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	if err := router.Run(addr); err != nil {
		panic(err)
	}
}

func saveFile(c *gin.Context, file *multipart.FileHeader) string {
	// Make sure tmp directory exists
	_, err := os.Open("upload")
	if err != nil {
		_ = os.Mkdir("upload", os.ModeDir)
	}

	// Upload the file to specific dst.
	id := uuid.New().String()
	path := "./upload/" + id + "." + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		panic("保存文件失败")
	}

	return path
}

func Encode(img image.Image, path, targetType string) error {
	_ = os.MkdirAll(filepath.Dir(path), os.ModeDir)
	fw, _ := os.Create(path)
	defer fw.Close()

	switch targetType {
	case "bmp":
		return bmp.Encode(fw, img)
	case "gif":
		return gif.Encode(fw, img, nil)
	case "jpeg", "jpg":
		return jpeg.Encode(fw, img, nil)
	case "png":
		return png.Encode(fw, img)
	case "tiff":
		return tiff.Encode(fw, img, nil)
	default:
		return errors.New(fmt.Sprintf("The type:[%s] not in support list", targetType))
	}
}

func Decode(path string) (img image.Image, err error) {
	f, _ := os.Open(path)
	defer f.Close()

	ext := filepath.Ext(path)
	switch ext {
	case ".bmp":
		img, err = bmp.Decode(f)
	case ".gif":
		img, err = gif.Decode(f)
	case ".jpeg", ".jpg":
		img, err = jpeg.Decode(f)
	case ".png":
		img, err = png.Decode(f)
	case ".tiff":
		img, err = tiff.Decode(f)
	case ".webp":
		img, err = webp.Decode(f)
	default:
		text := fmt.Sprintf("The type:[%s] not in support list", ext)
		return nil, errors.New(text)
	}

	return img, nil
}

func RemovePathExt(path string) string {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[:i]
		}
	}
	return path
}

func PanicNonNil(err error) {
	if err != nil {
		panic(err)
	}
}
