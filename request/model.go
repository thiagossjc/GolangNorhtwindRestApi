package request

type RequestOrder struct {
	Id                 int64   `json:"id"`
	IdCustomer         int64   `json:"idCustomer"`
	Customer           string  `json:"customer"`
	IdSupplier         int64   `json:"idSupplier"`
	Supplier           string  `json:"supplier"`
	DateRequest        string  `json:"dateRequest"`
	TimeRequest        string  `json:"timeRequest"`
	TotalQuantity      int64   `json:"totalQuantity"`
	TotalPrice         float32 `json:"totalPrice"`
	TotPriceWidhoutTax float32 `json:"totPriceWidhoutTax"`
	Tax                float32 `json:"tax"`
	Other_tax          float32 `json:"otherTax"`
	Country            string  `json:"country"`
	DocState           string  `json:"docState"`
	Email              string  `json:"email"`
	Cellphone          string  `json:"cellphone"`
	Description        string  `json:"description"`
	DtReg              string  `json:"dtReg"`
	IdUser             int64   `json:"idUser"`
	Cancel             bool    `json:"cancel"`
	IdCancel           int64   `json:"idCancel"`
	DeliverySupplier   bool    `json:"deliverySupplier"`
	IdStatus           int32   `json:"idStatus"`
	DescriptStatus     string  `json:"descriptStatus"`
	Id_event           int32   `json:"idEvent"`
	Event              string  `json:"event"`
}

func (req RequestOrder) SumTotalPrice(priceWTax float32, tax float32, otherTax float32) float32 {
	var total = priceWTax + tax + otherTax
	return total
}
