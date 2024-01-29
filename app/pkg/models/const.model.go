package models

// swagger:model Gender
type Gender string

const (
	Female    Gender = "female"
	Male      Gender = "male"
	NonBinary Gender = "non-binary"
	Intersex  Gender = "intersex"
)

var Genders = []Gender{
	Female, Male, NonBinary, Intersex,
}

func (c Gender) Valid() bool {
	for _, v := range Genders {
		if c == v {
			return true
		}
	}
	return false
}
