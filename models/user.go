package models

type User struct {
	UserID               int64   `json:"user_id"`
	UserTypeID           *int64  `json:"user_type_id"`
	Name                 string  `json:"name"`
	DtOfCommenceBusiness *string `json:"dt_of_commence_business"`
	MobileNum            string  `json:"mobile_num"`
	Email                *string `json:"email"`
	Address              *string `json:"address"`
	Pincode              *string `json:"pincode"`
	Location             *string `json:"location"`
	BusinessLicenseNo    *string `json:"business_license_no"`
	Validity             *string `json:"validity"`
	GstNo                *string `json:"gst_no"`
	ExpiryDt             *string `json:"expiry_dt"`
	BusinessName         *string `json:"business_name"`
	BusinessType         *string `json:"business_type"`
	MandiID              *int64  `json:"mandi_id"`
	MandiTypeID          *string `json:"mandi_type_id"`
	Remarks              *string `json:"remarks"`
	Col1                 *string `json:"col1"`
}
