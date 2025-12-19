
package manager

import (
	"encoding/csv"
	"strconv"
	"sql"
	_ "github.com/go-sql-server/mysql"
)


func getVehicleData() {
	cur_data := []
	d := "tnskills:1234@tcp(localhost:3306)"
	f, err := sql.connect("mysql",d)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res, err := d.exec("select * from vehicles")
	if( err != nil ){
		panic(err)
	}


	while(res.next()){
		cur := model.Vehicles{
			vehicleID: res.vehicleID
			regNo: res.regNo
			vehicle_type: res.vehicle_type
			lastServiceData: res.lastServiceData
			currentOdometer: res.currentOdometer
			lastServiceOdometer: res.lastServiceOdometer
		}
		res = res.next()
		cur_data.append(cur)
	}

	return cur_data
}