package controllers

import (
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *services.CommentService
}

func NewCommentContoller(service *services.CommentService) *CommentController {
	return &CommentController{
		commentService: service,
	}
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var req params.CreateComment

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, _ := strconv.Atoi(c.GetString("id"))

	req.UserID = userId

	result := cc.commentService.CreateComment(req)

	c.JSON(result.Status, result.Payload)
}

func (cc *CommentController) GetAllComments(c *gin.Context) {
	result := cc.commentService.GetAllComments()

	c.JSON(result.Status, result.Payload)
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	var req params.CreateComment

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, _ := strconv.Atoi(c.GetString("id"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	req.UserID = userId
	req.ID = commentId

	result := cc.commentService.UpdateComment(req)
	c.JSON(result.Status, result.Payload)
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	var req params.CreateComment

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, _ := strconv.Atoi(c.GetString("id"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	req.UserID = userId
	req.ID = commentId

	result := cc.commentService.DeleteComment(req)
	c.JSON(result.Status, result.Payload)
}
