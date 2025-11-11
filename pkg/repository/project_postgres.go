package repository

import "gorm.io/gorm"

type ProjectPostgres struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectPostgres { return &ProjectPostgres{db: db} }

func (r *ProjectPostgres) Create() {}

func (r *ProjectPostgres) GetAll() {}

func (r *ProjectPostgres) GetByID() {}

func (r *ProjectPostgres) Update() {}

func (r *ProjectPostgres) Delete() {}
