package Database

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

type DBConnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var connection DBConnection
var name string

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

func ExecuteQuerynew(query string) string {

	if err := connection.session.Query(query, "wavy@gmail.com").Consistency(gocql.One).Scan(&name); err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
	return name

}
