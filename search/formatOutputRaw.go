package search

type root struct {
	Data rootData `json:"data"`
}
type rootData struct {
	Presentation rootdataPresentation `json:"presentation"`
}

type rootdataPresentation struct {
	StaysSearch rootdatapresentationStaysSearch `json:"staysSearch"`
}

type rootdatapresentationStaysSearch struct {
	Results rootdatapresentationstayssearchResults `json:"results"`
}

type rootdatapresentationstayssearchResults struct {
	PaginationInfo paginationInfo `json:"paginationInfo"`
	Results        []resultData   `json:"searchResults"`
}

type paginationInfo struct {
	PageCursors        []string `json:"pageCursors"`
	PreviousPageCursor string   `json:"previousPageCursor"`
	NextPageCursor     string   `json:"nextPageCursor"`
}
type resultData struct {
	Listing      listingData     `json:"listing"`
	PricingQuote pricingWrapper1 `json:"pricingQuote"`
}

type pricingWrapper1 struct {
	StructuredStayDisplayPrice pricingWrapper2 `json:"structuredStayDisplayPrice"`
}

type pricingWrapper2 struct {
	PrimaryLine     priceData    `json:"primaryLine"`
	SecondaryLine   priceData    `json:"secondaryLine"`
	ExplanationData priceDetails `json:"explanationData"`
}

type listingData struct {
	AvgRatingA11yLabel string                  `json:"avgRatingA11yLabel"`
	AvgRatingLocalized string                  `json:"avgRatingLocalized"`
	City               string                  `json:"city"`
	ContextualPictures []pricture              `json:"contextualPictures"`
	Coordinate         coordinate              `json:"coordinate"`
	FormattedBadges    []formatedbadgeWrapper1 `json:"formattedBadges"`
	Id                 int64                   `json:"id,string"`
	ListingObjType     string                  `json:"listingObjType"`
	LocalizedCityName  string                  `json:"localizedCityName"`
	Name               string                  `json:"name"`
	PdpUrlType         string                  `json:"pdpUrlType"`
	RoomTypeCategory   string                  `json:"roomTypeCategory"`
	TierId             int                     `json:"tierId"`
	Title              string                  `json:"title"`
	TitleLocale        string                  `json:"titleLocale"`
}
type pricture struct {
	Picture string `json:"picture"`
}

type coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type formatedbadgeWrapper1 struct {
	LoggingContext formatedbadgeWrapper2 `json:"loggingContext"`
}

type formatedbadgeWrapper2 struct {
	BadgeType string `json:"badgeType"`
}

type priceData struct {
	DisplayComponentType string `json:"displayComponentType"`
	AccessibilityLabel   string `json:"accessibilityLabel"`
	Price                string `json:"price"`
	OriginalPrice        string `json:"originalPrice"`
	DiscountedPrice      string `json:"discountedPrice"`
	Qualifier            string `json:"qualifier"`
	ShortQualifier       string `json:"shortQualifier"`
	ConcatQualifierLeft  bool   `json:"concatQualifierLeft"`
}

type priceDetails struct {
	PriceDetails []items `json:"priceDetails"`
}

type items struct {
	Items []itemData `json:"items"`
}

type itemData struct {
	DisplayComponentType string `json:"displayComponentType"`
	Description          string `json:"description"`
	PriceString          string `json:"priceString"`
}
