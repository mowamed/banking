package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Mohamed Bakayoko", "Accra", "00233", "2000-01-01", "1"},
		{"1001", "Kouame Kwaku", "Kumassi", "00233", "2000-01-01", "1"},
		{"1001", "Franck Kone", "Abidjan", "00225", "2000-01-01", "1"},
		{"1001", "Moussa Dicko", "Bamako", "00223", "2000-01-01", "1"},
		{"1001", "Aliuone Thiam", "Dakar", "00226", "2000-01-01", "1"},
	}

	return CustomerRepositoryStub{customers}
}
