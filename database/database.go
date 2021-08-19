package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Defining our global variable and pointer to our connection
var (
	DBConn *gorm.DB
)
/* We are able to use our orm to talk to our database */

