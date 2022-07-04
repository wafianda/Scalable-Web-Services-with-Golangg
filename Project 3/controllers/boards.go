package controllers

import (
	"kanban-board/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Getboards returns all boards

// @Summary get all boards
// @Schemes
// @Description Returns all boards
// @Tags GetBoards
// @Accept json
// @Produce json
// @Success 200 {array} []models.Board
// @Failure 404 {string} error
// @Router /boards [get]
func GetBoards() gin.HandlerFunc {
	return func(c *gin.Context) {
		boards, err := models.BoardManager.GetAll()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, boards)
	}
}

// GetBoard returns a Board by id

// @Summary Return a Board by id
// @Schemes
// @Description Return a Board by id
// @Tags GetBoard
// @Accept json
// @Produce json
// @Param id path string true "Board ID"
// @Success 200 {object} models.Board
// @Failure 404 {string} error
// @Router /boards/:id [get]
func GetBoard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		board, err := models.BoardManager.GetById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, board)
	}
}

type BoardBody struct {
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}

// @Summary Create a new Board
// @Schemes
// @Description Create a new Board
// @Tags CreateBoard
// @Accept json
// @Produce json
// @Param Board body BoardBody true "Board"
// @Success 200 {object} models.Board
// @Failure 404 {string} error
// @Router /boards [post]
func CreateBoard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body BoardBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		board := models.Board{
			Name:      body.Name,
			StartDate: body.StartDate,
			EndDate:   body.EndDate,
		}
		if err := models.BoardManager.Create(&board); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, board)
	}
}

// @Summary Update a board
// @Schemes
// @Description Update a board
// @Tags UpdateBoard
// @Accept json
// @Produce json
// @Param id path string true "Board ID"
// @Param board body BoardBody true "Board"
// @Success 200 {object} models.Board
// @Failure 404 {string} error
// @Router /boards/:id [put]
func UpdateBoard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		var body BoardBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		board := models.Board{
			Name:      body.Name,
			StartDate: body.StartDate,
			EndDate:   body.EndDate,
		}
		if err := models.BoardManager.Update(&board, id); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, board)
	}
}

// @Summary Delete a board by id
// @Schemes
// @Description Delete a board by id
// @Tags DeleteBoard
// @Accept json
// @Produce json
// @Param id path string true "Board ID"
// @Success 200 {string} string
// @Failure 404 {string} error
// @Router /boards/:id [delete]
func DeleteBoard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		if err := models.BoardManager.DeleteById(id); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "Deleted")
	}
}
