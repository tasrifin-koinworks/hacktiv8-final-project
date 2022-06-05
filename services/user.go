package services

import (
	"hacktiv8-final-project/helpers"
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewItemService(repo repositories.UserRepo) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (u *UserService) UserRegister(request params.CreateUser) *params.Response {
	user := models.User{
		Age:      request.Age,
		Email:    request.Email,
		Password: request.Password,
		Username: request.Username,
	}

	userData, err := u.userRepo.CreateUser(&user)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: map[string]string{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusCreated,
		Payload: gin.H{
			"age":      userData.Age,
			"email":    userData.Email,
			"id":       userData.ID,
			"username": userData.Username,
		},
	}
}

func (u *UserService) Login(request params.CreateUser) *params.Response {
	var userDB models.User

	if request.Email == "" {
		return &params.Response{
			Status: http.StatusUnauthorized,
			Payload: gin.H{
				"message": "Email cannot be null",
			},
		}
	}

	if request.Password == "" {
		return &params.Response{
			Status: http.StatusUnauthorized,
			Payload: gin.H{
				"message": "Password cannot be null",
			},
		}
	}

	err := u.userRepo.CheckUser(request.Email, &userDB)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	dataIsOK := helpers.ComparePassword([]byte(userDB.Password), []byte(request.Password))

	if !dataIsOK {
		return &params.Response{
			Status: http.StatusUnauthorized,
			Payload: gin.H{
				"message": "Password not match",
			},
		}
	}

	token := helpers.GenerateToken(userDB.ID, userDB.Email)

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"token": token,
		},
	}
}

func (u *UserService) UpdateUser(id int, request params.CreateUser) *params.Response {

	if id != request.ID {
		return &params.Response{
			Status: http.StatusForbidden,
			Payload: gin.H{
				"error": "Forbidden - only owner can update data",
			},
		}
	}

	checkData, _ := u.userRepo.CheckUserByID(request.ID, &models.User{})

	if checkData.ID < 1 {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "User not found",
			},
		}
	}

	user := models.User{
		Email:    request.Email,
		Username: request.Username,
	}

	updateUser, err := u.userRepo.UpdateUser(request.ID, &user)

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
		Payload: &params.CreateUser{
			ID:        updateUser.ID,
			Email:     updateUser.Email,
			Username:  updateUser.Username,
			Age:       updateUser.Age,
			UpdatedAt: &updateUser.UpdatedAt,
		},
	}

}

func (u *UserService) DeleteUser(id int, request params.CreateUser) *params.Response {

	if id != request.ID {
		return &params.Response{
			Status: http.StatusForbidden,
			Payload: gin.H{
				"error": "Forbidden - only owner can delete data",
			},
		}
	}

	checkData, _ := u.userRepo.CheckUserByID(request.ID, &models.User{})

	if checkData.ID < 1 {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "User not found",
			},
		}
	}

	err := u.userRepo.DeleteUser(id)

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
			"message": "Your account has been successfully deleted",
		},
	}
}
