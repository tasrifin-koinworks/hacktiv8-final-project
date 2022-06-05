package controllers

import (
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(service *services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		socialMediaService: *service,
	}
}

func (s *SocialMediaController) CreateSocialMedia(c *gin.Context) {
	var req params.CreateSocialMedia

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, _ := strconv.Atoi(c.GetString("id"))

	req.UserID = userId

	result := s.socialMediaService.CreateSocialMedia(req)

	c.JSON(result.Status, result.Payload)
}

func (s *SocialMediaController) GetAllSocialMedias(c *gin.Context) {
	result := s.socialMediaService.GetAllSocialMedias()

	c.JSON(result.Status, result.Payload)
}

func (s *SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	var req params.CreateSocialMedia

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, _ := strconv.Atoi(c.GetString("id"))
	id, _ := strconv.Atoi(c.Param("socialMediaId"))
	req.UserID = userId
	req.ID = id

	result := s.socialMediaService.UpdateSocialMedia(req)

	c.JSON(result.Status, result.Payload)

}

func (s *SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	var req params.CreateSocialMedia

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, _ := strconv.Atoi(c.GetString("id"))
	id, _ := strconv.Atoi(c.Param("socialMediaId"))
	req.UserID = userId
	req.ID = id

	result := s.socialMediaService.DeleteSocialMedia(req)

	c.JSON(result.Status, result.Payload)

}
