package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	savePath := fmt.Sprintf("./uploads/%s", file.Filename)

	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't save file"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Upload successful!",
		"filename": file.Filename,
		"path":     savePath,
	})
}

func UploadMultipleFiles(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data upload is invalid"})
		return
	}

	files := form.File["files"]
	var paths []string

	for _, file := range files {
		savePath := fmt.Sprintf("./uploads/%s", file.Filename)
		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		paths = append(paths, savePath)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Uploaded %d file", len(files)),
		"files":   paths,
	})
}
