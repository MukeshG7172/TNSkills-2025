
package manager

import (
	"encoding/csv"
	"strconv"
	"sql"
	_ "github.com/go-sql-server/mysql"
)


//Fetch data
func getTripsData() {
	cur_data := []
	d := "tnskills:1234@tcp(localhost:3306)"
	f, err := sql.connect("mysql",d)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res, err := d.exec("select * from trips")
	if( err != nil ){
		panic(err)
	}


	while(res.next()){
		cur := model.Trips{
			tripID: res.tripID
			vehicleID: res.vehicleID
			driverName: res.driverName
			startDate: res.startDate
			endDate: res.endDate
		}
		res = res.next()
		cur_data.append(cur)
	}

	return cur_data
}