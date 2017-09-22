package node

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/zero-os/0-orchestrator/api/handlers"
	"github.com/zero-os/0-orchestrator/api/httperror"
)

// GetNode is the handler for GET /nodes/{nodeid}
// Get detailed information of a node
func (api *NodeAPI) GetNode(w http.ResponseWriter, r *http.Request) {
	// aysClient := tools.GetAysConnection(r, api)
	vars := mux.Vars(r)
	nodeID := vars["nodeid"]

	service, err := api.client.GetService("node", nodeID, "", nil)
	if err != nil {
		handlers.HandleError(w, err)
		return
	}

	var respBody NodeService
	if err := json.Unmarshal(service.Data, &respBody); err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err, "Error unmarshaling ays response")
		return
	}
	var node Node
	node.IPAddress = respBody.RedisAddr
	node.Status = respBody.Status
	node.Hostname = respBody.Hostname
	node.Id = service.Name

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&node)
}
