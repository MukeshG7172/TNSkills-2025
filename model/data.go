package model

import "time"

type Vehicles stuct{
	vehicleID int `json:"vehicleID"`
	regNo string `json:"regNo"`
	vehicle_type string `json:"vehicleType"`
	lastServiceDate time.Date(2025, 3, 10, 0, 0, 0, 0, time.UTC) `json:"lastServiceDate"`  
}

type Trips struct{
	tripID int `json:"tripID"`
	vehicleID int `json:"vehicleID"`
	driverName string `json:"driverName"`
	startDate time.Date(2025, 3, 10, 0, 0, 0, 0, time.UTC) `json:"startDate"`
	endDate time.Date(2025, 3, 10, 0, 0, 0, 0, time.UTC) `json:"endDate"`  
}

type Maintenance_Alerts struct{
	
}