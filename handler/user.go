package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/config"
	"github.com/golpo/dto"
	"github.com/golpo/util"
	"log"
)

func UserList(c *fiber.Ctx) {
	rows, err := config.DB.Raw("SELECT id, name, email, age FROM users order by id").Rows()
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	defer rows.Close()
	result := dto.Users{}

	for rows.Next() {
		user := dto.User{}
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
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	pStr := util.HashPassword(c, *u.Password)
	u.Password = &pStr
	res := config.DB.Raw("INSERT INTO users (name, email, password, age)VALUES ($1, $2, $3, $4)", u.Name, u.Email, pStr, u.Age).Row()
	//if err != nil {
	//	c.Status(500).Send(err)
	//	return
	//}
	//res := db.DB.Create(&u)
	log.Println(res.Scan())

	log.Println(res)
	u.Password = nil
	if err := c.JSON(u); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func UpdateUser(c *fiber.Ctx) {
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	res := config.DB.Raw("UPDATE users SET name=$1,age=$2 WHERE id=$3", u.Name, u.Age, u.ID).Row()
	//if err != nil {
	//	c.Status(500).Send(err)
	//	return
	//}
	log.Println(res)
	if err := c.Status(201).JSON(u); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func DeleteUser(c *fiber.Ctx) {
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	res := config.DB.Raw("DELETE FROM users WHERE id = $1", u.ID).Row()
	//if err != nil {
	//	c.Status(500).Send(err)
	//	return
	//}
	log.Println(res)
	if err := c.JSON("Deleted"); err != nil {
		c.Status(500).Send(err)
		return
	}
}
