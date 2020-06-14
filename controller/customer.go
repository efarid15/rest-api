package customer

import (
	"errors"
	respond "goproject/res"
	"net/http"
	"os"
	"plugin"
)

// get customer function
func GetCustomer(w http.ResponseWriter, r *http.Request) {

	plug, err := plugin.Open("lib/customer.so")
	if err != nil {
		println(err)
		os.Exit(1)
	}

	symCustomer, err := plug.Lookup("GetData")
	if err != nil {
		println("Salah simbol")
		os.Exit(1)
	}

	//fmt.Printf("%T\n", symCustomer)
	pgPtr, ok := symCustomer.(func() interface{})
	if !ok {
		panic(errors.New("gagal binding ke plugin interface"))
	}

	var datacustomer interface{}
	datacustomer = pgPtr()

	if datacustomer == nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}


	respond.RespondWithJSON(w, http.StatusOK, datacustomer)

}
