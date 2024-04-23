package details

type Data struct {
	Title                string
	URL                  string
	RoomType             string
	Language             string
	HomeTier             int
	PersonCapacity       int
	IsSuperHost          bool
	Price                Price
	Rating               Rating
	Coordinates          Coordinates
	Host                 Host
	CoHosts              []Cohost
	SubDescription       SubDescription
	Description          string
	Highlights           []Highlight
	Amenities            []AmenityGroup
	HouseRules           HouseRules
	LocationDescriptions []LocationDetail
	Images               []Img
}
type Price struct {
	Amount         float32
	CurrencySymbol string
	Qualifier      string
}
type Cohost struct {
	ID   string
	Name string
}
type Host struct {
	ID          string
	Name        string
	JoinedOn    string
	Description string
}
type HouseRules struct {
	Aditional string
	General   []HouseRule
}
type Img struct {
	Title       string
	URL         string
	ContentType string
	Extension   string
	Content     []byte `json:"-"`
}
type HouseRule struct {
	Title  string
	Values []HouseRuleValue
}
type HouseRuleValue struct {
	Title string
	Icon  string
}
type LocationDetail struct {
	Title   string
	Content string
}
type Rating struct {
	Accuracy          float32
	Checking          float32
	CleaningLiness    float32
	Comunication      float32
	Location          float32
	Value             float32
	GuestSatisfaction float32
	ReviewCount       int
}
type Coordinates struct {
	Latitude float64
	Longitud float64
}
type SubDescription struct {
	Title string
	Items []string
}
type AmenityGroup struct {
	Title  string
	Values []Amenity
}
type Amenity struct {
	Title     string
	Subtitle  string
	Available bool
	Icon      string
}
type Highlight struct {
	Title    string
	Subtitle string
	Icon     string
}
