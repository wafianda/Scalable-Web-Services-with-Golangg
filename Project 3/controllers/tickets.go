package controllers

import (
	"kanban-board/models"
	"kanban-board/serializers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get all tickets
// @Schemes
// @Description Returns all tickets
// @Tags Tickets list
// @Accept json
// @Produce json
// @Success 200 {array} []models.Ticket
// @Failure 404 {string} error
// @Router /tickets [get]
func GetTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets, err := models.TicketManager.GetAll()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, tickets)
	}
}

// @Summary Get a ticket by id
// @Schemes
// @Description Returns a ticket by id
// @Tags Get Ticket
// @Accept json
// @Produce json
// @Param id path string true "Ticket ID"
// @Success 200 {object} models.Ticket
// @Failure 404 {string} error
// @Router /tickets/:id [get]
func GetTicket() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}

		ticket, err := models.TicketManager.GetById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, ticket)
	}
}

// @Summary Create a new ticket
// @Schemes
// @Description Create a new ticket
// @Tags Create Ticket
// @Accept json
// @Produce json
// @Param ticket body serializers.TicketBody true "Ticket"
// @Success 200 {object} models.Ticket
// @Failure 404 {string} error
// @Router /tickets [post]
func CreateTicket() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body serializers.TicketBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ticket := body.ToModel()
		if err := models.TicketManager.Create(&ticket); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, ticket)
	}
}

// @Summary Update a ticket
// @Schemes
// @Description Update a ticket
// @Tags Update Ticket
// @Accept json
// @Produce json
// @Param id path string true "Ticket ID"
// @Param ticket body serializers.TicketBody true "Ticket"
// @Success 200 {object} models.Ticket
// @Failure 404 {string} error
// @Router /tickets/:id [put]
func UpdateTicket() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		var body serializers.TicketBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ticket := body.ToModel()
		if err := models.TicketManager.Update(&ticket, id); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, ticket)
	}
}

// @Summary Delete ticket by id
// @Schemes
// @Description Delete a ticket by id
// @Tags Delete Ticket
// @Accept json
// @Produce json
// @Param id path string true "Ticket ID"
// @Success 200 {string} string
// @Failure 404 {string} error
// @Router /tickets/:id [delete]
func DeleteTicket() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		if err := models.TicketManager.DeleteById(id); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "Deleted")
	}
}
