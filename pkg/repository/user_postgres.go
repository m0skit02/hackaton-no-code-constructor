package repository

import "gorm.io/gorm"

type UserPostgres struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserPostgres { return &UserPostgres{db: db} }

func (r *UserPostgres) Create() {}

func (r *UserPostgres) GetAll() {}

func (r *UserPostgres) GetByID() {}

func (r *UserPostgres) Update() {}

func (r *UserPostgres) Delete() {}
