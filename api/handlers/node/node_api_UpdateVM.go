package node

import (
	"encoding/json"
	"fmt"

	"github.com/zero-os/0-orchestrator/api/ays"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/zero-os/0-orchestrator/api/handlers"
	"github.com/zero-os/0-orchestrator/api/httperror"
)

// UpdateVM is the handler for PUT /nodes/{nodeid}/vms/{vmid}
// Updates the VM
func (api *NodeAPI) UpdateVM(w http.ResponseWriter, r *http.Request) {
	// aysClient := tools.GetAysConnection(r, api)
	var reqBody VMUpdate

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		httperror.WriteError(w, http.StatusBadRequest, err, "Error decoding request body")
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		httperror.WriteError(w, http.StatusBadRequest, err, "")
		return
	}

	vars := mux.Vars(r)
	vmID := vars["vmid"]

	srv, err := api.client.GetService("vm", vmID, "", nil)
	if err != nil {
		handlers.HandleError(w, err)
		return
	}
	// srv, res, err := aysClient.Ays.GetServiceByName(vmID, "vm", api.AysRepo, nil, nil)
	// if !tools.HandleAYSResponse(err, res, w, fmt.Sprintf("getting vm %s details", vmID)) {
	// 	return
	// }

	var vm VM
	if err := json.Unmarshal(srv.Data, &vm); err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err, "Error unmarshaling ays response")
		return
	}

	if (vm.Memory != reqBody.Memory || vm.Cpu != reqBody.Cpu) && vm.Status != "halted" {
		err = fmt.Errorf("Can't update memory or CPU of VM %s while it's running", vm.Id)
		httperror.WriteError(w, http.StatusBadRequest, err, "")
		return
	}

	bp := struct {
		Memory int         `yaml:"memory" json:"memory"`
		CPU    int         `yaml:"cpu" json:"cpu"`
		Nics   []NicLink   `yaml:"nics" json:"nics"`
		Disks  []VDiskLink `yaml:"disks" json:"disks"`
	}{
		Memory: reqBody.Memory,
		CPU:    reqBody.Cpu,
		Nics:   reqBody.Nics,
		Disks:  reqBody.Disks,
	}

	blueprint := ays.Blueprint{
		fmt.Sprintf("vm__%s", vmID): bp,
	}
	bpName := ays.BlueprintName("vm", vmID, "update")
	if err := api.client.CreateExec(bpName, blueprint); err != nil {
		handlers.HandleError(w, err)
		return
	}
	// obj[fmt.Sprintf("vm__%s", vmID)] = bp

	// _, err = aysClient.ExecuteBlueprint(api.AysRepo, "vm", vmID, "update", obj)

	// errmsg := fmt.Sprintf("error executing blueprint for vm %s creation ", vmID)
	// if !tools.HandleExecuteBlueprintResponse(err, w, errmsg) {
	// 	return
	// }

	w.WriteHeader(http.StatusNoContent)
}
