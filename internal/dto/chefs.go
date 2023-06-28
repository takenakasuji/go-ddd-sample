package dto

type Chefs struct {
	Chefs []*Chef `json:"chef"`
}

func NewChefs() *Chefs {
	return &Chefs{
		Chefs: []*Chef{},
	}
}

func (c *Chefs) Add(chef *Chef) {
	c.Chefs = append(c.Chefs, chef)
}
