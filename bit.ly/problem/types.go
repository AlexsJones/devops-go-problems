package problem

// Link contains all the information regarding shortening process
type Link interface {
	// Expired should only return true iff its TTL has passed
	Expired() bool
	// GetURL returns the original URL that was used to create the "shorten" link
	GetURL() string
	// GetShortLink returns the "shorten" link that is generated from the given url
	GetShortLink() string
}

// Store holds all references to current links and updates
// them when required.
type Store interface {
	// AddLink will create a new Link object
	// and return that link.
	// An error should be returned the url already exists.
	AddLink(url string) (Link, error)
	// GetLink will return the Link object stored internally
	// It should report an error if it doesn't exist.
	GetLink(link string) (Link, error)
}
