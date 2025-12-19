package manager

import (
	"encoding/csv"
	"strconv"
	"sql"
	_ "github.com/go-sql-server/mysql"
)

func putVehicleData() {
	cur_data := []
	d := "tnskills:1234@tcp(localhost:3306)"
	f, err := sql.connect("mysql",d)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f, err := os.Open("driver_logs_raw.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Read()

	data := []

	for {
		row, err := r.Read()
		if err != nil {
			break
		}
		d.exec("insert into vehicles values ($vehicleID,$regNo,$vehicle_type,$lastServiceDate,$currentOdometer,lastServiceOdometer)")
	}
}