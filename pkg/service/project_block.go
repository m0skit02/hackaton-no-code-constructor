package service

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"hackaton-no-code-constructor/pkg/dto/project_block_context"
	models "hackaton-no-code-constructor/pkg/model"
	"hackaton-no-code-constructor/pkg/repository"
)

type ProjectBlockService struct {
	repo repository.ProjectBlock
}

func NewProjectBlockService(repo repository.ProjectBlock) *ProjectBlockService {
	return &ProjectBlockService{repo: repo}
}

func (s *ProjectBlockService) CreateProjectBlock(input project_block_context.CreateProjectBlockInput) (*models.ProjectBlock, error) {
	logrus.Infof("[ProjectBlockService] Create Project Block input: %v", input)

	projectBlock := models.ProjectBlock{
		BlockTypeID: input.BlockTypeID,
		Position:    input.Position,
		Data:        input.Data,
	}

	createdProjectBlock, err := s.repo.Create(projectBlock)
	if err != nil {
		logrus.Errorf("[ProjectBlockService] Create Project Block Error: %v", err)
		return nil, err
	}

	logrus.Infof("[ProjectBlockService] Create Project Block success: %v", createdProjectBlock)
	return createdProjectBlock, nil
}

func (s *ProjectBlockService) GetAllProjectBlock() ([]models.ProjectBlock, error) {
	logrus.Infof("[ProjectBlockService] Get All Project Block")
	projectBlocks, err := s.repo.GetAll()
	if err != nil {
		logrus.Errorf("[ProjectBlockService] Get All Project Block Error: %v", err)
		return nil, err
	}

	logrus.Infof("[ProjectBlockService] Get All Project Block success: %v", projectBlocks)
	return projectBlocks, nil
}

func (s *ProjectBlockService) GetByIDProjectBlock(id uuid.UUID) (*models.ProjectBlock, error) {
	logrus.Infof("[ProjectBlockService] Get By ID %v", id)

	projectBlock, err := s.repo.GetByID(id)
	if err != nil {
		logrus.Errorf("[ProjectBlockService] Get By ID Error: %v", err)
		return nil, err
	}

	logrus.Infof("[ProjectBlockService] Get By ID success: %v", projectBlock)
	return projectBlock, nil
}

func (s *ProjectBlockService) UpdateProjectBlock(id uuid.UUID, input project_block_context.UpdateProjectBlockInput) (*models.ProjectBlock, error) {
	logrus.Infof("[ProjectBlockService] Update Project Block input: %+v", input)

	projectBlock, err := s.repo.GetByID(id)
	if err != nil {
		logrus.Errorf("[ProjectBlockService] Get Project Block Error: %v", err)
		return nil, err
	}

	if input.BlockTypeID != nil {
		projectBlock.BlockTypeID = *input.BlockTypeID
	}
	if input.Position != nil {
		projectBlock.Position = *input.Position
	}
	if input.Data != nil {
		projectBlock.Data = *input.Data
	}

	updatedProjectBlock, err := s.repo.Update(*projectBlock) // передаём значение, не указатель
	if err != nil {
		logrus.Errorf("[ProjectBlockService] Update Project Block Error: %v", err)
		return nil, err
	}

	logrus.Infof("[ProjectBlockService] Update Project Block success: %+v", updatedProjectBlock)
	return updatedProjectBlock, nil
}

func (s *ProjectBlockService) DeleteProjectBlock(id uuid.UUID) error {
	logrus.Infof("[ProjectBlockService] Delete By ID %v", id)
	if err := s.repo.Delete(id); err != nil {
		logrus.Errorf("[ProjectBlockService] Delete By ID Error: %v", err)
		return err
	}

	logrus.Infof("[ProjectBlockService] Delete By ID success: %v", id)
	return nil
}
