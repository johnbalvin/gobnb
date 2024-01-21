package search

type searchRequest struct {
	OperationName string     `json:"operationName"`
	Variables     variables  `json:"variables"`
	Extensions    extensions `json:"extensions"`
}

type variables struct {
	IncludeMapResults       bool                    `json:"includeMapResults"`
	IsLeanTreatment         bool                    `json:"isLeanTreatment"`
	StaysSearchRequest      staysSearchRequest      `json:"staysSearchRequest"`
	StaysMapSearchRequestV2 staysMapSearchRequestV2 `json:"staysMapSearchRequestV2"`
}

type staysSearchRequest struct {
	Cursor            string     `json:"cursor"`
	RequestedPageType string     `json:"requestedPageType"`
	MetadataOnly      bool       `json:"metadataOnly"`
	Source            string     `json:"source"`
	SearchType        string     `json:"searchType"`
	TreatmentFlags    []string   `json:"treatmentFlags"`
	RawParams         []rawParam `json:"rawParams"`
	MaxMapItems       int        `json:"maxMapItems"`
}

type staysMapSearchRequestV2 struct {
	Cursor            string     `json:"cursor"`
	RequestedPageType string     `json:"requestedPageType"`
	MetadataOnly      bool       `json:"metadataOnly"`
	Source            string     `json:"source"`
	SearchType        string     `json:"searchType"`
	TreatmentFlags    []string   `json:"treatmentFlags"`
	RawParams         []rawParam `json:"rawParams"`
}

type rawParam struct {
	FilterName   string   `json:"filterName"`
	FilterValues []string `json:"filterValues"`
}

type extensions struct {
	PersistedQuery persistedQuery `json:"persistedQuery"`
}

type persistedQuery struct {
	Version    int    `json:"version"`
	Sha256Hash string `json:"sha256Hash"`
}
