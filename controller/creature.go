package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ANijat/creature-crud/model"
	"github.com/ANijat/creature-crud/storage"
	"github.com/labstack/echo"
)

func GetCreatures(c echo.Context) error {
	db := storage.GetDBInstance()
	creatures := []model.Creatures{}

	if err := db.Find(&creatures).Error; err != nil {
		message := model.Error{err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}
	return c.JSON(http.StatusOK, creatures)
}

func GetCreaturesByid(c echo.Context) error {
	id := c.Param("id")
	db := storage.GetDBInstance()
	creature := model.Creatures{}
	if err := db.First(&creature, id).Error; err != nil {
		message := model.Error{err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}
	return c.JSON(http.StatusOK, creature)
}

func UpdateCreature(c echo.Context) error {
	creature := model.Creatures{}
	id := c.Param("id")
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &creature)
	if err != nil {
		return err
	}
	db := storage.GetDBInstance()
	if err := db.Model(&creature).Where("id = ?", id).Update("name", creature.Name).Error; err != nil {
		message := model.Error{err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}
	return c.JSON(http.StatusOK, GetCreaturesByid(c))
}

func CreateCreature(c echo.Context) error {
	db := storage.GetDBInstance()
	creature := model.Creatures{}
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &creature)
	if err != nil {
		return err
	}
	if err := db.Save(&creature).Error; err != nil {
		message := model.Error{err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}
	return c.JSON(http.StatusOK, creature)
}
