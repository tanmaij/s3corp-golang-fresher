package fake_data

import (
	"s3corp-golang-fresher/internal/models"
	"time"
)

var timeLoc = time.UTC
var UserSliceByYear = models.UserSlice{
	&models.User{
		Username:  "mai",
		Email:     "mai@gmail.com",
		Name:      "Mãi",
		CreatedAt: time.Date(2022, 6, 01, 0, 0, 0, 0, timeLoc)},
	&models.User{
		Username:  "thien",
		Email:     "thien@gmail.com",
		Name:      "thien",
		CreatedAt: time.Date(2022, 6, 02, 0, 0, 0, 0, timeLoc),
	},
	&models.User{
		Username:  "mai",
		Email:     "mai@gmail.com",
		Name:      "Mãi",
		CreatedAt: time.Date(2022, 6, 03, 0, 0, 0, 0, timeLoc),
	},
	&models.User{
		Username:  "thien",
		Email:     "thien@gmail.com",
		Name:      "thien",
		CreatedAt: time.Date(2022, 6, 04, 0, 0, 0, 0, timeLoc),
	},
	&models.User{
		Username:  "mai",
		Email:     "mai@gmail.com",
		Name:      "Mãi",
		CreatedAt: time.Date(2022, 6, 05, 0, 0, 0, 0, timeLoc),
	},
	&models.User{
		Username:  "thien",
		Email:     "thien@gmail.com",
		Name:      "thien",
		CreatedAt: time.Date(2022, 6, 06, 0, 0, 0, 0, timeLoc),
	},
	&models.User{
		Username:  "mai",
		Email:     "mai@gmail.com",
		Name:      "Mãi",
		CreatedAt: time.Date(2022, 6, 07, 0, 0, 0, 0, timeLoc),
	},
	&models.User{
		Username:  "thien",
		Email:     "thien@gmail.com",
		Name:      "thien",
		CreatedAt: time.Date(2022, 6, 8, 0, 0, 0, 0, timeLoc),
	},
	&models.User{
		Username:  "mai",
		Email:     "mai@gmail.com",
		Name:      "Mãi",
		CreatedAt: time.Date(2022, 6, 9, 0, 0, 0, 0, timeLoc),
	},
	&models.User{
		Username:  "thien",
		Email:     "thien@gmail.com",
		Name:      "thien",
		CreatedAt: time.Date(2022, 6, 10, 0, 0, 0, 0, timeLoc),
	},
}
