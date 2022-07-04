package models

import "reflect"

type ticketManager struct{}

var TicketManager = ticketManager{}

func (t *ticketManager) Create(ticket *Ticket) error {
	return DB.Create(&ticket).Error
}

func (t *ticketManager) Update(ticket *Ticket, id uint64) error {
	return DB.Model(Ticket{}).
		Where("id = ?", id).
		Updates(ticket).
		Error
}

func (t *ticketManager) GetById(id uint64) (*Ticket, error) {
	var ticket Ticket
	return &ticket, DB.Preload("Board").Preload("Column").First(&ticket, id).Error
}

func (t *ticketManager) GetAll() (*[]Ticket, error) {
	var tickets []Ticket
	return &tickets, DB.Preload("Board").Preload("Column").Find(&tickets).Error
}

func (t *ticketManager) GetAllByQuery(q interface{}) ([]Ticket, error) {
	var tickets []Ticket
	v := reflect.ValueOf(q)
	typeOfS := v.Type()
	qs := DB.
		Preload("Board").
		Joins("Board")

	for i := 0; i < v.NumField(); i++ {
		key := typeOfS.Field(i).Name    // Key
		value := v.Field(i).Interface() // vaalue
		qs.Where(key+" LIKE %?%", value)
	}

	return tickets, qs.Find(&tickets).Error
}

func (t *ticketManager) DeleteById(id uint64) error {
	return DB.Delete(&Ticket{}, id).Error
}
