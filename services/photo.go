package services

import (
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoService struct {
	photoRepo repositories.PhotoRepo
}

func NewPhotoService(repo repositories.PhotoRepo) *PhotoService {
	return &PhotoService{
		photoRepo: repo,
	}
}

var allPhotos []params.GetAllPhotos

func (p *PhotoService) CreatePhoto(request params.CreatePhoto) *params.Response {

	photo := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserID:   request.UserID,
	}

	photoData, err := p.photoRepo.CreatePhoto(&photo)

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
		Payload: &params.GetAllPhotos{
			ID:        photoData.ID,
			Title:     photoData.Title,
			Caption:   photoData.Caption,
			PhotoUrl:  photoData.PhotoUrl,
			UserID:    photoData.UserID,
			CreatedAt: photoData.CreatedAt,
		},
	}
}

func (p *PhotoService) GetAllPhotos() *params.Response {
	result, err := p.photoRepo.GetAllPhotos()

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	allPhotos = nil
	for _, p := range *result {
		allPhotos = append(allPhotos, params.GetAllPhotos{
			ID:        p.ID,
			Title:     p.Title,
			Caption:   p.Caption,
			PhotoUrl:  p.PhotoUrl,
			UserID:    p.UserID,
			CreatedAt: p.CreatedAt,
			UpdateAt:  p.UpdatedAt,
			User: &params.UserPhoto{
				Email:    p.User.Email,
				Username: p.User.Username,
			},
		})
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: allPhotos,
	}
}

func (p *PhotoService) UpdatePhoto(request *params.CreatePhoto) *params.Response {
	photo := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
	}

	checkData, err := p.photoRepo.GetPhotoByID(request.ID)

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

	updatePhoto, err := p.photoRepo.UpdatePhoto(request.ID, &photo)

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
		Payload: &params.GetAllPhotos{
			ID:       updatePhoto.ID,
			Title:    updatePhoto.Title,
			Caption:  updatePhoto.Caption,
			PhotoUrl: updatePhoto.PhotoUrl,
			UserID:   updatePhoto.UserID,
			UpdateAt: updatePhoto.UpdatedAt,
		},
	}

}

func (p *PhotoService) DeletePhoto(request *params.CreatePhoto) *params.Response {
	checkData, err := p.photoRepo.GetPhotoByID(request.ID)

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

	err = p.photoRepo.DeletePhoto(request.ID)

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
			"message": "Your photo has been successfully deleted",
		},
	}

}
