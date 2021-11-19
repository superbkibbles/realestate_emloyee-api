package employee

const (
	STATUS_ACTIVE   = "active"
	STATUS_DEACTIVE = "deactive"
)

type Employee struct {
	ID          string `json:"id"`
	AgencyID    string `json:"agency_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Photo       string `json:"photo"`
	Gender      string `json:"gender"`
	GPS         gps    `json:"gps"`
	City        string `json:"city"`
	Status      string `json:"status"`
	PhoneNumber string `json:"phone_number"`
	Position    string `json:"position"`
	DateCreated string `json:"created_at"`
}

type gps struct {
	Long string `json:"long"`
	Lat  string `json:"lat"`
}

type Employees []Employee
