package models

const UserMediaPath string = "media"

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

// swagger:parameters Language
type Language string

const (
	TR  Language = "tr"
	ENG Language = "eng"
)

var Languages = []Language{
	TR, ENG,
}

func (c Language) Valid() bool {
	for _, v := range Languages {
		if c == v {
			return true
		}
	}
	return false
}
