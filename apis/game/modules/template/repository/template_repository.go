package template_repository

import (
	"errors"

	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/enums/codes"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type TemplateRepo struct {
	data gbp.IDatastoreAdapter
}

func NewTemplateRepository(data gbp.IDatastoreAdapter) (r TemplateRepo) {
	return TemplateRepo{
		data: data,
	}
}

func (tr *TemplateRepo) Create(template *entity.Template) error {
	err := tr.data.DB().Omit(clause.Associations).Create(&template).Error
	return err
}

func (tr *TemplateRepo) Update(template *entity.Template) error {
	query := tr.data.DB()
	query = query.Omit(clause.Associations)
	query = query.Save(&template)
	err := query.Error
	return err
}

func (tr *TemplateRepo) Delete(templateID uuid.UUID, gameId uuid.UUID, deleter uuid.UUID) error {
	u := entity.Template{ID: templateID}
	u.DeletedBy = deleter
	query := tr.data.DB()
	query = query.Where("game_id = ?", gameId)
	query = query.Delete(&u)
	err := query.Error

	if err == nil && query.RowsAffected == 0 {
		return errors.New(codes.TemplateNotFound.String())
	}

	return err
}

func (tr *TemplateRepo) List(gameId uuid.UUID) (templates []entity.Template, err error) {
	err = tr.data.DB().Where("game_id = ?", gameId).Find(&templates).Error
	return templates, err
}

func (tr *TemplateRepo) Read(templateId uuid.UUID, gameId uuid.UUID) (template entity.Template, err error) {
	query := tr.data.DB()
	query = query.Preload("Game")
	query = query.First(&template, "id = ? and game_id = ?", templateId, gameId)
	err = query.Error

	if err == nil && query.RowsAffected == 0 {
		return template, errors.New(codes.TemplateNotFound.String())
	}

	return template, err
}

func (tr *TemplateRepo) GameBelongsToStudio(gameID uuid.UUID, studioID uuid.UUID) (has bool, err error) {
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

func (tr *TemplateRepo) TemplateBelongsToGame(templateID uuid.UUID, gameID uuid.UUID) (has bool, err error) {
	query := tr.data.DB()
	query = query.Raw(`
		select true 
		where exists (
			select tp.id from templates tp 
			where tp.id = ? 
			and tp.game_id = ?
			and tp.deleted_at is null)`,
		templateID, gameID)
	query = query.Scan(&has)
	err = query.Error
	return has, err
}

func (tr *TemplateRepo) CreateAttribute(attribute *entity.TemplateAttributes) error {
	err := tr.data.DB().Omit(clause.Associations).Create(&attribute).Error
	return err
}

func (tr *TemplateRepo) UpdateAttribute(attribute *entity.TemplateAttributes) error {
	query := tr.data.DB()
	query = query.Omit(clause.Associations)
	query = query.Save(&attribute)
	err := query.Error
	return err
}

func (tr *TemplateRepo) ListAttributes(templateID uuid.UUID) (attributes []entity.TemplateAttributes, err error) {
	query := tr.data.DB()
	query = query.Where("template_id = ?", templateID)
	err = query.Find(&attributes).Error
	return attributes, err
}

func (tr *TemplateRepo) ReadAttribute(templateID, attributeID uuid.UUID) (attribute entity.TemplateAttributes, err error) {
	query := tr.data.DB()
	query = query.Where("template_id = ?", templateID)
	query = query.Where("id = ?", attributeID)

	err = query.Find(&attribute).Error

	if err == nil && query.RowsAffected == 0 {
		return attribute, errors.New(codes.TemplateAttributeNotFound.String())
	}

	return attribute, err
}

func (tr *TemplateRepo) DeleteAttribute(attributeID, templateID, deleter uuid.UUID) error {
	u := entity.TemplateAttributes{ID: attributeID, TemplateID: templateID}
	u.DeletedBy = deleter
	query := tr.data.DB()
	query = query.Delete(&u)
	err := query.Error

	if err == nil && query.RowsAffected == 0 {
		return errors.New(codes.TemplateNotFound.String())
	}

	return err
}

func (tr *TemplateRepo) FirstFileByType(templateID uuid.UUID, filetype datatypes.FileType) (file entity.TemplateFiles, err error) {
	query := tr.data.DB()
	query = query.Where(&entity.TemplateFiles{TemplateID: templateID, Type: filetype})
	query = query.First(&file)
	err = query.Error
	return file, err
}

func (tr *TemplateRepo) SaveFile(file *entity.TemplateFiles) error {
	query := tr.data.DB()
	query = query.Omit(clause.Associations)
	query.Save(file)
	return query.Error
}

func (tr *TemplateRepo) ListFilesIgnoringType(templateID uuid.UUID, filetypeToIgnore datatypes.FileType) (files []entity.TemplateFiles, err error) {
	query := tr.data.DB()
	query = query.Where("template_id = ?", templateID)
	query = query.Where("type != ?", filetypeToIgnore)
	err = query.Find(&files).Error
	return files, err
}

func (tr *TemplateRepo) GetFile(fileID, templateID uuid.UUID) (file entity.TemplateFiles, err error) {
	query := tr.data.DB()
	query = query.Where(entity.TemplateFiles{ID: fileID, TemplateID: templateID})
	query = query.First(&file)
	err = query.Error
	return file, err
}

func (tr *TemplateRepo) DeleteFile(fileID, templateID uuid.UUID) error {
	u := entity.TemplateFiles{ID: fileID, TemplateID: templateID}
	query := tr.data.DB()
	query = query.Delete(&u)
	err := query.Error

	if err == nil && query.RowsAffected == 0 {
		return errors.New(codes.TemplateNotFound.String())
	}

	return err
}
