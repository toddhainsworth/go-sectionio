package sectionio

// Account represents a Section.IO Account object
type Account struct{}

// Application represents a Section.IO Application object
type Application struct{}

// Billing represents a Section.IO Billing object
type Billing struct{}

// Domain represents a Section.IO Domain object
type Domain struct{}

// Environment represents a Section.IO Environment object
type Environment struct{}

// Help represents a Section.IO Help object
type Help struct{}

// OutagePage represents a Section.IO OutagePage object
type OutagePage struct{}

// Proxy represents a Section.IO Proxy object
type Proxy struct{}

// Synthetics represents a Section.IO Synthetics object
type Synthetics struct{}

// User represents a Section.IO User object
type User struct{}

// Zone represents a Section.IO Zone object
type Zone struct{}

// Session represents a session within the API
// The Section.IO API doesn't strictly keep track
// of sessions so this is more of an internal thing than anything
type Session struct {
	Username string
	Password []byte
}
