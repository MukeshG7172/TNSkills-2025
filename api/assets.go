package api

import (
    "fleet-management-handler/manager"
    "fleet-management-handler/model"
    "fleet-management-handler/search"
    "encoding/json"
    "net/http"
    "strings"
    "time"
)

func AssetHandler(w http.ResponseWriter, r *http.Request) {
    dm := manager.getTripsData()
	q := r.URL.Query()

	st := q.Get("start")
	ed := q.Get("end")

	for(int i=0;i<dm.size();i++){
		if start != "" && end != "" {
			s, _ := time.Parse("2006-01-02", start)
			e, _ := time.Parse("2006-01-02", end)

			return !a.startDate.Before(s) && !a.endDate.After(e)
		}
	}
}