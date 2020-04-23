package service

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/db"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// User struct
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      string `json:"age"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

func GetUsers(c *fiber.Ctx) {
	rows, err := db.DB.Query("SELECT id, name, email, age FROM users order by id")
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	defer rows.Close()
	result := Users{}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		result.Users = append(result.Users, user)
	}
	if err := c.JSON(result); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func CreateUser(c *fiber.Ctx) {
	u := new(User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	res, err := db.DB.Query("INSERT INTO users (name, email, password, age)VALUES ($1, $2, $3, $4)", u.Name, u.Email, string(hashedPassword), u.Age)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	log.Println(res)
	if err := c.JSON(u); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func UpdateUser(c *fiber.Ctx) {
	u := new(User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	res, err := db.DB.Query("UPDATE users SET name=$1,age=$2 WHERE id=$3", u.Name, u.Age, u.ID)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	log.Println(res)
	if err := c.Status(201).JSON(u); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func DeleteUser(c *fiber.Ctx) {
	u := new(User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	res, err := db.DB.Query("DELETE FROM users WHERE id = $1", u.ID)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	log.Println(res)
	if err := c.JSON("Deleted"); err != nil {
		c.Status(500).Send(err)
		return
	}
}
