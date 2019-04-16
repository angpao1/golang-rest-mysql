package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func createUser(c echo.Context) error {
	emp := new(Employee)
	if err := c.Bind(emp); err != nil {
		return err
	}
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/goblog")
	if err != nil {
		fmt.Println(err.Error())
	}

	sql := "INSERT INTO employee (name, city) VALUES (?, ?)"

	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	_, err2 := stmt.Exec(emp.Name, emp.City)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	return c.JSON(http.StatusCreated, "Create Successful")
}

func getUser(c echo.Context) error {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/goblog")
	if err != nil {
		fmt.Println(err.Error())
	}
	rs, err := db.Query("SELECT id, name, city FROM employee WHERE 1 = 1")
	if err != nil {
		fmt.Println(err)
	}

	var users []Employee

	for rs.Next() {
		var user Employee

		err := rs.Scan(
			&user.ID,
			&user.Name,
			&user.City)

		if err != nil {
			return err
		}
		users = append(users, user)
	}
	defer rs.Close()
	defer db.Close()
	return c.JSON(http.StatusOK, users)
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/goblog")
	if err != nil {
		fmt.Println(err.Error())
	}
	var user Employee
	err = db.QueryRow("SELECT id, name, city FROM employee WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.City)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	emp := new(Employee)
	if err := c.Bind(emp); err != nil {
		return err
	}
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/goblog")
	if err != nil {
		fmt.Println(err.Error())
	}
	sql := "UPDATE employee SET name=?, city=? WHERE id=?"

	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	stmt.Exec(emp.Name, emp.City, id)
	// Make sure to cleanup after the program exits
	defer stmt.Close()
	defer db.Close()
	return c.JSON(http.StatusOK, "updateUser")
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/goblog")
	if err != nil {
		fmt.Println(err.Error())
	}
	sql := "DELETE FROM employee WHERE id=?"

	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	stmt.Exec(id)
	defer stmt.Close()
	defer db.Close()
	return c.NoContent(http.StatusNoContent)
}

func returnResquest(c echo.Context) error {
	u := new(Employee)

	if err := c.Bind(u); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, u)
}
