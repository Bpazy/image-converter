package image_converter

import (
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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

func PanicNonNil(err error) {
	if err != nil {
		panic(err)
	}
}
