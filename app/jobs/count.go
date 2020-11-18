package jobs

import (
	"fmt"
	"github.com/mr-klaatu/baseapp/app/controllers"
	"github.com/mr-klaatu/baseapp/app/models"
	"github.com/revel/modules/jobs/app/jobs"
	"github.com/revel/revel"
)

// Periodically count the users in the database.
type UserCounter struct{}

func (c UserCounter) Run() {
	users, err := controllers.Dbm.Select(&models.User{},
		`select * from User`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %d users.\n", len(users))
}

func init() {
	revel.OnAppStart(func() {
		jobs.Schedule("@every 1m", UserCounter{})
	})
}
