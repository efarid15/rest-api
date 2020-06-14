package main

type CustomerModel struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func main()  {
	GetData()
}

func GetData() (dc interface{}) {
	var dataCustomer interface{} = InsertCustomer()
	return dataCustomer
}

func InsertCustomer() (customers []CustomerModel)  {
	customers = append(customers, CustomerModel{FirstName: "Enal", LastName: "Farid"})
	customers = append(customers, CustomerModel{FirstName: "Samsul", LastName: "Efendi"})
	customers = append(customers, CustomerModel{FirstName: "John", LastName: "Doe"})
	customers = append(customers, CustomerModel{FirstName: "Dani", LastName: "Setiawan"})
	return customers
}