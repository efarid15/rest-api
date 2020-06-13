package customer

import (
	"net/http"
)
import "goproject/model"
import "goproject/res"

func GetCustomer(w http.ResponseWriter, r *http.Request) {

	datacustomer := customermodel.InsertCustomer()

	if datacustomer == nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}
	respond.RespondWithJSON(w, http.StatusOK, datacustomer)

}
