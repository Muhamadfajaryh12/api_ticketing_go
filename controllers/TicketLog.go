package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
)

func InsertTicketLog( ticketID uint, statusID uint) error {

	ticketLog := model.TicketLogForm{
		TicketID: ticketID,
		StatusID: statusID,
	}
	
	if err := config.DB.Create(&ticketLog).Error; err != nil{
		return err
	}
	
	return nil
}