package db

import (
	"log"
	"sync"

	"github.com/cottrellio/cottrellio_go/pkg/model"
)

// DB defines the expected db methods.
type DB interface {
	Connect() error
	// User
	UserCreate(model.User) (*model.User, error)
	UserList(map[string][]string, map[string]string) ([]*model.User, int64, error)
	UserDetail(string) (*model.User, error)
	UserUpdate(string, model.User) (*model.User, error)
	UserDelete(string) error
}

const (
	_ = iota
	// MONGODB signals to use the Mongo driver.
	MONGODB
)

var db DB          // will be used as a singleton db object
var once sync.Once // used for thread safety

// New instantiates a db singleton object.
func New(driver int) (DB, error) {
	var err error
	once.Do(func() {
		db = databaseFactory(driver)
		err = db.Connect()
	})
	return db, err
}

func databaseFactory(driver int) DB {
	switch driver {
	case MONGODB:
		return new(MongoDB)
	default:
		log.Fatal("Unsupported DB driver.")
		return nil
	}
}
