package routes

import (
	"go-dao/models"
	"go-dao/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupKelasRoutes(r *gin.Engine, service *services.KelasService) {
	kelasRoutes := r.Group("/kelas")

	kelasRoutes.POST("/", func(c *gin.Context) {
		var k models.Kelas
		if err := c.BindJSON(&k); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := service.CreateKelas(&k); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create kelas"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Kelas created successfully"})
	})

	kelasRoutes.PUT("/", func(c *gin.Context) {
		idStr := c.DefaultQuery("id", "")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var k models.Kelas
		if err := c.BindJSON(&k); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := service.UpdateKelas(id, &k); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update kelas"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Kelas updated successfully"})
	})

	kelasRoutes.GET("/", func(c *gin.Context) {
		namaKelas := c.Query("namaKelas")
		if namaKelas == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nama kelas is required"})
			return
		}

		kelas, err := service.GetKelasByName(namaKelas)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get kelas"})
			return
		}

		c.JSON(http.StatusOK, kelas)

	})

	kelasRoutes.GET("/all", func(c *gin.Context) {
		kelas, err := service.GetAllKelas()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get kelas"})
			return
		}
		c.JSON(http.StatusOK, kelas)
	})

	kelasRoutes.DELETE("/", func(c *gin.Context) {
		idStr := c.DefaultQuery("id", "")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := service.DeleteKelas(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete kelas"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Kelas deleted successfully"})
	})
}
