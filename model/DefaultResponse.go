package model

import (
	"time"
)

type DefaultResponse struct {
	Code string    `json:"code"`
	Date time.Time `json:"date"`
	Msgs []string  `json:"msgs"`
}

func (d *DefaultResponse) NewDefault(code string, message string) {
	d.Code = code
	d.addDateNow()
	d.AddMessage(message)
}

func (d *DefaultResponse) AddMessage(message string) {
	d.Msgs = append(d.Msgs, message)
}

func (d *DefaultResponse) addDateNow() {
	d.Date = time.Now()
}
