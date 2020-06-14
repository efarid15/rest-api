package customer

import (
	pluginaccess "goproject/helper"
	respond "goproject/res"
	"net/http"
)

// get customer function
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	var dataCustomer, data interface{}
	dataCustomer = pluginaccess.PluginAccess("lib/customer.so", "GetData", data)

	if dataCustomer == nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}
	respond.RespondWithJSON(w, http.StatusOK, dataCustomer)
}
