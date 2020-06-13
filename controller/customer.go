package customer

import (
	"net/http"
)
import "goproject/model"
import "goproject/res"

// get customer function
func GetCustomer(w http.ResponseWriter, r *http.Request) {

	// access customer model
	datacustomer := customermodel.InsertCustomer()

	if datacustomer == nil {
		// respond if error
		respond.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}

	// respond if success endpoint
	respond.RespondWithJSON(w, http.StatusOK, datacustomer)

}
