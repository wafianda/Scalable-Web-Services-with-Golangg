package controllers

import (
	"kanban-board/models"
	"kanban-board/serializers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get all columns
// @Schemes
// @Description Returns all columns
// @Tags GetColumns
// @Accept json
// @Produce json
// @Success 200 {array} []models.Column
// @Failure 404 {string} error
// @Router /columns [get]
func GetColumns() gin.HandlerFunc {
	return func(c *gin.Context) {
		columns, err := models.ColumnManager.GetAll()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, columns)
	}
}

// @Summary Get a column by id
// @Schemes
// @Description Returns a column by id
// @Tags GetColumn
// @Accept json
// @Produce json
// @Param id path string true "Column ID"
// @Success 200 {object} models.Column
// @Failure 404 {string} error
// @Router /columns/:id [get]
func GetColumn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		column, err := models.ColumnManager.GetById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, column)
	}
}

// @Summary Create a new column
// @Schemes
// @Description Create a new column
// @Tags CreateColumn
// @Accept json
// @Produce json
// @Param column body serializers.ColumnBody true "Column"
// @Success 200 {object} models.Column
// @Failure 404 {string} error
// @Router /columns [post]
func CreateColumn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body serializers.ColumnBody
		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		column := body.ToModel()
		if err := models.ColumnManager.Create(&column); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, column)
	}
}

// @Summary Update a column
// @Schemes
// @Description Update a column
// @Tags UpdateColumn
// @Accept json
// @Produce json
// @Param id path string true "Column ID"
// @Param column body serializers.ColumnBody true "Column"
// @Success 200 {object} models.Column
// @Failure 404 {string} error
// @Router /columns/:id [put]
func UpdateColumn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body serializers.ColumnBody
		id, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		column := body.ToModel()
		if err := models.ColumnManager.Update(&column, id); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, column)
	}
}

// @Summary Delete column by id
// @Schemes
// @Description Delete a column by id
// @Tags DeleteColumn
// @Accept json
// @Produce json
// @Param id path string true "Column ID"
// @Success 200 {string} string
// @Failure 404 {string} error
// @Router /columns/:id [delete]
func DeleteColumn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		err = models.ColumnManager.DeleteById(uint(id))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "Deleted")
	}
}
