package controllers

import (
	"contact-management-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllContact(c *gin.Context) {
	var contact []models.Contact

	models.DB.Find(&contact)
	c.JSON(200, gin.H{
		"statusCode": 200,
		"data":       contact,
	})
}

func ShowContactById(c *gin.Context) {
	var contact models.Contact
	id := c.Param("id")
	var findContact = models.DB.First(&contact, id)

	if err := findContact.Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(404, gin.H{
				"statusCode": 404,
				"message":    "Data dengan ID tersebut tidak ditemukan!",
			})
			return
		default:
			c.AbortWithStatusJSON(404, gin.H{
				"statusCode": 404,
				"message":    err.Error(),
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"statusCode": 200,
		"data":       contact,
	})
}

func CreateContact(c *gin.Context) {
	var contact models.Contact

	// mengecek apakah format json/data valid atau tidak
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"statusCode": 400,
			"message":    "Format JSON tidak valid!",
		})
		return
	}

	if err := models.DB.Create(&contact).Error; err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"statusCode": 401,
			"message":    "Gagal menambahkan kontak!",
		})
		return
	}

	models.DB.Create(&contact)

	c.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "Kontak berhasil ditambahkan!",
		"data":       contact,
	})
}

func UpdateContact(c *gin.Context) {
	var contact models.Contact
	id := c.Param("id")

	if err := models.DB.First(&contact, id).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"statusCode": 404,
			"message":    "Kontak dengan ID tersebut tidak ditemukan!",
		})
		return
	}

	var input models.Contact
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"statusCode": 400,
			"message":    "Format data tidak valid!",
		})
		return
	}

	models.DB.Model(&contact).Updates(input)

	c.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "Kontak berhasil diperbarui!",
		"data":       contact,
	})
}

func DeleteContact(c *gin.Context) {
	var contact models.Contact
	id := c.Param("id")
	var findContactById = models.DB.First(&contact, id)

	if err := findContactById.Error; err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"statusCode": 400,
			"message":    "Kontak dengan ID tersebut tidak ada!",
		})
	}

	if err := models.DB.Delete(&contact).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"statuscode": 404,
			"message":    "Gagal menghapus kontak!",
		})
	}

	models.DB.Delete(&contact)
	c.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "Kontak berhasil dihapus!",
	})
}
