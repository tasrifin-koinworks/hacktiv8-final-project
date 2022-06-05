package services

import (
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentService struct {
	commentRepo repositories.CommentRepo
}

func NewCommentService(repo repositories.CommentRepo) *CommentService {
	return &CommentService{
		commentRepo: repo,
	}
}

var allComments []params.GetAllCommentsWithUserAndPhoto

func (c *CommentService) CreateComment(request params.CreateComment) *params.Response {
	comment := models.Comment{
		UserID:  request.UserID,
		PhotoID: request.PhotoID,
		Message: request.Message,
	}

	commentData, err := c.commentRepo.CreateComment(&comment)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusCreated,
		Payload: &params.CreateComment{
			ID:        commentData.ID,
			Message:   commentData.Message,
			PhotoID:   commentData.PhotoID,
			UserID:    commentData.UserID,
			CreatedAt: commentData.CreatedAt,
		},
	}
}

func (c *CommentService) GetAllComments() *params.Response {
	allComments = nil

	comments, err := c.commentRepo.GetAllComments()

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	for _, com := range *comments {
		allComments = append(allComments, params.GetAllCommentsWithUserAndPhoto{
			ID:        com.ID,
			Message:   com.Message,
			PhotoID:   com.PhotoID,
			UserID:    com.UserID,
			UpdateAt:  com.UpdatedAt,
			CreatedAt: com.CreatedAt,
			User: &params.UserComment{
				ID:       com.User.ID,
				Email:    com.User.Email,
				Username: com.User.Username,
			},
			Photo: &params.PhotoComment{
				ID:       com.Photo.ID,
				Title:    com.Photo.Title,
				Caption:  com.Photo.Caption,
				PhotoUrl: com.Photo.PhotoUrl,
				UserID:   com.Photo.UserID,
			},
		})
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: allComments,
	}
}

func (c *CommentService) UpdateComment(request params.CreateComment) *params.Response {
	comment := models.Comment{
		Message: request.Message,
	}

	checkData, err := c.commentRepo.GetCommentByID(request.ID)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	if checkData.ID < 1 {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "record not found",
			},
		}
	}

	if request.UserID != checkData.UserID {
		return &params.Response{
			Status: http.StatusForbidden,
			Payload: gin.H{
				"error": "forbidden - only owner can update data",
			},
		}
	}

	updateData, err := c.commentRepo.UpdateComment(request.ID, &comment)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: params.CreateComment{
			ID:        updateData.ID,
			Message:   updateData.Message,
			PhotoID:   updateData.PhotoID,
			UserID:    updateData.UserID,
			UpdatedAt: updateData.UpdatedAt,
		},
	}
}

func (c *CommentService) DeleteComment(request params.CreateComment) *params.Response {
	checkData, err := c.commentRepo.GetCommentByID(request.ID)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	if checkData.ID < 1 {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "record not found",
			},
		}
	}

	if request.UserID != checkData.UserID {
		return &params.Response{
			Status: http.StatusForbidden,
			Payload: gin.H{
				"error": "forbidden - only owner can delete data",
			},
		}
	}

	err = c.commentRepo.DeleteComment(request.ID)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message": "Your comment has been successfully deleted",
		},
	}
}
