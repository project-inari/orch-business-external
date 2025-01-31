package dto

// CoreBusinessServerCreateReq represents the request body for creating a new business with core business server API
type CoreBusinessServerCreateReq struct {
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

// CoreBusinessServerCreateRes represents the response body for creating a new business with core business server API
type CoreBusinessServerCreateRes struct {
	BusinessID   int    `json:"businessId"`
	BusinessName string `json:"businessName"`
	Success      bool   `json:"success"`
}
