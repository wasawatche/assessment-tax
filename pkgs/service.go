package main

type Tax struct {
	TotalIncome float64      `json:"totalIncome"`
	Wht         float64      `json:"wht"`
	Allowances  []Allowances `json:"allowances"`
}

type Allowances struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}



func calculateTax(t Tax) struct {

}

func calIncome(income float64) float64 {

}

func calAllowance (a Allowances) float64 {
	switch a.AllowanceType {
		case "donation":
			  
	}
}