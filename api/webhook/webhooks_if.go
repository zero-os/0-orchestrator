package webhook

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"net/http"

	"github.com/gorilla/mux"
)

// WebhooksInterface is interface for /webhooks root endpoint
type WebhooksInterface interface { // DeleteWebhook is the handler for DELETE /webhooks/{webhookname}
	// Delete a webhook
	DeleteWebhook(http.ResponseWriter, *http.Request)
	// GetWebhook is the handler for GET /webhooks/{webhookname}
	// Get a webhook
	GetWebhook(http.ResponseWriter, *http.Request)
	// UpdateWebhook is the handler for PUT /webhooks/{webhookname}
	// Update a webhook
	UpdateWebhook(http.ResponseWriter, *http.Request)
	// ListWebhooks is the handler for GET /webhooks
	// List all webhooks
	ListWebhooks(http.ResponseWriter, *http.Request)
	// CreateWebhook is the handler for POST /webhooks
	// Create Webhook
	CreateWebhook(http.ResponseWriter, *http.Request)
}

// WebhooksInterfaceRoutes is routing for /webhooks root endpoint
func WebhooksInterfaceRoutes(r *mux.Router, i WebhooksInterface) {
	r.HandleFunc("/webhooks/{webhookname}", i.DeleteWebhook).Methods("DELETE")
	r.HandleFunc("/webhooks/{webhookname}", i.GetWebhook).Methods("GET")
	r.HandleFunc("/webhooks/{webhookname}", i.UpdateWebhook).Methods("PUT")
	r.HandleFunc("/webhooks", i.ListWebhooks).Methods("GET")
	r.HandleFunc("/webhooks", i.CreateWebhook).Methods("POST")
}
