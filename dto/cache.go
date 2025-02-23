package dto

// UserVerifiedAccountCache represents the cache for a verified user account
type UserVerifiedAccountCache struct {
	Username   string       `json:"username"`
	UID        string       `json:"uid"`
	FirstName  string       `json:"firstName"`
	LastName   string       `json:"lastName"`
	PhoneNo    string       `json:"phoneNo"`
	Email      string       `json:"email"`
	Businesses []Businesses `json:"businesses"`
}

// Businesses represents the businesses of a user
type Businesses struct {
	Name             string         `json:"name"`
	IndustryType     string         `json:"industryType"`
	BusinessType     string         `json:"businessType"`
	Description      string         `json:"description"`
	PhoneNo          string         `json:"phoneNo"`
	OperatingHours   OperatingHours `json:"operatingHours"`
	Address          string         `json:"address"`
	BusinessImageURL string         `json:"businessImageUrl"`
	OwnerUsername    string         `json:"ownerUsername"`
}
