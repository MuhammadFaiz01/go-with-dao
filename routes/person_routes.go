package routes

import (
	"go-dao/models"
	"go-dao/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupPersonRoutes(r *gin.Engine, service *services.PersonService) {
	personRoutes := r.Group("/person")

	personRoutes.POST("/", func(c *gin.Context) {
		var p models.Person
		if err := c.BindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := service.CreatePerson(&p); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create person"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Person created successfully"})
	})

	personRoutes.PUT("/", func(c *gin.Context) {
		idStr := c.Query("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var p models.Person
		if err := c.BindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := service.UpdatePerson(id, &p); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update person"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Person updated successfully"})
	})

	personRoutes.GET("/", func(c *gin.Context) {
		fullName := c.Query("full_name")
		persons, err := service.GetPersonByName(fullName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve person"})
			return
		}
		c.JSON(http.StatusOK, persons)
	})

	personRoutes.GET("/all", func(c *gin.Context) {
		persons, err := service.GetAllPersons()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve persons"})
			return
		}
		c.JSON(http.StatusOK, persons)
	})

	personRoutes.DELETE("/", func(c *gin.Context) {
		idStr := c.Query("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := service.DeletePerson(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete person"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
	})

	personRoutes.GET("/export", func(c *gin.Context) {
		filepath := "persons.xlsx"
		if err := service.ExportPersons(filepath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to export persons"})
			return
		}
		c.File(filepath)
	})

	personRoutes.POST("/import", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
			return
		}

		filePath := "uploaded_" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
			return
		}

		if err := service.ImportFromExcel(filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to import data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data imported successfully"})
	})
}
