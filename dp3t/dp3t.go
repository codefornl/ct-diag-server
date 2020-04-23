package dp3t

// ExposedOverview - A list of all SecretKeys
type ExposedOverview struct {
	Exposed []Exposee `json:"exposed,omitempty"`
}

// ExposeeAuthData - Authentication data used to verify the test result (base64 encoded)
type ExposeeAuthData struct {
	Value string `json:"value,omitempty"`
}

// ExposeeRequest - The request that adds an Exposee to the system
type ExposeeRequest struct {
	// The SecretKey used to generate EphID base64 encoded.
	Key string `json:"key"`
	// The onset date of the secret key. Format: yyyy-MM-dd
	Onset string `json:"onset"`

	AuthData *ExposeeAuthData `json:"authData"`
}

// Exposee - A Exposed Identity
type Exposee struct {
	// The SecretKey of a exposed as a base64 encoded string. The SecretKey consists of 32 bytes.
	Key string `json:"key"`
	// The onset of an exposed.
	Onset string `json:"onset"`
}
