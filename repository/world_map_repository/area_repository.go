package world_map_repository

import (
	"context"
	"fmt"
	"math"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// DAO
type Area struct {
	Uuid           string     `mapstructure:"uuid" json:"uuid" example:"fb621679-6284-4aac-b70f-148bd3c8e1d2"`
	Coordinate     Coordinate `mapstructure:"coordinate" json:"coordinate"`
	CoordinateName string     `mapstructure:"coordinateName" json:"coordinateName" example:"(1, 1)"`
	Radius         int64      `mapstructure:"radius" json:"radius" example:"1"`
	Region         Region     `mapstructure:"region" json:"region"`
}

type Coordinate struct {
	X int64 `mapstructure:"x" json:"x" example:"1"`
	Y int64 `mapstructure:"y" json:"y" example:"1"`
}

type AreaRepository interface {
	CreateArea(coordinate *Coordinate) (*string, error)
	GetAreas(radius *int64) ([]Area, error)
	GetArea(uuid string) (Area, error)
}

type AreaRepositoryImpl struct {
	Driver *neo4j.DriverWithContext
}

func (a AreaRepositoryImpl) GetAreas(radius *int64) ([]Area, error) {
	ctx := context.Background()

	radiusFilter := ""
	if radius != nil {
		radiusFilter = "WHERE a.radius = $radius"
	}

	cypher := "MATCH (a:Area) --> (r:Region)"
	cypher = fmt.Sprintln(cypher, radiusFilter)
	cypher = fmt.Sprintln(cypher, "RETURN a, r")

	parameters := map[string]interface{}{
		"radius": radius,
	}

	result, err := neo4j.ExecuteQuery(ctx, *a.Driver, cypher, parameters, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return []Area{}, err
	}

	areas := make([]Area, 0, len(result.Records))

	for _, record := range result.Records {
		areaNode, areaNil, err := neo4j.GetRecordValue[neo4j.Node](record, "a")
		if err != nil {
			return []Area{}, err
		}
		regionNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "r")
		if err != nil {
			return []Area{}, err
		}

		if areaNil {
			return []Area{}, nil
		}

		_uuid, err := neo4j.GetProperty[string](areaNode, "uuid")
		if err != nil {
			return []Area{}, err
		}
		coordX, err := neo4j.GetProperty[int64](areaNode, "coordinateX")
		if err != nil {
			return []Area{}, err
		}
		coordY, err := neo4j.GetProperty[int64](areaNode, "coordinateY")
		if err != nil {
			return []Area{}, err
		}
		coordName, err := neo4j.GetProperty[string](areaNode, "coordinateName")
		if err != nil {
			return []Area{}, err
		}
		radius, err := neo4j.GetProperty[int64](areaNode, "radius")
		if err != nil {
			return []Area{}, err
		}
		regionName, err := neo4j.GetProperty[string](regionNode, "name")
		if err != nil {
			return []Area{}, err
		}
		regionCode, err := neo4j.GetProperty[string](regionNode, "code")
		if err != nil {
			return []Area{}, err
		}

		area := Area{
			Uuid: _uuid,
			Coordinate: Coordinate{
				X: coordX,
				Y: coordY,
			},
			CoordinateName: coordName,
			Radius:         radius,
			Region: Region{
				Name: regionName,
				Code: regionCode,
			},
		}

		areas = append(areas, area)
	}

	return areas, nil
}

