package node

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	tools "github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// CreateVM is the handler for POST /nodes/{nodeid}/vms
// Creates the VM
func (api NodeAPI) UpdateVM(w http.ResponseWriter, r *http.Request) {
	var reqBody VMUpdate

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	vars := mux.Vars(r)
	vmid := vars["vmid"]

	bp := struct {
		Memory          int         `json:"memory"`
		Cpu             int         `json:"cpu"`
		Nic             []NicLink   `json:"nic"`
		Disks           []VDiskLink `json:"disks"`
	}{
		Memory:          reqBody.Memory,
		Cpu:             reqBody.Cpu,
		Nic:             reqBody.Nics,
		Disks:           reqBody.Disks,
	}

	obj := make(map[string]interface{})
	obj[fmt.Sprintf("vm__%s", vmid)] = bp

	if _, err := tools.ExecuteBlueprint(api.AysRepo, fmt.Sprintf("vm_%s_%+v", vmid, time.Now().Unix()), obj); err != nil {
		log.Errorf("error executing blueprint for vm %s creation : %+v", vmid, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
