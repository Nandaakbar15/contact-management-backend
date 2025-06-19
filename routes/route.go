package routes

import (
	"github.com/gin-gonic/gin"

	"contact-management-backend/controllers"
	// "contact-management-backend/models"
)

func Route(router *gin.Engine) {
	router.GET("/api/v1/getallcontacts", controllers.GetAllContact)       // get all data
	router.GET("/api/v1/showcontact/:id", controllers.ShowContactById)    // get the data by id
	router.POST("/api/v1/addnewcontact", controllers.CreateContact)       // add new data
	router.PUT("/api/v1/updatecontact/:id", controllers.UpdateContact)    // update the data
	router.DELETE("/api/v1/deletecontact/:id", controllers.DeleteContact) // delete the data
}
