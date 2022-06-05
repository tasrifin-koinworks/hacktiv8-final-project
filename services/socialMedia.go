package services

import (
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocialMediaService struct {
	socialMediaRepo repositories.SocialMediaRepo
}

func NewSocialMediaService(repo repositories.SocialMediaRepo) *SocialMediaService {
	return &SocialMediaService{
		socialMediaRepo: repo,
	}
}

var allSocialMedias []params.GetAllSocialMedias

func (s *SocialMediaService) CreateSocialMedia(request params.CreateSocialMedia) *params.Response {
	socialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
		UserID:         request.UserID,
	}

	socialMediaData, err := s.socialMediaRepo.CreateSocialMedia(&socialMedia)

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
		Payload: params.CreateSocialMedia{
			ID:             socialMediaData.ID,
			Name:           socialMediaData.Name,
			SocialMediaUrl: socialMediaData.SocialMediaUrl,
			UserID:         socialMediaData.UserID,
			CreatedAt:      socialMediaData.CreatedAt,
		},
	}

}

func (s *SocialMediaService) GetAllSocialMedias() *params.Response {
	allSocialMedias = nil
	socialMediaDatas, err := s.socialMediaRepo.GetAllSocialMedias()

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	for _, v := range *socialMediaDatas {
		allSocialMedias = append(allSocialMedias, params.GetAllSocialMedias{
			ID:             v.ID,
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			UserID:         v.UserID,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			User: &params.UserSocialMedia{
				ID:       v.User.ID,
				Username: v.User.Username,
			},
		})
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"social_medias": allSocialMedias,
		},
	}

}

func (s *SocialMediaService) UpdateSocialMedia(request params.CreateSocialMedia) *params.Response {
	socialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
	}

	checkData, err := s.socialMediaRepo.GetSocialMediaByID(request.ID)

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

	updateData, err := s.socialMediaRepo.UpdateSocialMedia(request.ID, &socialMedia)

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
		Payload: params.CreateSocialMedia{
			ID:             updateData.ID,
			Name:           updateData.Name,
			SocialMediaUrl: updateData.SocialMediaUrl,
			UserID:         updateData.UserID,
			CreatedAt:      updateData.CreatedAt,
			UpdatedAt:      updateData.UpdatedAt,
		},
	}
}

func (s *SocialMediaService) DeleteSocialMedia(request params.CreateSocialMedia) *params.Response {
	checkData, err := s.socialMediaRepo.GetSocialMediaByID(request.ID)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	if checkData.ID < 1 || checkData == nil {
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

	err = s.socialMediaRepo.DeleteSocialMedia(request.ID)

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
			"message": "Your social media has been successfully deleted",
		},
	}
}
