package service

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/db"
	"log"
)

// Employee struct
type Employee struct {
	ID     string `json: "id"`
	Name   string `json: "name"`
	Salary string `json: "salary"`
	Age    string `json: "age"`
}

// Employees struct
type Employees struct {
	Employees []Employee `json: "employees"`
}

func GetEmployees(c *fiber.Ctx) {
	rows, err := db.DB.Query("SELECT id, name, salary, age FROM employees order by id")
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age)
		// Exit if we get an error
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		// Append Employee to Employees
		result.Employees = append(result.Employees, employee)
	}
	// Return Employees in JSON format
	if err := c.JSON(result); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func CreateEmployee(c *fiber.Ctx) {
	// New Employee struct
	u := new(Employee)
	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	// Insert Employee into database
	res, err := db.DB.Query("INSERT INTO employees (name, salry, age)VALUES ($1, $2, $3)", u.Name, u.Salary, u.Age)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	// Print result
	log.Println(res)
	// Return Employee in JSON format
	if err := c.JSON(u); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func UpdateEmployee(c *fiber.Ctx) {
	// New Employee struct
	u := new(Employee)
	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	// Insert Employee into database
	res, err := db.DB.Query("UPDATE employees SET name=$1,salary=$2,age=$3 WHERE id=$5", u.Name, u.Salary, u.Age, u.ID)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	// Print result
	log.Println(res)
	// Return Employee in JSON format
	if err := c.Status(201).JSON(u); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func DeleteEmployee(c *fiber.Ctx) {
	// New Employee struct
	u := new(Employee)
	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	// Insert Employee into database
	res, err := db.DB.Query("DELETE FROM employees WHERE id = $1", u.ID)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	// Print result
	log.Println(res)
	// Return Employee in JSON format
	if err := c.JSON("Deleted"); err != nil {
		c.Status(500).Send(err)
		return
	}
}
