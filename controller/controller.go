package controller

import (
	"dba/repository"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetArmy(w http.ResponseWriter, r *http.Request) {
	// Parse any URL parameters if needed.
	r.ParseForm()

	// Get the armyID from the request parameters.
	armyIDStr := r.FormValue("armyid") // Replace with the actual request parameter name
	armyID, err := strconv.Atoi(armyIDStr)
	if err != nil {
		http.Error(w, "Invalid armyid parameter", http.StatusBadRequest)
		return
	}
	// Call the ShowArmyUnits function from the repository to get the army units.
	units, err := repository.QueryUnits(armyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal the units into JSON.
	jsonResponse, err := json.Marshal(units)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
