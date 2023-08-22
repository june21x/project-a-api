package player_repository

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"golang.org/x/crypto/bcrypt"
)

// DAO
type Player struct {
	Uuid       string `mapstructure:"uuid" json:"uuid" example:"fb621679-6284-4aac-b70f-148bd3c8e1d2"`
	PlayerName string `mapstructure:"playerName" json:"playerName" example:"June Eleutheria"`
	Email      string `mapstructure:"email" json:"email" example:"juneeleutheria@gmail.com"`
}

type PlayerRepository interface {
	CreatePlayer(player *Player, password string) (*string, error)
	GetPlayers() ([]Player, error)
	GetPlayer(uuid string) (Player, error)
}

type PlayerRepositoryImpl struct {
	Driver *neo4j.DriverWithContext
}

func (a PlayerRepositoryImpl) GetPlayers() ([]Player, error) {
	ctx := context.Background()

	cypher := "MATCH (p:Player)"
	cypher = fmt.Sprintln(cypher, "RETURN p")

	result, err := neo4j.ExecuteQuery(ctx, *a.Driver, cypher, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return []Player{}, err
	}

	players := make([]Player, 0, len(result.Records))

	for _, record := range result.Records {
		playerNode, playerNil, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
		if err != nil {
			return []Player{}, err
		}

		if playerNil {
			return []Player{}, nil
		}

		player := Player{}

		err = mapstructure.Decode(playerNode.Props, &player)
		if err != nil {
			return []Player{}, err
		}

		players = append(players, player)
	}

	return players, nil
}

func (u PlayerRepositoryImpl) GetPlayer(uuid string) (Player, error) {
	ctx := context.Background()

	cypher := `
			MATCH (p:Player {uuid: $uuid})
			RETURN p
			`

	parameters := map[string]interface{}{
		"uuid": uuid,
	}

	result, err := neo4j.ExecuteQuery(ctx, *u.Driver, cypher, parameters, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return Player{}, err
	}

	if len(result.Records) == 0 {
		return Player{}, nil
	}

	playerNode, playerNil, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "p")
	if err != nil {
		return Player{}, err
	}

	if playerNil {
		return Player{}, nil
	}

	player := Player{}

	err = mapstructure.Decode(playerNode.Props, &player)
	if err != nil {
		return Player{}, err
	}

	return player, nil
}

func (p PlayerRepositoryImpl) CreatePlayer(player *Player, password string) (*string, error) {
	ctx := context.Background()
	session := (*p.Driver).NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: "neo4j",
		AccessMode:   neo4j.AccessModeWrite,
	})

	defer session.Close(ctx)

	cypher := `
			CREATE (p:Player {uuid: randomUUID(), email: $email, playerName: $playerName, password: $password})
			RETURN p.uuid
			`
	hashedPassword, err := hash(password)
	if err != nil {
		return nil, err
	}
	parameters := map[string]interface{}{
		"email":      player.Email,
		"playerName": player.PlayerName,
		"password":   hashedPassword,
	}

	newPlayerUuid, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, cypher, parameters)
			if err != nil {
				return nil, err
			}
			record, err := result.Single(ctx)
			if err != nil {
				return nil, err
			}

			uuid, found := record.Get("p.uuid")
			if !found {
				return nil, nil
			}

			return uuid, nil
		})
	if err != nil {
		return nil, err
	}

	newPlayerUuidStr := newPlayerUuid.(string)

	return &newPlayerUuidStr, nil
}

func hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func PlayerRepositoryInit(driver *neo4j.DriverWithContext) *PlayerRepositoryImpl {
	return &PlayerRepositoryImpl{
		Driver: driver,
	}
}
