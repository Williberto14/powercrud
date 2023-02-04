package model

import (
	"time"

	"gorm.io/datatypes"
)

type Product struct {
	ProductId       string         `gorm:"type:varchar(45);primary_key; not null" json:"product_id"`
	Name            string         `gorm:"type:varchar(45)" json:"name"`
	Description     string         `gorm:"type:text(300)" json:"description"`
	Status          string         `gorm:"type:varchar(45)" json:"status"`
	CreationDate    time.Time      `json:"creation_date"`
	UpdateDate      time.Time      `json:"update_date"`
	AccountId       string         `gorm:"type:varchar(25)" json:"account_id"`
	FormatProduct   datatypes.JSON `json:"format_product"`
	ValueUnit       float64        `gorm:"type:decimal(15,2)" json:"value_unit"`
	UnitName        string         `gorm:"type:varchar(45)" json:"unit_name"`
	UnitDescription string         `gorm:"type:text(300)" json:"unit_description"`
	Stock           int            `json:"stock"`
}
