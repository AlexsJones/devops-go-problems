package types

type Letter interface {
  // GetStreetAddress returns the complete street address
  GetStreetAddress() string

  // GetState returns the unabbreviated state nanme
  GetState() string

  // GetPostcode returns the numerical postcode
  GetPostcode() int
}
