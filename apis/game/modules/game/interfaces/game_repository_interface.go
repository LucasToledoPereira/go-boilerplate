package game_interfaces

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

type IGameRepository interface {
	Create(game *entity.Game) error
	Update(game *entity.Game) error
	Delete(game uuid.UUID, studioID uuid.UUID, deleter uuid.UUID) error
	List(studioID uuid.UUID) (games []entity.Game, err error)
	Read(id uuid.UUID, studioID uuid.UUID) (game entity.Game, err error)
}
