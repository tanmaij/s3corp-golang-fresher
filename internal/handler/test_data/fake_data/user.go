package fake_data

import (
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/utils"
	"time"
)

var UserLogin = models.User{Username: "mai", Password: "1", Email: "mai@gmail.com", Name: "Mãi", CreatedAt: time.Date(2022, 6, 29, 0, 0, 0, 0, time.UTC)}
var TokenLogin = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1haSJ9.qYNphS_Xycc7-XY9MD9o_kTHocUjV6kCH0hD1EzTDk4"

var UserSlice = []models.User{models.User{Username: "mai", Email: "mai@gmail.com", Name: "Mãi", CreatedAt: time.Date(2022, 6, 29, 0, 0, 0, 0, time.UTC)}, models.User{Username: "thien", Email: "thien@gmail.com", Name: "thien", CreatedAt: time.Date(2022, 6, 29, 0, 0, 0, 0, time.UTC)}}
var Pagination = utils.Pagination{TotalPages: 3, Limit: 2, Page: 2, TotalRows: 5}
