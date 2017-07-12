package node

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NodesInterface is interface for /nodes root endpoint
type NodesInterface interface { // DeleteBridge is the handler for DELETE /nodes/{nodeid}/bridges/{bridgeid}
	// Remove bridge
	DeleteBridge(http.ResponseWriter, *http.Request)
	// GetBridge is the handler for GET /nodes/{nodeid}/bridges/{bridgeid}
	// Get bridge details
	GetBridge(http.ResponseWriter, *http.Request)
	// ListBridges is the handler for GET /nodes/{nodeid}/bridges
	// List bridges
	ListBridges(http.ResponseWriter, *http.Request)
	// CreateBridge is the handler for POST /nodes/{nodeid}/bridges
	// Creates a new bridge
	CreateBridge(http.ResponseWriter, *http.Request)
	// GetContainerCPUInfo is the handler for GET /nodes/{nodeid}/containers/{containername}/cpus
	// Get detailed information of all CPUs in the container
	GetContainerCPUInfo(http.ResponseWriter, *http.Request)
	// GetContainerDiskInfo is the handler for GET /nodes/{nodeid}/containers/{containername}/disks
	// Get detailed information of all the disks in the container
	// !!!!!!!GetContainerDiskInfo(http.ResponseWriter, *http.Request)
	// FileDelete is the handler for DELETE /nodes/{nodeid}/containers/{containername}/filesystem
	// Delete file from container
	FileDelete(http.ResponseWriter, *http.Request)
	// FileDownload is the handler for GET /nodes/{nodeid}/containers/{containername}/filesystem
	// Download file from container
	FileDownload(http.ResponseWriter, *http.Request)
	// FileUpload is the handler for POST /nodes/{nodeid}/containers/{containername}/filesystem
	// Upload file to container
	FileUpload(http.ResponseWriter, *http.Request)
	// GetContainerOSInfo is the handler for GET /nodes/{nodeid}/containers/{containername}/info
	// Get detailed information of the container OS
	GetContainerOSInfo(http.ResponseWriter, *http.Request)
	// KillContainerJob is the handler for DELETE /nodes/{nodeid}/containers/{containername}/jobs/{jobid}
	// Kills the job
	KillContainerJob(http.ResponseWriter, *http.Request)
	// GetContainerJob is the handler for GET /nodes/{nodeid}/containers/{containername}/jobs/{jobid}
	// Get details of a submitted job on the container
	GetContainerJob(http.ResponseWriter, *http.Request)
	// SendSignalToJob is the handler for POST /nodes/{nodeid}/containers/{containername}/jobs/{jobid}
	// Send signal to the job
	SendSignalJob(http.ResponseWriter, *http.Request)
	// KillAllContainerJobs is the handler for DELETE /nodes/{nodeid}/containers/{containername}/jobs
	// Kill all running jobs on the container
	KillAllContainerJobs(http.ResponseWriter, *http.Request)
	// ListContainerJobs is the handler for GET /nodes/{nodeid}/containers/{containername}/jobs
	// List running jobs on the container
	ListContainerJobs(http.ResponseWriter, *http.Request)
	// StartContainerJob is the handler for POST /nodes/{nodeid}/containers/{containername}/jobs
	// Start a new job in this container
	StartContainerJob(http.ResponseWriter, *http.Request)
	// GetContainerMemInfo is the handler for GET /nodes/{nodeid}/containers/{containername}/mem
	// Get detailed information about the memory in the container
	GetContainerMemInfo(http.ResponseWriter, *http.Request)
	// GetContainerNicInfo is the handler for GET /nodes/{nodeid}/containers/{containername}/nics
	// Get detailed information about the network interfaces in the container
	GetContainerNicInfo(http.ResponseWriter, *http.Request)
	// PingContainer is the handler for POST /nodes/{nodeid}/containers/{containername}/ping
	// Ping this container
	PingContainer(http.ResponseWriter, *http.Request)
	// KillContainerProcess is the handler for DELETE /nodes/{nodeid}/containers/{containername}/processes/{processid}
	// Kills the process by sending sigterm signal to the process. If it is still running, a sigkill signal will be sent to the process
	KillContainerProcess(http.ResponseWriter, *http.Request)
	// GetContainerProcess is the handler for GET /nodes/{nodeid}/containers/{containername}/processes/{processid}
	// Get process details
	GetContainerProcess(http.ResponseWriter, *http.Request)
	// SendSignalToProcess is the handler for POST /nodes/{nodeid}/containers/{containername}/processes/{processid}
	// Send signal to the process
	SendSignalProcess(http.ResponseWriter, *http.Request)
	// ListContainerProcesses is the handler for GET /nodes/{nodeid}/containers/{containername}/processes
	// Get running processes in this container
	ListContainerProcesses(http.ResponseWriter, *http.Request)
	// StartContainer is the handler for POST /nodes/{nodeid}/containers/{containername}/start
	// Start container instance
	StartContainer(http.ResponseWriter, *http.Request)
	// GetContainerState is the handler for GET /nodes/{nodeid}/containers/{containername}/state
	// Get aggregated consumption of container + all processes (CPU, memory, etc.)
	GetContainerState(http.ResponseWriter, *http.Request)
	// StopContainer is the handler for POST /nodes/{nodeid}/containers/{containername}/stop
	// Stop container instance
	StopContainer(http.ResponseWriter, *http.Request)
	// DeleteContainer is the handler for DELETE /nodes/{nodeid}/containers/{containername}
	// Delete container instance
	DeleteContainer(http.ResponseWriter, *http.Request)
	// GetContainer is the handler for GET /nodes/{nodeid}/containers/{containername}
	// Get container
	GetContainer(http.ResponseWriter, *http.Request)
	// ListContainers is the handler for GET /nodes/{nodeid}/containers
	// List running containers
	ListContainers(http.ResponseWriter, *http.Request)
	// CreateContainer is the handler for POST /nodes/{nodeid}/containers
	// Create a new container
	CreateContainer(http.ResponseWriter, *http.Request)
	// GetCPUInfo is the handler for GET /nodes/{nodeid}/cpus
	// Get detailed information of all CPUs in the node
	GetCPUInfo(http.ResponseWriter, *http.Request)
	// GetDiskInfo is the handler for GET /nodes/{nodeid}/disks
	// Get detailed information of all the disks in the node
	GetDiskInfo(http.ResponseWriter, *http.Request)
	// GetGWFWConfig is the handler for GET /nodes/{nodeid}/gws/{gwname}/advanced/firewall
	// Get current FW config
	GetGWFWConfig(http.ResponseWriter, *http.Request)
	// SetGWFWConfig is the handler for POST /nodes/{nodeid}/gws/{gwname}/advanced/firewall
	// Set FW config
	// Once used you can not use gw.portforwards any longer
	SetGWFWConfig(http.ResponseWriter, *http.Request)
	// GetGWHTTPConfig is the handler for GET /nodes/{nodeid}/gws/{gwname}/advanced/http
	// Get current HTTP config
	GetGWHTTPConfig(http.ResponseWriter, *http.Request)
	// SetGWHTTPConfig is the handler for POST /nodes/{nodeid}/gws/{gwname}/advanced/http
	// Set HTTP config
	// Once used you can not use gw.httpproxxies any longer
	SetGWHTTPConfig(http.ResponseWriter, *http.Request)
	// DeleteDHCPHost is the handler for DELETE /nodes/{nodeid}/gws/{gwname}/dhcp/{interface}/hosts/{macaddress}
	// Delete dhcp host
	DeleteDHCPHost(http.ResponseWriter, *http.Request)
	// ListGWDHCPHosts is the handler for GET /nodes/{nodeid}/gws/{gwname}/dhcp/{interface}/hosts
	// List DHCPHosts for specified interface
	ListGWDHCPHosts(http.ResponseWriter, *http.Request)
	// AddGWDHCPHost is the handler for POST /nodes/{nodeid}/gws/{gwname}/dhcp/{interface}/hosts
	// Add a dhcp host to a specified interface
	AddGWDHCPHost(http.ResponseWriter, *http.Request)
	// DeleteGWForward is the handler for DELETE /nodes/{nodeid}/gws/{gwname}/firewall/forwards/{forwardid}
	// Delete portforward, forwardid = srcip:srcport
	DeleteGWForward(http.ResponseWriter, *http.Request)
	// GetGWForwards is the handler for GET /nodes/{nodeid}/gws/{gwname}/firewall/forwards
	// Get list for IPv4 Forwards
	GetGWForwards(http.ResponseWriter, *http.Request)
	// CreateGWForwards is the handler for POST /nodes/{nodeid}/gws/{gwname}/firewall/forwards
	// Create a new Portforwarding
	CreateGWForwards(http.ResponseWriter, *http.Request)
	// DeleteHTTPProxies is the handler for DELETE /nodes/{nodeid}/gws/{gwname}/httpproxies/{proxyid}
	// Delete HTTP proxy
	DeleteHTTPProxies(http.ResponseWriter, *http.Request)
	// GetHTTPProxy is the handler for GET /nodes/{nodeid}/gws/{gwname}/httpproxies/{proxyid}
	// Get info of HTTP proxy
	GetHTTPProxy(http.ResponseWriter, *http.Request)
	// ListHTTPProxies is the handler for GET /nodes/{nodeid}/gws/{gwname}/httpproxies
	// List for HTTP proxies
	ListHTTPProxies(http.ResponseWriter, *http.Request)
	// CreateHTTPProxies is the handler for POST /nodes/{nodeid}/gws/{gwname}/httpproxies
	// Create new HTTP proxies
	CreateHTTPProxies(http.ResponseWriter, *http.Request)
	// StartGateway is the handler for POST /nodes/{nodeid}/gws/{gwname}/start
	// Start Gateway instance
	StartGateway(http.ResponseWriter, *http.Request)
	// StopGateway is the handler for POST /nodes/{nodeid}/gws/{gwname}/stop
	// Stop gateway instance
	StopGateway(http.ResponseWriter, *http.Request)
	// DeleteGateway is the handler for DELETE /nodes/{nodeid}/gws/{gwname}
	// Delete gateway instance
	DeleteGateway(http.ResponseWriter, *http.Request)
	// GetGateway is the handler for GET /nodes/{nodeid}/gws/{gwname}
	// Get gateway
	GetGateway(http.ResponseWriter, *http.Request)
	// UpdateGateway is the handler for PUT /nodes/{nodeid}/gws/{gwname}
	// Update Gateway
	UpdateGateway(http.ResponseWriter, *http.Request)
	// ListGateways is the handler for GET /nodes/{nodeid}/gws
	// List running gateways
	ListGateways(http.ResponseWriter, *http.Request)
	// CreateGW is the handler for POST /nodes/{nodeid}/gws
	// Create a new gateway
	CreateGW(http.ResponseWriter, *http.Request)
	// GetNodeOSInfo is the handler for GET /nodes/{nodeid}/info
	// Get detailed information of the OS of the node
	GetNodeOSInfo(http.ResponseWriter, *http.Request)
	// KillNodeJob is the handler for DELETE /nodes/{nodeid}/jobs/{jobid}
	// Kills the job
	KillNodeJob(http.ResponseWriter, *http.Request)
	// GetNodeJob is the handler for GET /nodes/{nodeid}/jobs/{jobid}
	// Get the details of a submitted job
	GetNodeJob(http.ResponseWriter, *http.Request)
	// KillAllNodeJobs is the handler for DELETE /nodes/{nodeid}/jobs
	// Kill all running jobs
	KillAllNodeJobs(http.ResponseWriter, *http.Request)
	// ListNodeJobs is the handler for GET /nodes/{nodeid}/jobs
	// List running jobs
	ListNodeJobs(http.ResponseWriter, *http.Request)
	// GetMemInfo is the handler for GET /nodes/{nodeid}/mem
	// Get detailed information about the memory in the node
	GetMemInfo(http.ResponseWriter, *http.Request)
	// GetNodeMounts is the handler for GET /nodes/{nodeid}/mounts
	// Get detailed information of all the mountpoints on the node
	GetNodeMounts(http.ResponseWriter, *http.Request)
	// GetNicInfo is the handler for GET /nodes/{nodeid}/nics
	// Get detailed information about the network interfaces in the node
	GetNicInfo(http.ResponseWriter, *http.Request)
	// PingNode is the handler for POST /nodes/{nodeid}/ping
	// Ping this node
	PingNode(http.ResponseWriter, *http.Request)
	// KillNodeProcess is the handler for DELETE /nodes/{nodeid}/processes/{processid}
	// Kills the process by sending sigterm signal to the process. If it is still running, a sigkill signal will be sent to the process
	KillNodeProcess(http.ResponseWriter, *http.Request)
	// GetNodeProcess is the handler for GET /nodes/{nodeid}/processes/{processid}
	// Get process details
	GetNodeProcess(http.ResponseWriter, *http.Request)
	// ListNodeProcesses is the handler for GET /nodes/{nodeid}/processes
	// Get processes
	ListNodeProcesses(http.ResponseWriter, *http.Request)
	// RebootNode is the handler for POST /nodes/{nodeid}/reboot
	// Immediately reboot the machine
	RebootNode(http.ResponseWriter, *http.Request)
	// GetNodeState is the handler for GET /nodes/{nodeid}/state
	// The aggregated consumption of node + all processes (cpu, memory, etc...)
	GetNodeState(http.ResponseWriter, *http.Request)
	// DeleteStoragePoolDevice is the handler for DELETE /nodes/{nodeid}/storagepools/{storagepoolname}/devices/{deviceuuid}
	// Removes the device from the storage pool
	DeleteStoragePoolDevice(http.ResponseWriter, *http.Request)
	// GetStoragePoolDeviceInfo is the handler for GET /nodes/{nodeid}/storagepools/{storagepoolname}/devices/{deviceuuid}
	// Get information of the device
	GetStoragePoolDeviceInfo(http.ResponseWriter, *http.Request)
	// ListStoragePoolDevices is the handler for GET /nodes/{nodeid}/storagepools/{storagepoolname}/devices
	// List the devices in the storage pool
	ListStoragePoolDevices(http.ResponseWriter, *http.Request)
	// CreateStoragePoolDevices is the handler for POST /nodes/{nodeid}/storagepools/{storagepoolname}/devices
	// Add extra devices to this storage pool
	CreateStoragePoolDevices(http.ResponseWriter, *http.Request)
	// RollbackFilesystemSnapshot is the handler for POST /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}/rollback
	// Rollback the file system to the state at the moment the snapshot was taken
	RollbackFilesystemSnapshot(http.ResponseWriter, *http.Request)
	// DeleteFilesystemSnapshot is the handler for DELETE /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}
	// Delete snapshot
	DeleteFilesystemSnapshot(http.ResponseWriter, *http.Request)
	// GetFilesystemSnapshotInfo is the handler for GET /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}
	// Get detailed information on the snapshot
	GetFilesystemSnapshotInfo(http.ResponseWriter, *http.Request)
	// ListFilesystemSnapshots is the handler for GET /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots
	// List snapshots of this file system
	ListFilesystemSnapshots(http.ResponseWriter, *http.Request)
	// CreateSnapshot is the handler for POST /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots
	// Create a new read-only snapshot of the current state of the vdisk
	CreateSnapshot(http.ResponseWriter, *http.Request)
	// DeleteFilesystem is the handler for DELETE /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}
	// Delete file system
	DeleteFilesystem(http.ResponseWriter, *http.Request)
	// GetFilesystemInfo is the handler for GET /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}
	// Get detailed file system information
	GetFilesystemInfo(http.ResponseWriter, *http.Request)
	// ListFilesystems is the handler for GET /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems
	// List all file systems
	ListFilesystems(http.ResponseWriter, *http.Request)
	// CreateFilesystem is the handler for POST /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems
	// Create a new file system
	CreateFilesystem(http.ResponseWriter, *http.Request)
	// DeleteStoragePool is the handler for DELETE /nodes/{nodeid}/storagepools/{storagepoolname}
	// Delete the storage pool
	DeleteStoragePool(http.ResponseWriter, *http.Request)
	// GetStoragePoolInfo is the handler for GET /nodes/{nodeid}/storagepools/{storagepoolname}
	// Get detailed information of this storage pool
	GetStoragePoolInfo(http.ResponseWriter, *http.Request)
	// ListStoragePools is the handler for GET /nodes/{nodeid}/storagepools
	// List storage pools present in the node
	ListStoragePools(http.ResponseWriter, *http.Request)
	// CreateStoragePool is the handler for POST /nodes/{nodeid}/storagepools
	// Create a new storage pool in the node
	CreateStoragePool(http.ResponseWriter, *http.Request)
	// GetVMInfo is the handler for GET /nodes/{nodeid}/vms/{vmid}/info
	// Get statistical information about the virtual machine.
	GetVMInfo(http.ResponseWriter, *http.Request)
	// MigrateVM is the handler for POST /nodes/{nodeid}/vms/{vmid}/migrate
	// Migrate the virtual machine to another host
	MigrateVM(http.ResponseWriter, *http.Request)
	// PauseVM is the handler for POST /nodes/{nodeid}/vms/{vmid}/pause
	// Pauses the VM
	PauseVM(http.ResponseWriter, *http.Request)
	// ResumeVM is the handler for POST /nodes/{nodeid}/vms/{vmid}/resume
	// Resumes the virtual machine
	ResumeVM(http.ResponseWriter, *http.Request)
	// ShutdownVM is the handler for POST /nodes/{nodeid}/vms/{vmid}/shutdown
	// Gracefully shutdown the virtual machine
	ShutdownVM(http.ResponseWriter, *http.Request)
	// StartVM is the handler for POST /nodes/{nodeid}/vms/{vmid}/start
	// Start the virtual machine
	StartVM(http.ResponseWriter, *http.Request)
	// StopVM is the handler for POST /nodes/{nodeid}/vms/{vmid}/stop
	// Stops the VM
	StopVM(http.ResponseWriter, *http.Request)
	// DeleteVM is the handler for DELETE /nodes/{nodeid}/vms/{vmid}
	// Deletes the virtual machine
	DeleteVM(http.ResponseWriter, *http.Request)
	// GetVM is the handler for GET /nodes/{nodeid}/vms/{vmid}
	// Get the virtual machine object
	GetVM(http.ResponseWriter, *http.Request)
	// UpdateVM is the handler for PUT /nodes/{nodeid}/vms/{vmid}
	// Updates the virtual machine
	UpdateVM(http.ResponseWriter, *http.Request)
	// ListVMs is the handler for GET /nodes/{nodeid}/vms
	// List all virtual machines
	ListVMs(http.ResponseWriter, *http.Request)
	// CreateVM is the handler for POST /nodes/{nodeid}/vms
	// Creates a new virtual machine
	CreateVM(http.ResponseWriter, *http.Request)
	// ExitZerotier is the handler for DELETE /nodes/{nodeid}/zerotiers/{zerotierid}
	// Exit the ZeroTier network
	ExitZerotier(http.ResponseWriter, *http.Request)
	// GetZerotier is the handler for GET /nodes/{nodeid}/zerotiers/{zerotierid}
	// Get ZeroTier network details
	GetZerotier(http.ResponseWriter, *http.Request)
	// ListZerotier is the handler for GET /nodes/{nodeid}/zerotiers
	// List running ZeroTier networks
	ListZerotier(http.ResponseWriter, *http.Request)
	// JoinZerotier is the handler for POST /nodes/{nodeid}/zerotiers
	// Join ZeroTier network
	JoinZerotier(http.ResponseWriter, *http.Request)
	// DeleteNode is the handler for DELETE /nodes/{nodeid}
	// Delete a node
	DeleteNode(http.ResponseWriter, *http.Request)
	// GetNode is the handler for GET /nodes/{nodeid}
	// Get detailed information of a node
	GetNode(http.ResponseWriter, *http.Request)
	// ListNodes is the handler for GET /nodes
	// List all nodes
	ListNodes(http.ResponseWriter, *http.Request)
	// DeleteDashboard is the handler for DELETE /graphs/{graphid}/dashboards/{dashboardid}
	// Delete dashboard
	DeleteDashboard(http.ResponseWriter, *http.Request)
	// GetDashboard is the handler for GET /graphs/{graphid}/dashboards/{dashboardid}
	// Get dashboard
	GetDashboard(http.ResponseWriter, *http.Request)
	// CreateDashboard is the handler for POST /graphs/{graphid}/dashboards
	// Create dashboard
	CreateDashboard(http.ResponseWriter, *http.Request)
	// ListDashboards is the handler for GET /graphs/{graphid}/dashboards
	// List all dashboards
	ListDashboards(http.ResponseWriter, *http.Request)
	// GetGraph is the handler for GET /graphs/{graphid}
	// Get a graph
	GetGraph(http.ResponseWriter, *http.Request)
	// ListGraphs is the handler for GET /graphs
	// List all graphs
	ListGraphs(http.ResponseWriter, *http.Request)
}

