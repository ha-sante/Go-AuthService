package database

import (
	"fmt"
	"log"

	faunadb "github.com/fauna/faunadb-go/faunadb"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDB used to connect auth service to faunadb
func ConnectDB(){

	connection, err := faunadb.NewFaunaClient("fnADod0i9aACFDHEKDvmsLUdgbCMt-F9Afh9jAaw")

	if err != nil {

		fmt.Println("auth_service_DB: CONNECTION FAILED %v", err)
		logFatal(err)
	}

	fmt.Println("auth_service_DB: CONNECTION SUCCESSFULL" )

	return connection
}
