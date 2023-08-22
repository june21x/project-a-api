package player_service

import (
	"github.com/june21x/project-a-api/repository/player_repository"
)

type PlayerService interface {
	RegisterPlayer(player *player_repository.Player, password string) (player_repository.Player, error)
	GetPlayers() ([]player_repository.Player, error)
	GetPlayer(uuid string) (player_repository.Player, error)
	// UpdatePlayer()
	// DeletePlayer()
}

type PlayerServiceImpl struct {
	playerRepository player_repository.PlayerRepository
}

// Register player and return player
func (p PlayerServiceImpl) RegisterPlayer(player *player_repository.Player, password string) (player_repository.Player, error) {
	// TODO validate data uniqueness (no duplicated playerName)

	// insert player
	uuid, err := p.playerRepository.CreatePlayer(player, password)
	if err != nil || uuid == nil {
		return player_repository.Player{}, err
	}

	// get player by uuid
	createdPlayer, err := p.playerRepository.GetPlayer(*uuid)
	if err != nil {
		return player_repository.Player{}, err
	}

	return createdPlayer, nil
}

func (p PlayerServiceImpl) GetPlayers() ([]player_repository.Player, error) {
	players, err := p.playerRepository.GetPlayers()
	if err != nil {
		return []player_repository.Player{}, err
	}

	return players, nil
}

func (p PlayerServiceImpl) GetPlayer(uuid string) (player_repository.Player, error) {
	player, err := p.playerRepository.GetPlayer(uuid)
	if err != nil {
		return player_repository.Player{}, err
	}

	return player, nil
}

func PlayerServiceInit(playerRepository player_repository.PlayerRepository) *PlayerServiceImpl {
	return &PlayerServiceImpl{
		playerRepository: playerRepository,
	}
}
