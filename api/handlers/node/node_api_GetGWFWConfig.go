package node

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	client "github.com/zero-os/0-core/client/go-client"
	"github.com/zero-os/0-orchestrator/api/httperror"
)

// GetGWFWConfig is the handler for GET /nodes/{nodeid}/gws/{gwname}/advanced/firewall
// Get current FW config
// Once used you can not use gw.portforwards any longer
func (api *NodeAPI) GetGWFWConfig(w http.ResponseWriter, r *http.Request) {
	var config bytes.Buffer

	vars := mux.Vars(r)
	gwname := vars["gwname"]

	node, err := api.client.GetNodeConnection(r)
	containerID, err := api.client.GetContainerID(r, gwname)
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err, "Failed to establish connection to node")
		return
	}

	containerClient := client.Container(node).Client(containerID)
	err = client.Filesystem(containerClient).Download("/etc/nftables.conf", &config)
	if err != nil {
		errmsg := fmt.Sprintf("Error getting  file from container '%s' at path '%s'.\n", gwname, "/etc/nftables.conf")
		httperror.WriteError(w, http.StatusInternalServerError, err, errmsg)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(config.Bytes())
}
