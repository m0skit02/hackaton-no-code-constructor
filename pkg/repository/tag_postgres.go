package repository

import (
	models "hackaton-no-code-constructor/pkg/model"

	"gorm.io/gorm"
)

type TagPostgres struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) *TagPostgres { return &TagPostgres{db: db} }

func (r *TagPostgres) Create(tag models.Tag) (*models.Tag, error) {
	if err := r.db.Create(&tag).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r *TagPostgres) GetAll() ([]models.Tag, error) {
	var tags []models.Tag
	if err := r.db.Order("created_at desc").Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *TagPostgres) GetByID(id string) (*models.Tag, error) {
	var tag models.Tag
	if err := r.db.First(&tag, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r *TagPostgres) Update(tag *models.Tag) (*models.Tag, error) {
	if err := r.db.Save(&tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

func (r *TagPostgres) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Tag{}).Error; err != nil {
		return err
	}

	return nil
}
