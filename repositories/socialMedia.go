package repositories

import (
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepo interface {
	CreateSocialMedia(socialmedia *models.SocialMedia) (*models.SocialMedia, error)
	GetAllSocialMedias() (*[]models.SocialMedia, error)
	GetSocialMediaByID(socialMediaId int) (*models.SocialMedia, error)
	UpdateSocialMedia(socialMediaId int, socialmedia *models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMedia(socialMediaId int) error
}

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepo(db *gorm.DB) SocialMediaRepo {
	return &socialMediaRepo{db}
}

func (s *socialMediaRepo) CreateSocialMedia(socialmedia *models.SocialMedia) (*models.SocialMedia, error) {
	return socialmedia, s.db.Create(&socialmedia).Error
}

func (s *socialMediaRepo) GetAllSocialMedias() (*[]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := s.db.Preload("User").Find(&socialMedias).Error
	return &socialMedias, err
}

func (s *socialMediaRepo) GetSocialMediaByID(socialMediaId int) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia

	err := s.db.Preload("User").Where("id=?", socialMediaId).Find(&socialMedia).Error
	return &socialMedia, err
}

func (s *socialMediaRepo) UpdateSocialMedia(socialMediaId int, socialMediaUpdate *models.SocialMedia) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia

	err := s.db.Model(&socialMedia).Clauses(clause.Returning{}).Where("id=?", socialMediaId).Updates(socialMediaUpdate).Error
	return &socialMedia, err
}

func (s *socialMediaRepo) DeleteSocialMedia(socialMediaId int) error {
	var socialMedia models.SocialMedia

	err := s.db.Where("id=?", socialMediaId).Delete(&socialMedia).Error
	return err
}
