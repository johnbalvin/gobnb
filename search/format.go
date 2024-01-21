package search

type Data struct {
	RoomID           int64
	Badges           []string
	Name             string
	Title            string
	Type             string
	Kind             string
	Category         string
	Rating           Rating
	Coordinates      Coordinates
	Fee              Fee
	Price            PriceData
	LongStayDiscount Price
	Images           []Img
}
type Fee struct {
	Cleaning Price
	Airbn    Price
}
type PriceData struct {
	Total     Price
	Unit      UnitPrice
	BreakDown []PriceBreakDown
}
type PriceBreakDown struct {
	Description    string
	Amount         float32
	CurrencySymbol string
}
type Rating struct {
	Value       float32
	ReviewCount int
}
type Coordinates struct {
	Latitude float64
	Longitud float64
}

type Img struct {
	URL         string
	ContentType string
	Extension   string
	Content     []byte `json:"-"`
}
type UnitPrice struct {
	Qualifier string
	Discount  float32
	Price
}
type Price struct {
	Amount         float32
	CurrencySymbol string
}