// NodesInterfaceRoutes is routing for /nodes root endpoint
func NodesInterfaceRoutes(r *mux.Router, i NodesInterface, org string) {
	r.HandleFunc("/nodes/{nodeid}/bridges/{bridgeid}", i.DeleteBridge).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/bridges/{bridgeid}", i.GetBridge).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/bridges", i.ListBridges).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/bridges", i.CreateBridge).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/cpus", i.GetContainerCPUInfo).Methods("GET")
	// r.HandleFunc("/nodes/{nodeid}/containers/{containername}/disks", i.GetContainerDiskInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/filesystem", i.FileDelete).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/filesystem", i.FileDownload).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/filesystem", i.FileUpload).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/info", i.GetContainerOSInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs/{jobid}", i.KillContainerJob).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs/{jobid}", i.GetContainerJob).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs/{jobid}", i.SendSignalJob).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs", i.KillAllContainerJobs).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs", i.ListContainerJobs).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs", i.StartContainerJob).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/mem", i.GetContainerMemInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/nics", i.GetContainerNicInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/ping", i.PingContainer).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/processes/{processid}", i.KillContainerProcess).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/processes/{processid}", i.GetContainerProcess).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/processes/{processid}", i.SendSignalProcess).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/processes", i.ListContainerProcesses).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/start", i.StartContainer).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/state", i.GetContainerState).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/stop", i.StopContainer).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}", i.DeleteContainer).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}", i.GetContainer).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers", i.ListContainers).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers", i.CreateContainer).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/cpus", i.GetCPUInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/disks", i.GetDiskInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/advanced/firewall", i.GetGWFWConfig).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/advanced/firewall", i.SetGWFWConfig).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/advanced/http", i.GetGWHTTPConfig).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/advanced/http", i.SetGWHTTPConfig).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/dhcp/{interface}/hosts/{macaddress}", i.DeleteDHCPHost).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/dhcp/{interface}/hosts", i.ListGWDHCPHosts).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/dhcp/{interface}/hosts", i.AddGWDHCPHost).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/firewall/forwards/{forwardid}", i.DeleteGWForward).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/firewall/forwards", i.GetGWForwards).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/firewall/forwards", i.CreateGWForwards).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/httpproxies/{proxyid}", i.DeleteHTTPProxies).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/httpproxies/{proxyid}", i.GetHTTPProxy).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/httpproxies", i.ListHTTPProxies).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/httpproxies", i.CreateHTTPProxies).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/start", i.StartGateway).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}/stop", i.StopGateway).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}", i.DeleteGateway).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}", i.GetGateway).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/gws/{gwname}", i.UpdateGateway).Methods("PUT")
	r.HandleFunc("/nodes/{nodeid}/gws", i.ListGateways).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/gws", i.CreateGW).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/info", i.GetNodeOSInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/jobs/{jobid}", i.KillNodeJob).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/jobs/{jobid}", i.GetNodeJob).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/jobs", i.KillAllNodeJobs).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/jobs", i.ListNodeJobs).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/mem", i.GetMemInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/mounts", i.GetNodeMounts).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/nics", i.GetNicInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/ping", i.PingNode).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/processes/{processid}", i.KillNodeProcess).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/processes/{processid}", i.GetNodeProcess).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/processes", i.ListNodeProcesses).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/reboot", i.RebootNode).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/state", i.GetNodeState).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/devices/{deviceuuid}", i.DeleteStoragePoolDevice).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/devices/{deviceuuid}", i.GetStoragePoolDeviceInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/devices", i.ListStoragePoolDevices).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/devices", i.CreateStoragePoolDevices).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}/rollback", i.RollbackFilesystemSnapshot).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}", i.DeleteFilesystemSnapshot).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}", i.GetFilesystemSnapshotInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots", i.ListFilesystemSnapshots).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots", i.CreateSnapshot).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}", i.DeleteFilesystem).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}", i.GetFilesystemInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems", i.ListFilesystems).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems", i.CreateFilesystem).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}", i.DeleteStoragePool).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}", i.GetStoragePoolInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools", i.ListStoragePools).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools", i.CreateStoragePool).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/info", i.GetVMInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/migrate", i.MigrateVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/pause", i.PauseVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/resume", i.ResumeVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/shutdown", i.ShutdownVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/start", i.StartVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/stop", i.StopVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}", i.DeleteVM).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}", i.GetVM).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}", i.UpdateVM).Methods("PUT")
	r.HandleFunc("/nodes/{nodeid}/vms", i.ListVMs).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/vms", i.CreateVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/zerotiers/{zerotierid}", i.ExitZerotier).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/zerotiers/{zerotierid}", i.GetZerotier).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/zerotiers", i.ListZerotier).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/zerotiers", i.JoinZerotier).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}", i.DeleteNode).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}", i.GetNode).Methods("GET")
	r.HandleFunc("/nodes", i.ListNodes).Methods("GET")
	r.HandleFunc("/graphs/{graphid}/dashboards/{dashboardid}", i.DeleteDashboard).Methods("DELETE")
	r.HandleFunc("/graphs/{graphid}/dashboards/{dashboardid}", i.GetDashboard).Methods("GET")
	r.HandleFunc("/graphs/{graphid}/dashboards", i.CreateDashboard).Methods("POST")
	r.HandleFunc("/graphs/{graphid}/dashboards", i.ListDashboards).Methods("GET")
	r.HandleFunc("/graphs/{graphid}", i.GetGraph).Methods("GET")
	r.HandleFunc("/graphs", i.ListGraphs).Methods("GET")
}
