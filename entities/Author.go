package author

import (
	"time"
)

type Author struct {
	id       int
	name     string
	username string
	password string
	cdate    time.Time
	udate    time.Time
}
