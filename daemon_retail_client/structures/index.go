package structures

type MessageFromNSQ struct {
	ID  string
	New bool
}

type MindboxClient struct {
	Customer struct {
		Ids struct {
			RetailCRMID string `json:"retailCRMID"`
			WebsiteID   string `json:"websiteID"`
		} `json:"ids,omitempty"`
		BirthDate    string `json:"birthDate,omitempty"`
		Password     string `json:"password,omitempty"`
		FullName     string `json:"fullName,omitempty"`
		MobilePhone  string `json:"mobilePhone,omitempty"`
		Email        string `json:"email,omitempty"`
		CustomFields struct {
			Address string `json:"address"`
		} `json:"customFields,omitempty"`
		Subscriptions []struct {
			IsSubscribed bool `json:"isSubscribed"`
		} `json:"subscriptions"`
	} `json:"customer"`
	ExecutionDateTimeUtc string `json:"executionDateTimeUtc,omitempty"`
	New                  bool   `json:"new"`
}
