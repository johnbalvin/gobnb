package details

import "encoding/json"

type niobeMinimalClientDataWrapper struct {
	NiobeMinimalClientData [][]json.RawMessage `json:"niobeMinimalClientData"`
}
type PriceDependencyInput struct {
	ProducID    string
	ImpresionID string
	ApiKey      string
}
type metadataData struct {
	Data      metadataData_data `json:"data"`
	Variables metadataVariables `json:"variables"`
}
type metadataData_data struct {
	Presentation metadataData_data_presentation `json:"presentation"`
}
type metadataData_data_presentation struct {
	Typename              string                                               `json:"__typename"`
	StayProductDetailPage metadataData_data_presentation_stayProductDetailPage `json:"stayProductDetailPage"`
}
type metadataData_data_presentation_stayProductDetailPage struct {
	Sections sections `json:"sections"`
}
type sections struct {
	Typename string       `json:"__typename"`
	Sections []subSection `json:"sections"`
	Metadata metadata     `json:"metadata"`
	SbuiData SbuiData     `json:"sbuiData"`
}
type loggingContext struct {
	EventDataLogging eventDataLogging `json:"eventDataLogging"`
}
type eventDataLogging struct {
	RoomType                 string  `json:"roomType"`
	Lat                      float64 `json:"listingLat"`
	HomeTier                 int     `json:"homeTier"`
	Long                     float64 `json:"listingLng"`
	PersonCapacity           int     `json:"personCapacity"`
	DescriptionLanguage      string  `json:"descriptionLanguage"`
	IsSuperhost              bool    `json:"isSuperhost"`
	AccuracyRating           float32 `json:"accuracyRating"`
	CheckinRating            float32 `json:"checkinRating"`
	CleanlinessRating        float32 `json:"cleanlinessRating"`
	CommunicationRating      float32 `json:"communicationRating"`
	LocationRating           float32 `json:"locationRating"`
	ValueRating              float32 `json:"valueRating"`
	GuestSatisfactionOverall float32 `json:"guestSatisfactionOverall"`
	VisibleReviewCount       int     `json:"visibleReviewCount,string"`
}
type metadata struct {
	Typename       string         `json:"__typename"`
	LoggingContext loggingContext `json:"loggingContext"`
}
type SbuiData struct {
	Typename             string               `json:"__typename"`
	SectionConfiguration sectionConfiguration `json:"sectionConfiguration"`
}
type sectionConfiguration struct {
	Typename string `json:"__typename"`
	Root     root   `json:"root"`
}
type root struct {
	Typename string     `json:"__typename"`
	Sections []section4 `json:"sections"`
}
type section4 struct {
	Typename    string      `json:"__typename"`
	SectionId   string      `json:"sectionId"`
	SectionData sectionData `json:"sectionData"`
}
type sectionData struct {
	Typename      string          `json:"__typename"`
	Title         string          `json:"title"`
	HostAvatar    hostAvatar      `json:"hostAvatar"`
	OverviewItems []overviewItems `json:"overviewItems"`
}
type overviewItems struct {
	Typename string `json:"__typename"`
	Title    string `json:"title"`
}
type subSection struct {
	ID       string   `json:"id"`
	Typename string   `json:"__typename"`
	Section  section3 `json:"section"`
}
type SeeAllAmenitiesGroup struct {
	Typename string    `json:"__typename"`
	Title    string    `json:"title"`
	Amenity  []amenity `json:"amenities"`
}
type amenity struct {
	Typename  string `json:"__typename"`
	ID        string `json:"id"`
	Available bool   `json:"available"`
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Icon      string `json:"icon"`
}
type section3 struct {
	Typename               string                 `json:"__typename"`
	Title                  string                 `json:"title"`
	Subtitle               string                 `json:"subtitle"`
	Highlights             []highlights           `json:"highlights"`
	SeeAllAmenitiesGroups  []SeeAllAmenitiesGroup `json:"seeAllAmenitiesGroups"`
	HtmlDescription        htmlDescription        `json:"htmlDescription"`
	HouseRules             []houseRule            `json:"houseRulesSections"`
	SeeAllLocationDetails  []seeAllLocationDetail `json:"seeAllLocationDetails"`
	MediaItems             []mediaItem            `json:"mediaItems"`
	HostAvatar             hostAvatar             `json:"hostAvatar"`
	AdditionalHosts        []cohostAvatar         `json:"additionalHosts"`
	HostProfileDescription hostProfileDescription `json:"hostProfileDescription"`
}
type hostProfileDescription struct {
	HtmlText string `json:"htmlText"`
}
type cohostAvatar struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
type hostAvatar struct {
	Title            string           `json:"title"`
	Subtitle         string           `json:"subtitle"`
	Badge            string           `json:"badge"`
	LoggingEventData loggingEventData `json:"loggingEventData"`
	UserID           string           `json:"userId"`
}
type loggingEventData struct {
	EventData eventData `json:"eventData"`
}
type eventData struct {
	PdpContext pdpContext `json:"pdpContext"`
}
type pdpContext struct {
	HostID string `json:"hostId"`
}
type mediaItem struct {
	AccessibilityLabel string `json:"accessibilityLabel"`
	BaseURL            string `json:"baseUrl"`
}
type seeAllLocationDetail struct {
	Title   string          `json:"title"`
	Content contentLocation `json:"content"`
}
type contentLocation struct {
	HTMLText string `json:"htmlText"`
}
type houseRule struct {
	Title string `json:"title"`
	Items []Item `json:"items"`
}
type htmlDescription struct {
	Typename string `json:"__typename"`
	HtmlText string `json:"htmlText"`
}
type highlights struct {
	Typename     string       `json:"__typename"`
	Title        string       `json:"title"`
	Subtitle     string       `json:"subtitle"`
	SubtitleHTML subtitleHtml `json:"subtitleHtml"`
	ICon         string       `json:"icon"`
}
type subtitleHtml struct {
	HtmlText string `json:"htmlText"`
}
type Item struct {
	Typename string `json:"__typename"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Icon     string `json:"icon"`
	HTML     HTML   `json:"html"`
}
type HTML struct {
	Typename string `json:"__typename"`
	HTMLText string `json:"htmlText"`
}
type DescriptionItems struct {
	Typename string `json:"__typename"`
	Title    string `json:"title"`
}
type metadataVariables struct {
	ID                 string             `json:"id"`
	PdpSectionsRequest PdpSectionsRequest `json:"pdpSectionsRequest"`
}
type metadatExtension struct {
	PersistedQuery persistedQuery `json:"persistedQuery"`
}
type persistedQuery struct {
	Version    int    `json:"version"`
	Sha256Hash string `json:"sha256Hash"`
}
type PdpSectionsRequest struct {
	Adults                        string   `json:"adults"`
	BypassTargetings              bool     `json:"bypassTargetings"`
	CategoryTag                   *string  `json:"categoryTag"`
	CauseId                       *string  `json:"causeId"`
	Children                      *string  `json:"children"`
	DisasterId                    *string  `json:"disasterId"`
	DiscountedGuestFeeVersion     *string  `json:"discountedGuestFeeVersion"`
	DisplayExtensions             *string  `json:"displayExtensions"`
	FederatedSearchId             *string  `json:"federatedSearchId"`
	ForceBoostPriorityMessageType *string  `json:"forceBoostPriorityMessageType"`
	Infants                       *string  `json:"infants"`
	InteractionType               *string  `json:"interactionType"`
	Layouts                       []string `json:"layouts"`
	Pets                          int      `json:"pets"`
	PdpTypeOverride               *string  `json:"pdpTypeOverride"`
	PhotoId                       *string  `json:"photoId"`
	Preview                       bool     `json:"preview"`
	PreviousStateCheckIn          *string  `json:"previousStateCheckIn"`
	PreviousStateCheckOut         *string  `json:"previousStateCheckOut"`
	PriceDropSource               *string  `json:"priceDropSource"`
	PrivateBooking                bool     `json:"privateBooking"`
	PromotionUuid                 *string  `json:"promotionUuid"`
	RelaxedAmenityIds             *string  `json:"relaxedAmenityIds"`
	SearchId                      *string  `json:"searchId"`
	SelectedCancellationPolicyId  *string  `json:"selectedCancellationPolicyId"`
	SelectedRatePlanId            *string  `json:"selectedRatePlanId"`
	SplitStays                    *string  `json:"splitStays"`
	StaysBookingMigrationEnabled  bool     `json:"staysBookingMigrationEnabled"`
	TranslateUgc                  *string  `json:"translateUgc"`
	UseNewSectionWrapperApi       bool     `json:"useNewSectionWrapperApi"`
	SectionIds                    []string `json:"sectionIds"`
	CheckIn                       *string  `json:"checkIn"`
	CheckOut                      *string  `json:"checkOut"`
	P3ImpressionId                string   `json:"p3ImpressionId"`
}
