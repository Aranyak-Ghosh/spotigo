package repository

import uuid "github.com/satori/go.uuid"

type BaseRepository interface {
	GetById(id uuid.UUID) (*interface{}, error)
	List(searchParam string, offset int, limit int) ([]interface{}, int, error)
	Create(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
}