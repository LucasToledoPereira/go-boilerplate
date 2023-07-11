package game_handler

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	game_commands "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/game/commands"
	interfaces "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/game/interfaces"
	"github.com/google/uuid"
)

type GameHandler struct {
	repository interfaces.IGameRepository
	filestore  gbp.IFilestoreAdapter
}

func NewGameHandler(repo interfaces.IGameRepository, filestore gbp.IFilestoreAdapter) (handler GameHandler) {
	return GameHandler{
		repository: repo,
		filestore:  filestore,
	}
}

func (h GameHandler) Create(command game_commands.CreateGameCommandRequest) (res game_commands.CreateGameCommandResponse, err error) {
	game := command.ToEntity()
	err = h.repository.Create(&game)
	res.FromEntity(&game)
	return res, err
}

func (h GameHandler) List(studioID uuid.UUID) (games []game_commands.ListGameCommandResponse, err error) {
	entities, err := h.repository.List(studioID)
	games = game_commands.ListResponseFromEntities(entities)
	return games, err
}

func (h GameHandler) Read(id uuid.UUID, studioID uuid.UUID) (res game_commands.ReadGameCommandResponse, err error) {
	game, err := h.repository.Read(id, studioID)
	res.FromEntity(&game)
	return res, err
}

func (h GameHandler) Update(command game_commands.UpdateGameCommandRequest) (res game_commands.UpdateGameCommandResponse, err error) {
	game, err := h.repository.Read(command.GameID, command.StudioID)
	if err != nil {
		return game_commands.UpdateGameCommandResponse{}, err
	}

	command.ToEntity(&game)
	err = h.repository.Update(&game)
	res.FromEntity(&game)
	return res, err
}

func (h GameHandler) Delete(userID uuid.UUID, studioID uuid.UUID, deleter uuid.UUID) (err error) {
	err = h.repository.Delete(userID, studioID, deleter)
	return err
}
