package dto

// CreateNewBusinessReq represents the request body for creating a new business
type CreateNewBusinessReq struct {
	Name             string         `json:"name" validate:"required"`
	IndustryType     string         `json:"industryType" validate:"required"`
	BusinessType     string         `json:"businessType" validate:"required"`
	Description      string         `json:"description"`
	PhoneNo          string         `json:"phoneNo" validate:"required"`
	OperatingHours   OperatingHours `json:"operatingHours"`
	Address          string         `json:"address"`
	BusinessImageURL string         `json:"businessImageUrl"`
	OwnerUsername    string         `json:"ownerUsername" validate:"required"`
}

// OperatingHours represents the operating hours of a business for the week
type OperatingHours struct {
	Monday    OpenTime `json:"monday"`
	Tuesday   OpenTime `json:"tuesday"`
	Wednesday OpenTime `json:"wednesday"`
	Thursday  OpenTime `json:"thursday"`
	Friday    OpenTime `json:"friday"`
	Saturday  OpenTime `json:"saturday"`
	Sunday    OpenTime `json:"sunday"`
}

// OpenTime represents the open and close time of a business
type OpenTime struct {
	Open      bool   `json:"open"`
	OpenTime  string `json:"openTime,omitempty"`
	CloseTime string `json:"closeTime,omitempty"`
}

// CreateNewBusinessRes represents the response body for creating a new business
type CreateNewBusinessRes struct {
	BusinessID   int    `json:"businessId"`
	BusinessName string `json:"businessName"`
	Success      bool   `json:"success"`
}
