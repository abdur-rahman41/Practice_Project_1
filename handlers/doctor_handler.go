package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/yourproject/models"
)

var doctors = []models.Doctor{}

func GetDoctors(c echo.Context) error {
	return c.JSON(http.StatusOK, doctors)
}

func GetDoctorByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	for _, doctor := range doctors {
		if doctor.ID == id {
			return c.JSON(http.StatusOK, doctor)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Doctor not found"})
}

func CreateDoctor(c echo.Context) error {
	var doctor models.Doctor
	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid data"})
	}
	doctor.ID = len(doctors) + 1
	doctors = append(doctors, doctor)
	return c.JSON(http.StatusCreated, doctor)
}

func UpdateDoctor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	for i, doctor := range doctors {
		if doctor.ID == id {
			var updatedDoctor models.Doctor
			if err := c.Bind(&updatedDoctor); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid data"})
			}
			updatedDoctor.ID = doctor.ID
			doctors[i] = updatedDoctor
			return c.JSON(http.StatusOK, updatedDoctor)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Doctor not found"})
}

func DeleteDoctor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	for i, doctor := range doctors {
		if doctor.ID == id {
			doctors = append(doctors[:i], doctors[i+1:]...)
			return c.JSON(http.StatusOK, doctor)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Doctor not found"})
}
