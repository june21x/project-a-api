package initialization

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/viper"
)

func ConnectToDB() *neo4j.DriverWithContext {
	ctx := context.Background()

	uri := viper.GetString("neo_4j.uri")
	username := viper.GetString("neo_4j.username")
	password := viper.GetString("neo_4j.password")

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic(err)
	}
	// defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	return &driver
}
