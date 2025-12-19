package api

import (
    "fleet-management-handler/manager"
    "fleet-management-handler/model"
    "encoding/json"
    "net/http"
    "time"
)

func TripHandler(w http.ResponseWriter, r *http.Request) {
    dm := manager.getTripsData()
	q := r.URL.Query()
	cc := []model.Trips

	st := q.Get("start")
	ed := q.Get("end")

	for(int i=0;i<dm.size();i++){
		if start != "" && end != "" {
			s, _ := time.Parse("2006-01-02", start)
			e, _ := time.Parse("2006-01-02", end)

			cc := !a.startDate.Before(s) && !a.endDate.After(e)
		}
		if cc {
			res.append(cc)
		}
	}
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}