package repository

import (
	// "fmt"
	"InstagramMock-v2/model"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

var timeNow time.Time = time.Now()
var posts = []model.Post{
	model.Post{
		ID:          1,
		UrlPhoto:    "abc.jpg",
		Description: "abc",
		Likes:       0,
		CreatedAt:   timeNow,
		Username:    "abc",
	},
	model.Post{
		ID:          2,
		UrlPhoto:    "def.jpg",
		Description: "def",
		Likes:       0,
		CreatedAt:   timeNow,
		Username:    "def",
	},
	model.Post{
		ID:          1,
		UrlPhoto:    "ghi.jpg",
		Description: "ghi",
		Likes:       0,
		CreatedAt:   timeNow,
		Username:    "ghi",
	},
}

func TestShowAll(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	dialector := mysql.New(mysql.Config{
		Conn:       db,
		DriverName: "mysql",
	})

	gdb, _ := gorm.Open(dialector, &gorm.Config{})
	postRepo := NewPostRepository(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `posts` ORDER BY created_at desc").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "url_photo", "description", "likes", "created_at", "username"}).
				AddRow(posts[0].ID, posts[0].UrlPhoto, posts[0].Description, posts[0].Likes, posts[0].CreatedAt, posts[0].Username).
				AddRow(posts[1].ID, posts[1].UrlPhoto, posts[1].Description, posts[1].Likes, posts[1].CreatedAt, posts[1].Username).
				AddRow(posts[2].ID, posts[2].UrlPhoto, posts[2].Description, posts[2].Likes, posts[2].CreatedAt, posts[2].Username))
	res, err := postRepo.ShowAll()

	assert.NoError(t, err)
	assert.Equal(t, posts, res)
}

func TestShowAllError(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	dialector := mysql.New(mysql.Config{
		Conn:       db,
		DriverName: "mysql",
	})

	gdb, _ := gorm.Open(dialector, &gorm.Config{})
	postRepo := NewPostRepository(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `posts`").
		WillReturnError(errors.New("Error Database"))
	res, err := postRepo.ShowAll()

	assert.Error(t, err)
	assert.Nil(t, res)
}

//ERROR
// func TestFindById(t *testing.T) {

// 	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	defer db.Close()

// 	dialector := mysql.New(mysql.Config{
// 		Conn:       db,
// 		DriverName: "mysql",
// 	})

// 	gdb, _ := gorm.Open(dialector, &gorm.Config{})
// 	postRepo := NewPostRepository(gdb)

// 	// const query := `SELECT * FROM "posts" WHERE "posts"."id" = $1`
// 	mock.ExpectQuery(regexp.QuoteMeta(
// 		"SELECT * FROM `posts` WHERE `posts`.`id` = $1 ORDER BY `posts`.`id` LIMIT 1")).
// 		WithArgs(1).
// 		WillReturnRows(
// 			sqlmock.NewRows([]string{"id", "url_photo", "description", "likes", "created_at", "username"}).
// 				AddRow(posts[0].ID, posts[0].UrlPhoto, posts[0].Description, posts[0].Likes, posts[0].CreatedAt, posts[0].Username))
// 	res, err := postRepo.FindByID(1)

// 	assert.NoError(t, err)
// 	assert.Equal(t, posts[0], res)
// }
