package pet

type Pet struct {
	Name    string
	Type    string
	Sex     string
	Address address
	Age     uint8
	Weight  uint8
	Race    string
}

type address struct {
	Number uint8
	City   string
	Street string
}
