package pet

type Pet struct {
	Name    string
	Type    string
	Sex     string
	Address address
	Age     string
	Weight  string
	Race    string
}

type address struct {
	Number string
	City   string
	Street string
}
