package customermodel

type CustomerModel struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func InsertCustomer() (customers []CustomerModel)  {
	customers = append(customers, CustomerModel{FirstName: "Enal", LastName: "Farid"})
	customers = append(customers, CustomerModel{FirstName: "Samsul", LastName: "Efendi"})
	customers = append(customers, CustomerModel{FirstName: "John", LastName: "Doe"})
	return
}