func (a AreaRepositoryImpl) GetArea(uuid string) (Area, error) {
	ctx := context.Background()

	cypher := `
			MATCH (a:Area {uuid: $uuid})
			-->
			(r:Region)
			RETURN a, r
			`

	parameters := map[string]interface{}{
		"uuid": uuid,
	}

	result, err := neo4j.ExecuteQuery(ctx, *a.Driver, cypher, parameters, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return Area{}, err
	}

	if len(result.Records) == 0 {
		return Area{}, nil
	}

	areaNode, areaNil, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "a")
	if err != nil {
		return Area{}, err
	}
	regionNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "r")
	if err != nil {
		return Area{}, err
	}

	if areaNil {
		return Area{}, nil
	}

	_uuid, err := neo4j.GetProperty[string](areaNode, "uuid")
	if err != nil {
		return Area{}, err
	}
	coordX, err := neo4j.GetProperty[int64](areaNode, "coordinateX")
	if err != nil {
		return Area{}, err
	}
	coordY, err := neo4j.GetProperty[int64](areaNode, "coordinateY")
	if err != nil {
		return Area{}, err
	}
	coordName, err := neo4j.GetProperty[string](areaNode, "coordinateName")
	if err != nil {
		return Area{}, err
	}
	radius, err := neo4j.GetProperty[int64](areaNode, "radius")
	if err != nil {
		return Area{}, err
	}
	regionName, err := neo4j.GetProperty[string](regionNode, "name")
	if err != nil {
		return Area{}, err
	}
	regionCode, err := neo4j.GetProperty[string](regionNode, "code")
	if err != nil {
		return Area{}, err
	}

	area := Area{
		Uuid: _uuid,
		Coordinate: Coordinate{
			X: coordX,
			Y: coordY,
		},
		CoordinateName: coordName,
		Radius:         radius,
		Region: Region{
			Name: regionName,
			Code: regionCode,
		},
	}

	return area, nil
}

func (a AreaRepositoryImpl) CreateArea(coordinate *Coordinate) (*string, error) {
	ctx := context.Background()
	session := (*a.Driver).NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: "neo4j",
		AccessMode:   neo4j.AccessModeWrite,
	})

	defer session.Close(ctx)

	cypher := `
			MERGE (a:Area {
				coordinateX: $coordinateX,
				coordinateY: $coordinateY,
				coordinateName: $coordinateName
			})
			MERGE (r:Region {
				name: $regionName,
				code: $regionCode
			})
			MERGE (a)-[:REGION_OF]->(r)
			ON CREATE
				SET
					a.uuid = randomUUID(),
					a.radius = $radius
			RETURN a.uuid
			`

	radius := GetRadiusByCoordinate(*coordinate)
	region := GetRegionByCoordinate(*coordinate)

	parameters := map[string]interface{}{
		"coordinateName": GetCoordinateNameByXY(coordinate.X, coordinate.Y),
		"coordinateX":    coordinate.X,
		"coordinateY":    coordinate.Y,
		"radius":         radius,
		"regionName":     region.Name,
		"regionCode":     region.Code,
	}

	newAreaUuid, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, cypher, parameters)
			if err != nil {
				return nil, err
			}
			record, err := result.Single(ctx)
			if err != nil {
				return nil, err
			}

			uuid, found := record.Get("a.uuid")
			if !found {
				return nil, nil
			}

			return uuid, nil
		})
	if err != nil {
		return nil, err
	}

	newAreaUuidStr := newAreaUuid.(string)

	return &newAreaUuidStr, nil
}

func GetCoordinateNameByXY(x int64, y int64) string {
	return fmt.Sprintf("(%d, %d)", x, y)
}

func GetRadiusByCoordinate(coordinate Coordinate) int64 {
	absX := math.Abs(float64(coordinate.X))
	absY := math.Abs(float64(coordinate.Y))

	if absX > absY {
		return int64(absX)
	}
	return int64(absY)
}

func GetAllRegions() []Region {
	return []Region{
		{Name: "Northeast", Code: "NE"},
		{Name: "Southeast", Code: "SE"},
		{Name: "Southwest", Code: "SW"},
		{Name: "Northwest", Code: "NW"},
	}
}

func GetRegionByCoordinate(coordinate Coordinate) Region {
	if coordinate.X > 0 {
		if coordinate.Y > 0 {
			return GetAllRegions()[0]
		}
		return GetAllRegions()[1]
	} else {
		if coordinate.Y < 0 {
			return GetAllRegions()[2]
		}
		return GetAllRegions()[3]
	}
}

func AreaRepositoryInit(driver *neo4j.DriverWithContext) *AreaRepositoryImpl {
	return &AreaRepositoryImpl{
		Driver: driver,
	}
}
