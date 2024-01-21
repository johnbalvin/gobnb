package details

type metadataData2 struct {
	Data metadataData_data2 `json:"data"`
}
type metadataData_data2 struct {
	Presentation presentation2 `json:"presentation"`
}
type presentation2 struct {
	StayProductDetailPage stayProductDetailPage `json:"stayProductDetailPage"`
}
type stayProductDetailPage struct {
	Sections section7 `json:"sections"`
}
type section7 struct {
	Section  []section8 `json:"sections"`
	Metadata metadata5  `json:"metadata"`
}
type section8 struct {
	Section9 section9 `json:"section"`
}
type section9 struct {
	Typename               string                 `json:"__typename"`
	StructuredDisplayPrice structuredDisplayPrice `json:"structuredDisplayPrice"`
}
type structuredDisplayPrice struct {
	PrimaryLine primaryLine `json:"primaryLine"`
}
type primaryLine struct {
	Price     string `json:"price"`
	Qualifier string `json:"qualifier"`
}
type metadata5 struct {
	BookingPrefetchData bookingPrefetchData `json:"bookingPrefetchData"`
}
type bookingPrefetchData struct {
	P3_display_rate p3_display_rate `json:"p3_display_rate"`
}
type p3_display_rate struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}
