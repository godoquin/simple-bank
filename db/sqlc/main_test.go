package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/godoquin/simple-bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load configuration: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

/* func TestLoadConfig(t *testing.T) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, config)
}
*/
