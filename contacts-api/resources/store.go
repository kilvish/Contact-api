package resources

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/anil/contacts-api/internal/users"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Store interface abstracts the store layer
type Store interface {
	users.Repo
}

type MysqlConnection struct {
	Host   string `json:"host"`
	DBName string `json:"dbname"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
}

//DB is singelton DB
type mysqldb struct {
	db *gorm.DB
}

var MysqlConnString string

// Init initializes the store

func Init(Type string) (Store, error) {
	switch Type {
	case "mysql":
		raw, err := ioutil.ReadFile("../resources/mysqlConfig.json")
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		var myconn MysqlConnection
		json.Unmarshal(raw, &myconn)
		MysqlConnString = myconn.User + ":" + myconn.Pass + "@tcp(" + myconn.Host + ")/" + myconn.DBName + "?charset=utf8&parseTime=True&loc=Local"
		fmt.Println(MysqlConnString)
		DB, err := gorm.Open("mysql", MysqlConnString)
		defer DB.Close()
		if err != nil {
			fmt.Println(err)
			log.Fatal("Connection to DB failed")
			return nil, errors.New("connection Failed")
		}
		return &mysqldb{
			db: DB,
		}, nil
	default:
		{
			return nil, errors.New("Connection now allowed")
		}
	}
}

func (db *mysqldb) AddUser(ctx context.Context, user *users.User) error {
	db.db.AutoMigrate(user)
	fmt.Println(user)
	e := db.db.Create(user)
	if e != nil {
		errors.New("DB Operation Failed")
	}
	return nil
}

func (db *mysqldb) deleteUser(ctx context.Context, user *users.User) error {
	e := db.db.Delete(user)
	if e != nil {
		errors.New("DB Operation Failed")
	}
	return nil
}

func (db *mysqldb) GetUserByEmail(ctx context.Context, email string) (*users.User, error) {
	user := &users.User{}
	e := db.db.Table("User").Where("email = ?", email).First(user).Error
	if e != nil {
		errors.New("DB Operation Failed")
	}
	return user, nil
}

func (db *mysqldb) GetUserByID(ctx context.Context, id string) (*users.User, error) {
	user := &users.User{}
	e := db.db.Table("User").Where("user_id = ?", id).First(user).Error
	if e != nil {
		errors.New("DB Operation Failed")
	}
	return user, nil
}
