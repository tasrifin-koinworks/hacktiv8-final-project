package repositories

import (
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error)
	CheckUser(email string, user *models.User) error
	CheckUserByID(id int, user *models.User) (*models.User, error)
	DeleteUser(userId int) error
	UpdateUser(userId int, user *models.User) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (u *userRepo) CreateUser(user *models.User) (*models.User, error) {
	return user, u.db.Create(&user).Error
}

func (u *userRepo) CheckUser(email string, user *models.User) error {
	return u.db.Where("email=?", email).Take(&user).Error
}

func (u *userRepo) CheckUserByID(id int, user *models.User) (*models.User, error) {
	return user, u.db.Where("id=?", id).Take(&user).Error
}

func (u *userRepo) DeleteUser(userId int) error {
	var user models.User
	return u.db.Where("id=?", userId).Delete(&user).Error
}

func (u *userRepo) UpdateUser(userId int, userUpdate *models.User) (*models.User, error) {
	var user models.User

	result := u.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", userId).Updates(userUpdate)
	return &user, result.Error
}
