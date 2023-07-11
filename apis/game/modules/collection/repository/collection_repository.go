package collection_repository

import (
	"errors"

	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/enums/codes"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type CollectionRepo struct {
	data gbp.IDatastoreAdapter
}

func NewCollectionRepository(data gbp.IDatastoreAdapter) (r CollectionRepo) {
	return CollectionRepo{
		data: data,
	}
}

func (tr *CollectionRepo) Create(collection *entity.Collection) error {
	err := tr.data.DB().Omit(clause.Associations).Create(&collection).Error
	return err
}

func (tr *CollectionRepo) Update(collection *entity.Collection) error {
	query := tr.data.DB()
	query = query.Omit(clause.Associations)
	query = query.Save(&collection)
	err := query.Error
	return err
}

func (tr *CollectionRepo) Delete(collectionID, gameId, deleter uuid.UUID) error {
	u := entity.Collection{ID: collectionID}
	u.DeletedBy = deleter
	query := tr.data.DB()
	query = query.Where("game_id = ?", gameId)
	query = query.Delete(&u)
	err := query.Error

	if err == nil && query.RowsAffected == 0 {
		return errors.New(codes.CollectionNotFound.String())
	}

	return err
}

func (tr *CollectionRepo) List(gameId uuid.UUID) (collections []entity.Collection, err error) {
	err = tr.data.DB().Where("game_id = ?", gameId).Find(&collections).Error
	return collections, err
}

func (tr *CollectionRepo) Read(collectionID, gameId uuid.UUID) (collection entity.Collection, err error) {
	query := tr.data.DB()
	query = query.Preload("Game")
	query = query.First(&collection, "id = ? and game_id = ?", collectionID, gameId)
	err = query.Error

	if err == nil && query.RowsAffected == 0 {
		return collection, errors.New(codes.CollectionNotFound.String())
	}

	return collection, err
}

func (tr *CollectionRepo) CollectionBelongsToGame(collectionID, gameID uuid.UUID) (has bool, err error) {
	query := tr.data.DB()
	query = query.Raw(`
		select true 
		where exists (
			select cl.id from collections cl 
			where cl.id = ? 
			and cl.game_id = ?
			and cl.deleted_at is null)`,
		collectionID, gameID)
	query = query.Scan(&has)
	err = query.Error
	return has, err
}

func (tr *CollectionRepo) GameBelongsToStudio(gameID, studioID uuid.UUID) (has bool, err error) {
	query := tr.data.DB()
	query = query.Raw(`
		select true 
		where exists (
			select g.id from games g 
			where g.id = ? 
			and g.studio_id = ?
			and g.deleted_at is null)`,
		gameID, studioID)
	query = query.Scan(&has)
	err = query.Error
	return has, err
}

func (tr *CollectionRepo) SetCoverKey(collection *entity.Collection, key string) error {
	query := tr.data.DB()
	query = query.Omit(clause.Associations)
	query = query.Model(&collection)
	query = query.Update("cover_key", key)
	return query.Error
}

func (tr *CollectionRepo) SetAvatarKey(collection *entity.Collection, key string) error {
	query := tr.data.DB()
	query = query.Omit(clause.Associations)
	query = query.Model(&collection)
	query = query.Update("avatar_key", key)
	return query.Error
}

func (tr *CollectionRepo) CreateAttribute(attribute *entity.CollectionAttributes) error {
	err := tr.data.DB().Omit(clause.Associations).Create(&attribute).Error
	return err
}

func (tr *CollectionRepo) ReadAttribute(collectionID, attributeID uuid.UUID) (attribute entity.CollectionAttributes, err error) {
	query := tr.data.DB()
	query = query.Where("collection_id = ?", collectionID)
	query = query.Where("id = ?", attributeID)

	err = query.Find(&attribute).Error

	if err == nil && query.RowsAffected == 0 {
		return attribute, errors.New(codes.CollectionAttributeNotFound.String())
	}

	return attribute, err
}

func (tr *CollectionRepo) UpdateAttribute(attribute *entity.CollectionAttributes) error {
	query := tr.data.DB()
	query = query.Omit(clause.Associations)
	query = query.Save(&attribute)
	err := query.Error
	return err
}

func (tr *CollectionRepo) ListAttributes(collectionID uuid.UUID) (attributes []entity.CollectionAttributes, err error) {
	query := tr.data.DB()
	query = query.Where("collection_id = ?", collectionID)
	err = query.Find(&attributes).Error
	return attributes, err
}

func (tr *CollectionRepo) DeleteAttribute(attributeID, collectionID, deleter uuid.UUID) error {
	u := entity.CollectionAttributes{ID: attributeID, CollectionID: collectionID}
	u.DeletedBy = deleter
	query := tr.data.DB()
	query = query.Delete(&u)
	err := query.Error

	if err == nil && query.RowsAffected == 0 {
		return errors.New(codes.CollectionNotFound.String())
	}

	return err
}
