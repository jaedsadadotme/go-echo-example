package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) GetAll(c echo.Context) error {
	users := []User{}

	h.DB.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetOne(c echo.Context) error {
	id := c.Param("id")
	user := User{}

	if err := h.DB.Find(&user, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Create(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&user); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Update(c echo.Context) error {
	id := c.Param("id")
	user := User{}

	if err := h.DB.Find(&user, id).Error; err != nil {
		fmt.Println(err)
	}
	if err := c.Bind(&user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := h.DB.Save(&user); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Destroy(c echo.Context) error {
	id := c.Param("id")
	user := User{}

	if err := h.DB.Delete(&user, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)

}
