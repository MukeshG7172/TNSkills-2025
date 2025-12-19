package manager

import (
	"strconv"
	"sql"
	_ "github.com/go-sql-server/mysql"
)

func checkMaintenance() {
	vehicle_data := getVehicleData()
	cur_Date := time.Date()
	req := []

	for(int i=0;i<vehicle_data.size();i++){
		cur_vch = vehicle_data[i]
		if(cur_vch.lastServiceDate - cur_Date > 180) {
			req.push(cur_vch)
		}
	}

	return req
}