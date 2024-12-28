package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucksvinee/api-students/db"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/students", getStudents)
  e.POST("/students", createStudent)
  e.GET("/students/:id", getStudent)
  e.PUT("/students/:id", updateStudent)
  e.DELETE("/students/:id", deleteStudent)


  // Start server
  if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
    slog.Error("failed to start server", "error", err)
  }
}

// Handler
func getStudents(c echo.Context) error {
  return c.String(http.StatusOK, "List of all students")
}

func createStudent(c echo.Context) error {
  student := db.Student{}
  if err := c.Bind(&student); err != nil {
    return err
  }
  if err := db.AddStudent(student); err != nil {
    return c.String(http.StatusInternalServerError, "Error to create student")
  }
  return c.String(http.StatusOK, "Create a new student")
}

func getStudent(c echo.Context) error {
  id := c.Param("id")
  getStud := fmt.Sprintf("Get %s student", id)
  return c.String(http.StatusOK, getStud)
}

func updateStudent(c echo.Context) error {
  id := c.Param("id")
  updateStud := fmt.Sprintf("Update %s student", id)  
  return c.String(http.StatusOK, updateStud)
}

func deleteStudent(c echo.Context) error {
  id := c.Param("id")
  deleteStud := fmt.Sprintf("Delete %s student", id)
  return c.String(http.StatusOK, deleteStud)
}