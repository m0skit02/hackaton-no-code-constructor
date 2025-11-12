package repository

import (
	models "hackaton-no-code-constructor/pkg/model"

	"gorm.io/gorm"
)

type BlockTypePostgres struct {
	db *gorm.DB
}

func NewBlockTypeRepo(db *gorm.DB) *BlockTypePostgres { return &BlockTypePostgres{db: db} }

func (r *BlockTypePostgres) Create(blockType models.BlockType) (*models.BlockType, error) {
	if err := r.db.Create(&blockType).Error; err != nil {
		return nil, err
	}
	return &blockType, nil
}

func (r *BlockTypePostgres) GetAll() ([]models.BlockType, error) {
	var blockTypes []models.BlockType
	if err := r.db.Order("created_at desc").Find(&blockTypes).Error; err != nil {
		return nil, err
	}
	return blockTypes, nil
}

func (r *BlockTypePostgres) GetByID(id string) (*models.BlockType, error) {
	var blockType models.BlockType
	if err := r.db.First(&blockType, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &blockType, nil
}

func (r *BlockTypePostgres) Update(blockType *models.BlockType) (*models.BlockType, error) {
	if err := r.db.Save(&blockType).Error; err != nil {
		return nil, err
	}

	return blockType, nil
}

func (r *BlockTypePostgres) Delete(id string) error {
	if err := r.db.Delete(&models.BlockType{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
