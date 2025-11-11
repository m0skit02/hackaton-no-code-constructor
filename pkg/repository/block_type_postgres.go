package repository

import "gorm.io/gorm"

type BlockTypePostgres struct {
	db *gorm.DB
}

func NewBlockTypeRepo(db *gorm.DB) *BlockTypePostgres { return &BlockTypePostgres{db: db} }

func (r *BlockTypePostgres) Create() {}

func (r *BlockTypePostgres) GetAll() {}

func (r *BlockTypePostgres) GetByID() {}

func (r *BlockTypePostgres) Update() {}

func (r *BlockTypePostgres) Delete() {}
