package repository

import "gorm.io/gorm"

type ProjectBlockPostgres struct {
	db *gorm.DB
}

func NewProjectBlockRepo(db *gorm.DB) *ProjectBlockPostgres { return &ProjectBlockPostgres{db: db} }

func (r *ProjectBlockPostgres) Create() {}

func (r *ProjectBlockPostgres) GetAll() {}

func (r *ProjectBlockPostgres) GetByID() {}

func (r *ProjectBlockPostgres) Update() {}

func (r *ProjectBlockPostgres) Delete() {}
