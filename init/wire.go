//go:build wireinject
// +build wireinject

package initialization

import (
	"github.com/google/wire"
	"github.com/june21x/project-a-api/api/controller"
	"github.com/june21x/project-a-api/internal_service/player_service"
	"github.com/june21x/project-a-api/internal_service/world_map_service"
	"github.com/june21x/project-a-api/repository/player_repository"
	"github.com/june21x/project-a-api/repository/world_map_repository"
)

var db = wire.NewSet(ConnectToDB)

var healthCheckCtrlSet = wire.NewSet(controller.HealthCheckControllerInit,
	wire.Bind(new(controller.HealthCheckController), new(*controller.HealthCheckControllerImpl)),
)

var playerCtrlSet = wire.NewSet(controller.PlayerControllerInit,
	wire.Bind(new(controller.PlayerController), new(*controller.PlayerControllerImpl)),
)

var areaCtrlSet = wire.NewSet(controller.AreaControllerInit,
	wire.Bind(new(controller.AreaController), new(*controller.AreaControllerImpl)),
)

var playerServiceSet = wire.NewSet(player_service.PlayerServiceInit,
	wire.Bind(new(player_service.PlayerService), new(*player_service.PlayerServiceImpl)),
)

var areaServiceSet = wire.NewSet(world_map_service.AreaServiceInit,
	wire.Bind(new(world_map_service.AreaService), new(*world_map_service.AreaServiceImpl)),
)

var playerRepoSet = wire.NewSet(player_repository.PlayerRepositoryInit,
	wire.Bind(new(player_repository.PlayerRepository), new(*player_repository.PlayerRepositoryImpl)),
)

var areaRepoSet = wire.NewSet(world_map_repository.AreaRepositoryInit,
	wire.Bind(new(world_map_repository.AreaRepository), new(*world_map_repository.AreaRepositoryImpl)),
)

func Initialize() *Initialization {
	wire.Build(NewInitialization, db, healthCheckCtrlSet, playerCtrlSet, areaCtrlSet, playerServiceSet, areaServiceSet, playerRepoSet, areaRepoSet)
	return nil
}
