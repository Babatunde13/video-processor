package main

import (
	"fmt"
	"io"
	"log"
	"media_processor/internals/files"
	"media_processor/internals/stream"
	"os"

	"github.com/gin-gonic/gin"
)

func StreamHandler(c *gin.Context) {
	filename := c.Param("filename")

	file, err := os.Open("srt-data/" + filename)
	if err != nil {
		log.Println("Error opening file", err)
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	defer file.Close()
	// turn it to a buffer
	c.Stream(func(w io.Writer) bool {
		_, err := io.Copy(w, file)
		if err != nil {
			log.Println("Error copying file", err)
		}
		return false
	})
}

func UploadHandler(c *gin.Context) {
	// get file buffer
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error getting file", err)
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	// save file to disk
	filename := "output.mp4"
	err = files.SaveFile(file, filename)
	if err != nil {
		log.Println("Error saving file", err)
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "file uploaded",
	})
}

func StartProcessingHandler(c *gin.Context) {
	go stream.Ingest("output.m3u8")
	c.JSON(200, gin.H{
		"message": "processing started",
	})
}

func startRouter() {
	api := gin.Default()
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	api.Static("/hls/stream", stream.FolderName)
	api.POST("/upload", UploadHandler)
	api.POST("/start-processing", StartProcessingHandler)
	err := api.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}

func main() {
	os.MkdirAll(stream.FolderName, os.ModePerm)
	fmt.Println("starting hls server")
	startRouter()
}
