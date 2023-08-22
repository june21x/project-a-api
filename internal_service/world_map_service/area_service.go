package world_map_service

import (
	"github.com/june21x/project-a-api/repository/world_map_repository"
)

type AreaService interface {
	// GetAllArea(c *gin.Context)
	// AddAreaData(c *gin.Context)
	// UpdateAreaData(c *gin.Context)
	// DeleteArea(c *gin.Context)
	CreateArea(coordinate *world_map_repository.Coordinate) (world_map_repository.Area, error)
	GetAreas(radius *int64) ([]world_map_repository.Area, error)
	GetArea(uuid string) (world_map_repository.Area, error)
}

type AreaServiceImpl struct {
	areaRepository world_map_repository.AreaRepository
}

func (a AreaServiceImpl) CreateArea(coordinate *world_map_repository.Coordinate) (world_map_repository.Area, error) {
	// TODO validate data uniqueness (no duplicated coordinate)

	// create area
	uuid, err := a.areaRepository.CreateArea(coordinate)
	if err != nil || uuid == nil {
		return world_map_repository.Area{}, err
	}

	// get area by uuid
	createdArea, err := a.areaRepository.GetArea(*uuid)
	if err != nil {
		return world_map_repository.Area{}, err
	}

	return createdArea, nil
}

func (a AreaServiceImpl) GetAreas(radius *int64) ([]world_map_repository.Area, error) {
	areas, err := a.areaRepository.GetAreas(radius)
	if err != nil {
		return []world_map_repository.Area{}, err
	}

	return areas, nil
}

func (a AreaServiceImpl) GetArea(uuid string) (world_map_repository.Area, error) {
	area, err := a.areaRepository.GetArea(uuid)
	if err != nil {
		return world_map_repository.Area{}, err
	}

	return area, nil
}

func AreaServiceInit(areaRepository world_map_repository.AreaRepository) *AreaServiceImpl {
	return &AreaServiceImpl{
		areaRepository: areaRepository,
	}
}
