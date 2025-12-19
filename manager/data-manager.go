
package manager

import (
	"fleet-management-handler/model"
	"encoding/csv"
	"strconv"
	"os"
)

type Cur struct {
	VehicleID int
	Driver string 
	Date time.Date(2025, 3, 10, 0, 0, 0, 0, time.UTC)
	Odometer_Reading long
	Trip_Distance long
}

type VehicleManager struct{
	loaded []Cur
}

var inst *VehicleManager

func VehicleManager() *VehicleManager {
	once.Do(func() {
		instance = &VehicleManager{
			Cur: loadCSV(),
		}
	})
	return instance
}


//Data Cleaning and Insert
func loadCSV() {
	f, err := os.Open("driver_logs_raw.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Read()

	data = []

	for {
		row, err := r.Read()
		if err != nil {
			break
		}

		id, _ := strconv.Atoi(row[0])
		reading, err := strconv.Atoi(row[3])
		if(err != nil) continue
		dist, err := strconv.Atoi(row[4])
		if(err != nil) continue
		d, _ := row[2]

		new_date = "";

		if(!(d[2] == '-' && d[7] == '-') || !(d[2] == '/' && d[7] == '/')) continue

		if(d[2] == '-' || d[2] == '/'){
			new_date += d[6]
			new_date += d[7]
			new_date += d[8]
			new_date += d[9]
			new_date += '-'
			new_date += d[4]
			new_date += d[5]
			new_date += '-'
			new_date += d[1]
			new_date += d[2]
		}
		else if(d[4] == '/'){
			new_date = date
			new_date[4] = '-'
			new_date[7] = '-'
		}
		else{
			new_date = date
		}

		new_date = time.Parse("2006-01-02", new_date)

		pr := Cur{
			VehicleID: id,
			Driver: row[1],
			Date:  new_date,
			Odometer_Reading: reading,
			Trip_Distance: dist,
		}

		data.append(Cur)
	}
	return cur
}