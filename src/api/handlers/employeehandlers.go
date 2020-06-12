package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Employee struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Birthday   string `json:"birthday"`
	Gender     string `json:"gender"`
	Address    string `json:"address"`
	City       string `json:"city"`
	University string `json:"university"`
	StartDate  string `json:"startdate"`
	RefUser    string `json:"refuser"`
	IsAdmin    string `json:"isadmin"`
	// Salary     string `json: "salary"`
	// Age        string `json : "age"`
}
type Employees struct {
	Employees []Employee `json:"employees"`
}

func AddEmployee(c echo.Context) error {

	//Connect DB
	var err error
	db, err := sql.Open("postgres", "user=postgres password=root dbname=books_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	u := new(Employee)
	if err := c.Bind(u); err != nil {
		return err
	}

	//Check employee

	// Add employee
	sqlStatement := "INSERT INTO employees (username, password , email, name, birthday, gender, address, city, university, startdate, refuser, isadmin)VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	res, err := db.Query(sqlStatement,
		u.Username, u.Password, u.Email, u.Name, u.Birthday, u.Gender, u.Address, u.City, u.University, u.StartDate, u.RefUser, u.IsAdmin)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "employee added")
}

func DeleteEmployee(c echo.Context) error {

	//Connect DB
	var err error
	db, err := sql.Open("postgres", "user=postgres password=root dbname=books_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	u := new(Employee)
	if err := c.Bind(u); err != nil {
		return err
	}

	id := c.Param("id")
	sqlStatement := "DELETE FROM employees WHERE id = $1"
	res, err := db.Query(sqlStatement, id)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusOK, "Deleted")
	}
	return c.String(http.StatusOK, id+"Deleted")
}

//UPDATE
func UpdateEmployee(c echo.Context) error {
	//Connect DB
	var err error
	db, err := sql.Open("postgres", "user=postgres password=root dbname=books_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	u := new(Employee)
	if err := c.Bind(u); err != nil {
		return err
	}
	sqlStatement := "UPDATE employees SET userame=$1,WHERE id=$2"
	res, err := db.Query(sqlStatement, u.Username, u.Id)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, u.Id)
}

// func ShowEmployee(c echo.Context) error {

// 	//Connect DB
// 	var err error
// 	db, err := sql.Open("postgres", "user=postgres password=root dbname=books_database sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err = db.Ping(); err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("DB Connected...")
// 	}

// 	u := new(Employee)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}

// 	sqlStatement := "SELECT id, username FROM employees order by id"
// 	rows, err := db.Query(sqlStatement)
// 	if err != nil {
// 		fmt.Println(err)
// 		//return c.JSON(http.StatusCreated, u);
// 	}
// 	defer rows.Close()
// 	result := Employees{}

// 	for rows.Next() {
// 		employee := Employees{}
// 		err2 := rows.Scan(&employee.Id, &employee.Username)
// 		// Exit if we get an error
// 		if err2 != nil {
// 			return err2
// 		}
// 		result.Employees = append(result.Employees, Employee)
// 	}
// 	return c.JSON(http.StatusCreated, result)

// 	//return c.String(http.StatusOK, "ok")

// }
