package image_converter

import (
	"errors"
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
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

		path := saveFile(c, file)
		img, err := imgio.Open(path)
		PanicNonNil(err)
		imgio.Save("output.jpg", img, imgio.JPEGEncoder(100))

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	if err := router.Run(addr); err != nil {
		panic(err)
	}
}

func saveFile(c *gin.Context, file *multipart.FileHeader) string {
	// Make sure tmp directory exists
	_, err := os.Open("tmp")
	if err != nil {
		os.Mkdir("tmp", os.ModeDir)
	}

	// Upload the file to specific dst.
	id := uuid.New().String()
	path := "./tmp/" + id + "." + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		panic("保存文件失败")
	}

	return path
}

func Encode(img image.Image, filename, Type string) error {
	fw, _ := os.Create(filename)
	defer fw.Close()

	switch Type {
	case "bmp":
		bmp.Encode(fw, img)
	case "gif":
		gif.Encode(fw, img, nil)
	case "jpeg", "jpg":
		jpeg.Encode(fw, img, nil)
	case "png":
		png.Encode(fw, img)
	case "tiff":
		tiff.Encode(fw, img, nil)
	default:
		text := fmt.Sprintf("The type:[%s] not in support list", Type)
		fmt.Println(text)
	}

	fmt.Printf("Convert %s success\n", filename)

	return nil
}

func Decode(filename string) (img image.Image, err error) {
	f, _ := os.Open(filename)
	defer f.Close()

	ext := filepath.Ext(filename)
	switch ext {
	case ".bmp":
		img, err = bmp.Decode(f)
	case ".gif":
		img, err = gif.Decode(f)
	case "jpeg", "jpg":
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
