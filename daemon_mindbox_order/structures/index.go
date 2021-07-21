package structures

type MessageFromNSQ struct {
	Customer struct {
		IDs struct {
			WebsiteID   string `json:"websiteID,omitempty"`
			RetailCRMID string `json:"retailCRMID,omitempty"`
		} `json:"ids"`
		FullName    string `json:"fullName,omitempty"`
		Email       string `json:"email,omitempty"`
		MobilePhone string `json:"mobilePhone,omitempty"`
	} `json:"customer"`
	Order struct {
		IDs struct {
			MindboxId int    `json:"mindboxId,omitempty"`
			WebsiteID string `json:"websiteID,omitempty"`
		} `json:"ids"`
		DeliveryCost float64 `json:"deliveryCost,omitempty"`
		CustomFields struct {
			DeliveryType    string `json:"deliveryType,omitempty"`
			DeliveryAddress string `json:"deliveryAddress,omitempty"`
			DeliveryDate    string `json:"deliveryDate,omitempty"`
			DeliveryTime    string `json:"deliveryTime,omitempty"`
			Name            string `json:"name,omitempty"`
			ShipmentStore   string `json:"shipmentStore,omitempty"`
			Status          string `json:"status,omitempty"`
		} `json:"customFields,omitempty"`
		TotalPrice float64 `json:"totalPrice,omitempty"`
		Discounts  []struct {
			Type      string `json:"type,omitempty"`
			PromoCode struct {
				IDs struct {
					Code string `json:"code,omitempty"`
				} `json:"ids,omitempty"`
			} `json:"promoCode,omitempty"`
			Amount float64 `json:"amount,omitempty"`
		} `json:"discounts,omitempty"`
		Payments []struct {
			Type   string  `json:"type,omitempty"`
			Amount float64 `json:"amount,omitempty"`
		} `json:"payments,omitempty"`
		Lines []struct {
			MinPricePerItem        float64     `json:"minPricePerItem,omitempty"`
			CostPricePerItem       float64     `json:"costPricePerItem,omitempty"`
			BasePricePerItem       float64     `json:"basePricePerItem,omitempty"`
			Quantity               interface{} `json:"quantity,omitempty"`
			QuantityType           string      `json:"quantityType,omitempty"`
			DiscountedPricePerLine float64     `json:"discountedPricePerLine,omitempty"`
			LineID                 string      `json:"lineId,omitempty"`
			LineNumber             int         `json:"lineNumber,omitempty"`
			Discounts              []struct {
				Type                string `json:"type,omitempty"`
				ExternalPromoAction struct {
					IDs struct {
						ExternalId string `json:"externalId,omitempty"`
					} `json:"ids,omitempty"`
				} `json:"externalPromoAction,omitempty"`
				Amount float64 `json:"amount,omitempty"`
			} `json:"discounts,omitempty"`
			Product struct {
				IDs struct {
					Website string `json:"website,omitempty"`
					ERP     string `json:"eRP,omitempty"`
				} `json:"ids,omitempty"`
			} `json:"product,omitempty"`
			Status string `json:"status,omitempty"`
		} `json:"lines,omitempty"`
		Email       string `json:"email,omitempty"`
		MobilePhone string `json:"mobilePhone,omitempty"`
	} `json:"order"`
	PointOfContact       string `json:"pointOfContact,omitempty"`
	ExecutionDateTimeUtc string `json:"executionDateTimeUtc,omitempty"`
	New                  bool   `json:"new,omitempty"`
}
