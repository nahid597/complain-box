package controller

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randLetter(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func UploadFile(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["image"]
	fileName := ""

	fmt.Println("file", files)

	for _, file := range files {
		fileName = randLetter(5) + "-" + file.Filename

		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return nil
		}
	}

	fmt.Println("filename", fileName)

	c.Status(200)

	return c.JSON(fiber.Map{
		"message": "File save successfully",
		"url":     "http://localhost:3000/api/uploads/" + fileName,
	})

}
