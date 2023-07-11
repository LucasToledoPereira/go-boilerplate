package studio_repository

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"

	"gorm.io/gorm/clause"
)

type StudioRepo struct {
	data gbp.IDatastoreAdapter
}

func NewStudioRepository(data gbp.IDatastoreAdapter) (r StudioRepo) {
	return StudioRepo{
		data: data,
	}
}

func (sr *StudioRepo) Create(studio *entity.Studio) error {
	err := sr.data.DB().Create(&studio).Error
	return err
}

func (sr *StudioRepo) Update(studio *entity.Studio) error {
	err := sr.data.DB().Save(&studio).Error
	return err
}

func (sr *StudioRepo) Delete(id uuid.UUID) error {
	err := sr.data.DB().Delete(&entity.Studio{}, id).Error
	return err
}

func (sr *StudioRepo) List() (studios []entity.Studio, err error) {
	err = sr.data.DB().Find(&studios).Error
	return studios, err
}

func (sr *StudioRepo) Read(id uuid.UUID) (studio entity.Studio, err error) {
	err = sr.data.DB().First(&studio, id).Error
	return studio, err
}

func (sr *StudioRepo) UpdateImage(studio entity.Studio, key string) error {
	query := sr.data.DB()
	query = query.Omit(clause.Associations)
	query = query.Model(&studio)
	query = query.Update("filekey", key)
	return query.Error
}
