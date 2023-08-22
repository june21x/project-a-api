package initialization

import (
	"github.com/june21x/project-a-api/api/controller"
	"github.com/june21x/project-a-api/internal_service/player_service"
	"github.com/june21x/project-a-api/internal_service/world_map_service"
	"github.com/june21x/project-a-api/repository/player_repository"
	"github.com/june21x/project-a-api/repository/world_map_repository"
)

type Initialization struct {
	playerRepo      player_repository.PlayerRepository
	areaRepo        world_map_repository.AreaRepository
	playerSvc       player_service.PlayerService
	areaSvc         world_map_service.AreaService
	HealthCheckCtrl controller.HealthCheckController
	PlayerCtrl      controller.PlayerController
	AreaCtrl        controller.AreaController
}

func NewInitialization(playerRepo player_repository.PlayerRepository,
	areaRepo world_map_repository.AreaRepository,
	playerService player_service.PlayerService,
	areaService world_map_service.AreaService,
	healthCheckCtrl controller.HealthCheckController,
	playerCtrl controller.PlayerController,
	areaCtrl controller.AreaController,
) *Initialization {
	return &Initialization{
		playerRepo:      playerRepo,
		areaRepo:        areaRepo,
		playerSvc:       playerService,
		areaSvc:         areaService,
		HealthCheckCtrl: healthCheckCtrl,
		PlayerCtrl:      playerCtrl,
		AreaCtrl:        areaCtrl,
	}
}
