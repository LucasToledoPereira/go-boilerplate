package game_repository

import (
	"errors"

	"github.com/LucasToledoPereira/go-boilerplate/adapters/datastore"
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type GameRepo struct {
	data datastore.IDatastoreAdapter
}

func NewGameRepository(data datastore.IDatastoreAdapter) GameRepo {
	return GameRepo{
		data: data,
	}
}

func (gr *GameRepo) Create(game *entity.Game) error {
	err := gr.data.DB().Omit(clause.Associations).Create(&game).Error
	return err
}

func (gr *GameRepo) Update(game *entity.Game) error {
	query := gr.data.DB()
	query = query.Omit(clause.Associations)
	query = query.Save(&game)
	err := query.Error
	return err
}

func (gr *GameRepo) Delete(gameID uuid.UUID, studioID uuid.UUID, deleter uuid.UUID) error {
	u := entity.Game{ID: gameID}
	u.DeletedBy = deleter
	query := gr.data.DB()
	query = query.Where("studio_id = ?", studioID)
	query = query.Delete(&u)
	err := query.Error

	if err == nil && query.RowsAffected == 0 {
		return errors.New(codes.GameNotFound.String())
	}

	return err
}

func (gr *GameRepo) List(studioID uuid.UUID) (games []entity.Game, err error) {
	err = gr.data.DB().Where("studio_id = ?", studioID).Find(&games).Error
	return games, err
}

func (gr *GameRepo) Read(id uuid.UUID, studioID uuid.UUID) (game entity.Game, err error) {
	query := gr.data.DB()
	query = query.Preload("Studio")
	query = query.First(&game, "id = ? and studio_id = ?", id, studioID)
	err = query.Error

	return game, err
}
