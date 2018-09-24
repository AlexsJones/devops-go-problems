package types

type Letter interface {
	// GetStreetAddress returns the complete street address
	GetStreetAddress() string

	// GetState returns the unabbreviated state nanme
	GetState() string

	// GetPostcode returns the numerical postcode
	GetPostcode() int

	// Valid will check to see if it is a valid letter and if it can
	// be sent to the correct address
	Valid() bool
}
