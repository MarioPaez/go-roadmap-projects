package model

const (
	TEMPERATURE Types = "temperature"
	WEIGHT      Types = "weight"
	LENGTH      Types = "length"
)

type Types string

type Conversion struct {
	From  string
	Value float64
	To    string
	Type  Types
}
