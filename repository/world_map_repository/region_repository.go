package world_map_repository

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// DAO
type Region struct {
	Name string `mapstructure:"name" json:"name" example:"Northeast"`
	Code string `mapstructure:"code" json:"code" example:"NE"`
}

type RegionRepository interface {
	GetRegion(uuid string) (Region, error)
	GetRegions() ([]Region, error)
}

type RegionRepositoryImpl struct {
	Driver *neo4j.DriverWithContext
}

func (a RegionRepositoryImpl) GetRegion(code string) (Region, error) {
	ctx := context.Background()

	cypher := `
			MATCH (r:Region {code: $code})
			RETURN r
			`

	parameters := map[string]interface{}{
		"code": code,
	}

	region := Region{}

	result, err := neo4j.ExecuteQuery(ctx, *a.Driver, cypher, parameters, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return Region{}, err
	}

	if len(result.Records) == 0 {
		return Region{}, nil
	}

	value, found := result.Records[0].Get("r")
	if !found {
		return Region{}, nil
	}

	node := value.(neo4j.Node)
	props := node.Props

	err = mapstructure.Decode(props, &region)
	if err != nil {
		return Region{}, err
	}

	return region, nil
}

func (a RegionRepositoryImpl) GetRegions() ([]Region, error) {
	ctx := context.Background()

	cypher := `
			MATCH (r:Region)
			RETURN r
			`

	result, err := neo4j.ExecuteQuery(ctx, *a.Driver, cypher, nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		return nil, err
	}

	regions := make([]Region, 0, len(result.Records))

	// Loop through results and do something with them
	for _, record := range result.Records {
		name, _, _ := neo4j.GetRecordValue[string](record, "name")
		code, _, _ := neo4j.GetRecordValue[string](record, "code")

		regions = append(regions, Region{Name: name, Code: code})
	}

	// Summary information
	fmt.Printf("The query `%v` returned %v records in %+v.\n",
		result.Summary.Query().Text(), len(result.Records),
		result.Summary.ResultAvailableAfter())

	return regions, nil
}

func RegionRepositoryInit(driver *neo4j.DriverWithContext) *RegionRepositoryImpl {
	return &RegionRepositoryImpl{
		Driver: driver,
	}
}
