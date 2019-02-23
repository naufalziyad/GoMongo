package models

type User struct {
	Name string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	PersonalInfo `json:"personalinfo,omitempty"`
	CampaignInfo `json:"campaigninfo,omitempty"`
}

type PersonalInfo struct {
	Address string `json:"address,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	MobilePhone string `json:"mobilephone,omitempty"`
}

type CampaignInfo struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	BannerUrl string `json:"bannerurl,omitempty"`
	Budget int `json:"budget,omitempty"`
}
