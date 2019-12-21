package route_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/isgo-golgo13/gorilla-restsvc/data"
)

//IndexHandler
func GetEngines(w http.ResponseWriter, r *http.Request) {
	engines := []data.Engine{
		{
			ID:            "100000001",
			SerialID:      "VW_100000001_QUATTRO",
			Configuration: "V8",
		},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(engines); err != nil {
		panic(err)
	}
}
