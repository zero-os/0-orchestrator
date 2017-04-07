package node

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/g8os/grid/api/tools"
)

// ListVMs is the handler for GET /node/{nodeid}/vm
// List VMs
func (api NodeAPI) ListVMs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	queryParams := map[string]interface{}{
		"fields": "status,id",
		"parent": vars["nodeid"],
	}
	services, res, err := api.AysAPI.Ays.ListServicesByRole("vm", api.AysRepo, nil, queryParams)
	if !tools.HandleAYSResponse(err, res, w, "listing vms") {
		return
	}

	var respBody = make([]VMListItem, len(services))
	for i, service := range services {
		vm := VMListItem{
			Id: service.Name,
		}
		if err := json.Unmarshal(service.Data, &vm); err != nil {
			tools.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		respBody[i] = vm
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)
}
