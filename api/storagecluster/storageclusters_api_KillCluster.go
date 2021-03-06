package storagecluster

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/zero-os/0-orchestrator/api/tools"
)

// KillCluster is the handler for DELETE /storageclusters/{label}
// Kill cluster
func (api *StorageclustersAPI) KillCluster(w http.ResponseWriter, r *http.Request) {
	aysClient, err := tools.GetAysConnection(api)
	if err != nil {
		tools.WriteError(w, http.StatusUnauthorized, err, "")
		return
	}
	vars := mux.Vars(r)
	storageCluster := vars["label"]

	service, resp, err := aysClient.Ays.GetServiceByName(storageCluster, "storagecluster", api.AysRepo, nil, nil)
	if err != nil {
		errmsg := fmt.Sprintf("error getting storagecluster %s service", storageCluster)
		tools.WriteError(w, http.StatusInternalServerError, err, errmsg)
		return
	}

	if resp.StatusCode == http.StatusNotFound {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Prevent deletion of nonempty clusters
	query := map[string]interface{}{
		"consume": fmt.Sprintf("storagecluster!%s", storageCluster),
	}
	services, res, err := aysClient.Ays.ListServicesByRole("vdiskstorage", api.AysRepo, nil, query)
	if !tools.HandleAYSResponse(err, res, w, "listing vdisks") {
		return
	}

	if len(services) > 0 {
		err := fmt.Errorf("Can't delete storage clusters with attached vdiskstorage")
		tools.WriteError(w, http.StatusBadRequest, err, "")
		return
	}

	// execute the delete action
	blueprint := map[string]interface{}{
		"actions": []tools.ActionBlock{{
			Action:  "delete",
			Actor:   service.Actor,
			Service: storageCluster,
		}},
	}

	run, err := aysClient.ExecuteBlueprint(api.AysRepo, "storagecluster", storageCluster, "delete", blueprint)
	if err != nil {
		httpErr := err.(tools.HTTPError)
		tools.WriteError(w, httpErr.Resp.StatusCode, httpErr, "Error executing blueprint for storagecluster deletion")
		return
	}

	// Wait for the delete job to be finshed before we delete the service
	if _, err = aysClient.WaitRunDone(run.Key, api.AysRepo); err != nil {
		httpErr, ok := err.(tools.HTTPError)
		if ok {
			tools.WriteError(w, httpErr.Resp.StatusCode, httpErr, "Error running blueprint for storagecluster deletion")
		} else {
			tools.WriteError(w, http.StatusInternalServerError, err, "Error running blueprint for storagecluster deletion")
		}
		return
	}

	res, err = aysClient.Ays.DeleteServiceByName(storageCluster, "storagecluster", api.AysRepo, nil, nil)
	if !tools.HandleAYSDeleteResponse(err, res, w, "deleting storage_cluster") {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
