package repository

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	models "hackaton-no-code-constructor/pkg/model"
)

type ProjectBlockPostgres struct {
	db *gorm.DB
}

func NewProjectBlockRepo(db *gorm.DB) *ProjectBlockPostgres { return &ProjectBlockPostgres{db: db} }

func (r *ProjectBlockPostgres) Create(projectBlock models.ProjectBlock) (*models.ProjectBlock, error) {
	logrus.Infof("[UserRepo] Create projectBlock %+v", projectBlock)
	if err := r.db.Create(&projectBlock).Error; err != nil {
		logrus.Errorf("[UserRepo] Create projectBlock err %v", err)
		return nil, err
	}
	logrus.Infof("[UserRepo] Create projectBlock successfully: ID=%s", projectBlock.ID)
	return &projectBlock, nil
}

func (r *ProjectBlockPostgres) GetAll() ([]models.ProjectBlock, error) {
	logrus.Infof("[UserRepo] GetAll projectBlocks")
	var projectBlocks []models.ProjectBlock
	if err := r.db.Find(&projectBlocks).Error; err != nil {
		logrus.Errorf("[UserRepo] GetAll projectBlocks err %v", err)
		return nil, err
	}

	logrus.Infof("[UserRepo] GetAll projectBlocks successfully")
	return projectBlocks, nil
}

func (r *ProjectBlockPostgres) GetByID(id uuid.UUID) (*models.ProjectBlock, error) {
	logrus.Infof("[UserRepo] GetByID projectBlock %+v", id)
	var projectBlock models.ProjectBlock
	if err := r.db.First(&projectBlock, id).Error; err != nil {
		if err := gorm.ErrRecordNotFound; err != nil {
			logrus.Errorf("[UserRepo] GetByID projectBlock err %v", err)
			return nil, nil
		}
		logrus.Infof("[UserRepo] GetByID projectBlock successfully")
		return nil, err
	}
	logrus.Infof("[UserRepo] GetByID projectBlock successfully")
	return &projectBlock, nil
}

func (r *ProjectBlockPostgres) Update(projectBlock models.ProjectBlock) (*models.ProjectBlock, error) {
	logrus.Infof("[UserRepo] Update projectBlock %+v", projectBlock)
	if err := r.db.Save(&projectBlock).Error; err != nil {
		logrus.Errorf("[UserRepo] Update projectBlock err %v", err)
		return nil, err
	}
	logrus.Infof("[UserRepo] Update projectBlock successfully: ID=%s", projectBlock.ID)
	return &projectBlock, nil
}

func (r *ProjectBlockPostgres) Delete(id uuid.UUID) error {
	logrus.Infof("[UserRepo] Delete projectBlock %+v", id)
	if err := r.db.Delete(&models.ProjectBlock{}, id).Error; err != nil {
		logrus.Errorf("[UserRepo] Delete projectBlock err %v", err)
		return err
	}
	logrus.Infof("[UserRepo] Delete projectBlock successfully: ID=%s", id)
	return nil
}
