package repositories

import (
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepo interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	GetAllComments() (*[]models.Comment, error)
	GetCommentByID(commentId int) (*models.Comment, error)
	UpdateComment(commentId int, comment *models.Comment) (*models.Comment, error)
	DeleteComment(commentId int) error
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepo{db}
}

func (c *commentRepo) CreateComment(comment *models.Comment) (*models.Comment, error) {
	return comment, c.db.Create(&comment).Error
}

func (c *commentRepo) GetAllComments() (*[]models.Comment, error) {
	var comments []models.Comment
	err := c.db.Preload("User").Preload("Photo").Find(&comments).Error
	return &comments, err

}

func (c *commentRepo) GetCommentByID(commentId int) (*models.Comment, error) {
	var comment models.Comment
	err := c.db.Preload("User").Preload("Photo").Where("id=?", commentId).First(&comment).Error
	return &comment, err
}

func (c *commentRepo) UpdateComment(commentId int, updateComment *models.Comment) (*models.Comment, error) {
	var comment models.Comment

	err := c.db.Model(&comment).Clauses(clause.Returning{}).Where("id=?", commentId).Updates(&updateComment).Error
	return &comment, err
}

func (c *commentRepo) DeleteComment(commentId int) error {
	var comment models.Comment
	err := c.db.Where("id=?", commentId).Delete(&comment).Error
	return err
}
