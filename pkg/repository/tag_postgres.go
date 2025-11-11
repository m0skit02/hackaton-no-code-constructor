package repository

import "gorm.io/gorm"

type TagPostgres struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) *TagPostgres { return &TagPostgres{db: db} }

func (r *TagPostgres) Create() {}

func (r *TagPostgres) GetAll() {}

func (r *TagPostgres) GetByID() {}

func (r *TagPostgres) Update() {}

func (r *TagPostgres) Delete() {}
