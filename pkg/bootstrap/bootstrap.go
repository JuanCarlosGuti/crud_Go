package bootstrap

import (
	"github.com/JuanCarlosGuti/Go_users.git/internal/domain"
	"github.com/JuanCarlosGuti/Go_users.git/internal/user"
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewDB() user.DB {
	return user.DB{
		Users: []domain.User{{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@gmail.com",
		}, {
			ID:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane.doe@gmail.com",
		}, {
			ID:        3,
			FirstName: "Janet",
			LastName:  "Does",
			Email:     "janet.does@gmail.com",
		},
		},
		MaxUserID: 3,
	}

}
