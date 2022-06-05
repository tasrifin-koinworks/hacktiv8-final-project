package repositories

import (
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepo interface {
	CreatePhoto(photo *models.Photo) (*models.Photo, error)
	GetAllPhotos() (*[]models.Photo, error)
	GetPhotoByID(photoId int) (*models.Photo, error)
	UpdatePhoto(photoId int, updatePhoto *models.Photo) (*models.Photo, error)
	DeletePhoto(photoId int) error
}

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) PhotoRepo {
	return &photoRepo{db}
}

func (p *photoRepo) CreatePhoto(photo *models.Photo) (*models.Photo, error) {
	return photo, p.db.Create(&photo).Error
}

func (p *photoRepo) GetAllPhotos() (*[]models.Photo, error) {
	var photo []models.Photo
	err := p.db.Preload("User").Find(&photo).Error
	return &photo, err
}

func (p *photoRepo) GetPhotoByID(photoId int) (*models.Photo, error) {
	var photo models.Photo
	err := p.db.Preload("User").Where("id=?", photoId).Find(&photo).Error
	return &photo, err
}

func (p *photoRepo) UpdatePhoto(photoId int, updatePhoto *models.Photo) (*models.Photo, error) {
	var photo models.Photo

	err := p.db.Model(&photo).Clauses(clause.Returning{}).Where("id=?", photoId).Updates(updatePhoto).Error
	return &photo, err
}

func (p *photoRepo) DeletePhoto(photoId int) error {
	var photo models.Photo

	err := p.db.Where("id=?", photoId).Delete(&photo).Error
	return err
}
