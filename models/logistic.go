package models

import "time"

type Logistics struct {
	Id					string `json:"id"`
	LogisticName		string `json:"logistic_name"`
	Amount				float32 `json:"amount"`
	DestinationName		string `json:"destination_name"`
	OriginName			string `json:"origin_name"`
	Duration			string `json:"duration"`
	IsActive			bool `json:"is_active"`
	CreatedAt			time.Time `json:"created_at"`
	CreatedBy			string `json:"created_by"`
	UpdatedAt			time.Time `json:"updated_at"`
	UpdatedBy			string `json:"updated_by"`
}

type LogisticsCreate struct { 
	LogisticName		string `json:"logistic_name"`
	Amount				float32 `json:"amount"`
	DestinationName		string `json:"destination_name"`
	OriginName			string `json:"origin_name"`
	Duration			string `json:"duration"` 
}