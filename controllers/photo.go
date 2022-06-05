package controllers

import (
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService *services.PhotoService
}

func NewPhotoController(service *services.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: service,
	}
}

func (p *PhotoController) CreatePhoto(c *gin.Context) {
	var req params.CreatePhoto

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	userId, _ := strconv.Atoi(c.GetString("id"))

	if userId < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": "No login information",
			},
		})

		return
	}

	req.UserID = int(userId)
	result := p.photoService.CreatePhoto(req)

	c.JSON(result.Status, result.Payload)
}

func (p *PhotoController) GetAllPhotos(c *gin.Context) {
	result := p.photoService.GetAllPhotos()

	c.JSON(result.Status, result.Payload)
}

func (p *PhotoController) UpdatePhoto(c *gin.Context) {
	var req params.CreatePhoto
	userId, _ := strconv.Atoi(c.GetString("id"))
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	req.ID = photoId
	req.UserID = userId

	result := p.photoService.UpdatePhoto(&req)

	c.JSON(result.Status, result.Payload)

}

func (p *PhotoController) DeletePhoto(c *gin.Context) {
	var req params.CreatePhoto
	userId, _ := strconv.Atoi(c.GetString("id"))
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	req.ID = photoId
	req.UserID = userId

	result := p.photoService.DeletePhoto(&req)

	c.JSON(result.Status, result.Payload)
}
