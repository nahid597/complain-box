package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nahidh597/complain-box/database"
	"github.com/nahidh597/complain-box/models"
	"github.com/nahidh597/complain-box/utils"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var blogPost models.Blog

	if err := c.BodyParser(&blogPost); err != nil {
		fmt.Println("Unable to body parser")
	}

	err := database.DB.Create(&blogPost).Error

	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Congratulations, you post is live now",
	})

}

func AllPosts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getblogs []models.Blog

	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblogs)
	database.DB.Model(&models.Blog{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": getblogs,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func DetailsPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogPost models.Blog

	database.DB.Where("id=?", id).Preload("User").First(&blogPost)

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Successfully get the post",
		"data":    blogPost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	blog := models.Blog{
		Id: uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&blog).Updates(blog)

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Post updated successfully",
	})
}

func UniquePostsByUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := utils.ParseJwt(cookie)
	var blog []models.Blog

	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)

	c.Status(200)

	return c.JSON(fiber.Map{
		"message": "Successfully get all posts",
		"data":    blog,
	})
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	blog := models.Blog{
		Id: uint(id),
	}

	deleteQuery := database.DB.Delete(&blog)

	fmt.Println(deleteQuery.Error)

	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Oops!, record not found",
		})
	}

	c.Status(200)

	return c.JSON(fiber.Map{
		"message": "post deleted successfully",
	})
}
