package studion_interfaces

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

type IStudioRepository interface {
	Create(studio *entity.Studio) error
	Update(studio *entity.Studio) error
	Delete(id uuid.UUID) error
	List() (studios []entity.Studio, err error)
	Read(id uuid.UUID) (studio entity.Studio, err error)
	UpdateImage(user entity.Studio, filekey string) error
}
