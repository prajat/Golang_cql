package Database

import (
	"log"

	"github.com/gocql/gocql"
)

type DBConnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var connection DBConnection

func SetupDBConnection() {
	connection.cluster = gocql.NewCluster("127.0.0.1")
	connection.cluster.Consistency = gocql.Quorum
	connection.cluster.Keyspace = "golang_cql"
	connection.session, _ = connection.cluster.CreateSession()

}

func ExecuteQuery(query string, values ...interface{}) {

	if err := connection.session.Query(query).Bind(values...).Exec(); err != nil {
		log.Fatal(err)
	}

}
