package collection_interfaces

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

type ICollectionRepository interface {
	Create(collection *entity.Collection) error
	Update(collection *entity.Collection) error
	Delete(collectionID, gameId, deleter uuid.UUID) error
	List(gameId uuid.UUID) (collections []entity.Collection, err error)
	Read(collectionID, gameId uuid.UUID) (collection entity.Collection, err error)
	CollectionBelongsToGame(collectionID, gameID uuid.UUID) (has bool, err error)
	GameBelongsToStudio(gameID, StudioID uuid.UUID) (has bool, err error)
	SetCoverKey(collection *entity.Collection, key string) error
	SetAvatarKey(collection *entity.Collection, key string) error
	CreateAttribute(attribute *entity.CollectionAttributes) error
	ReadAttribute(collectionID, attributeID uuid.UUID) (attribute entity.CollectionAttributes, err error)
	UpdateAttribute(attribute *entity.CollectionAttributes) error
	ListAttributes(collectionID uuid.UUID) (attributes []entity.CollectionAttributes, err error)
	DeleteAttribute(attributeID, collectionID, deleter uuid.UUID) error
}
