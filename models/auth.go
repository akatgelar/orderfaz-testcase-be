package models

import "time"

type Auth struct {
	Id					string `json:"id"`
	Msisdn				string `json:"msisdn"`
	Username			string `json:"username"`
	Password			string `json:"password"`
	Name				string `json:"name"`  
	CreatedAt			time.Time `json:"created_at"`
	CreatedBy			string `json:"created_by"` 
}

type AuthLogin struct { 
	Msisdn				string `json:"msisdn"` 
	Password			string `json:"password"` 
}

type AuthRegister struct {  
	Msisdn				string `json:"msisdn"`
	Username			string `json:"username"`
	Password			string `json:"password"`
	Name				string `json:"name"`   
}

// Override default table name
func (Auth) TableName() string {
    return "auth"
}
