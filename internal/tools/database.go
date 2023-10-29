package tools

import log "github.com/sirupsen/logrus"

// Database collection
type LoginDetails struct{
	AuthToken string
	Username string
}

type CoinDetails struct{
	Coins int64
	Username string
}

type DatabaseInterface interface{
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error){
	var database DatabaseInterface=&mockDB{}

// 	In Go, interfaces are typically satisfied using pointers to types to avoid unnecessary copying of values. This also allows the database variable to hold any value that implements the DatabaseInterface interface, whether it's a pointer to a struct or an actual struct, as long as the methods are implemented correctly.

// You can't directly pass mockDB to database because mockDB is a struct type, not an interface type. To assign it to an interface variable, you need to create an instance of it and use a pointer to that instance to satisfy the interface. This is a common pattern in Go when working with interfaces and their concrete implementations.

//Interface is a reference type in go

	var err error=database.SetupDatabase()
	if err!=nil{
		log.Error(err)
		return nil,err
	}

	return &database, nil
}