
package manager

import (
	"fleet-management-handler/model"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Cur struct {
	VehicleID int
	Driver string 
	Date time.Date(2025, 3, 10, 0, 0, 0, 0, time.UTC)
	Odometer_Reading long
	Trip_Distance long
}

var inst *VehicleManager
var once sync.Once

func VehicleManager() *VehicleManager {
	once.Do(func() {
		instance = &VehicleManager{
			vehicles: seedVehicles(),
		}
	})
	return instance
}

func loadCSV() {
	f, err := os.Open("C:\Users\Administrator\Documents\tn-skills\TNSkills-2025\driver_logs_raw.csv")
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
		d, _ := row[2]

		new_date = "";

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

		pr := Cur{
			VehicleID: id,
			Driver: row[1],
			Date:  new_date,
			Price:     price,
			Available: avail,
		}

		data.append(Cur)
	}
}