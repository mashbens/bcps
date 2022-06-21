package migrate

import (
	user "github.com/mashbens/cps/repository/user"
)

func AutoMigrate() {
	var user = user.User{}
	_ = user
}
