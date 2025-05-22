package models

import (
	"gorm.io/gorm"
	"strings"
)

type Contact struct {
	gorm.Model
	Name  string
	Email string
}

func (c *Contact) Validate() map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(c.Name) == "" {
		errors["Name"] = "Name is required"
	}
	if strings.TrimSpace(c.Email) == "" {
		errors["Email"] = "Email is required"
	}
	return errors
}
