package details

import "path/filepath"

func (meta metadataData) standardize() Data {
	ev := meta.Data.Presentation.StayProductDetailPage.Sections.Metadata.LoggingContext.EventDataLogging
	data := Data{
		Coordinates:    Coordinates{Latitude: ev.Lat, Longitud: ev.Long},
		RoomType:       ev.RoomType,
		IsSuperHost:    ev.IsSuperhost,
		HomeTier:       ev.HomeTier,
		PersonCapacity: ev.PersonCapacity,
		Rating: Rating{
			Accuracy:          ev.AccuracyRating,
			Checking:          ev.CheckinRating,
			CleaningLiness:    ev.CleanlinessRating,
			Comunication:      ev.CommunicationRating,
			Location:          ev.LocationRating,
			Value:             ev.LocationRating,
			GuestSatisfaction: ev.GuestSatisfactionOverall,
			ReviewCount:       ev.VisibleReviewCount,
		},
	}
	sd := meta.Data.Presentation.StayProductDetailPage.Sections.SbuiData
	for _, section := range sd.SectionConfiguration.Root.Sections {
		switch section.SectionData.Typename {
		case "PdpHostOverviewDefaultSection":
			data.Host.ID = section.SectionData.HostAvatar.LoggingEventData.EventData.PdpContext.HostID
			data.Host.Name = section.SectionData.Title

		case "PdpOverviewV2Section":
			data.SubDescription.Title = section.SectionData.Title
			for _, item := range section.SectionData.OverviewItems {
				data.SubDescription.Items = append(data.SubDescription.Items, item.Title)
			}
		}
	}
	for _, section := range meta.Data.Presentation.StayProductDetailPage.Sections.Sections {
		switch section.Section.Typename {
		case "HostProfileSection":
			if section.Section.HostAvatar.UserID != "" {
				data.Host.ID = section.Section.HostAvatar.UserID
			}
			if section.Section.Title != "" {
				data.Host.Name = section.Section.Title
			}
			data.Host.JoinedOn = section.Section.Subtitle
			data.Host.Description = section.Section.HostProfileDescription.HtmlText
			for _, cohost := range section.Section.AdditionalHosts {
				data.CoHosts = append(data.CoHosts, Cohost{
					ID:   cohost.ID,
					Name: cohost.Name,
				})
			}
		case "PhotoTourModalSection":
			for _, mediaItem := range section.Section.MediaItems {
				img := Img{
					Title:     mediaItem.AccessibilityLabel,
					URL:       mediaItem.BaseURL,
					Extension: filepath.Ext(mediaItem.BaseURL),
				}
				data.Images = append(data.Images, img)
			}
		case "PoliciesSection":
			for _, houseRulesSection := range section.Section.HouseRules {
				houseRule := HouseRule{
					Title: houseRulesSection.Title,
				}
				for _, item := range houseRulesSection.Items {
					if item.Title == "Additional rules" {
						data.HouseRules.Aditional = item.HTML.HTMLText
						continue
					}
					houseRule.Values = append(houseRule.Values, HouseRuleValue{
						Title: item.Title,
						Icon:  item.Icon,
					})
				}
				data.HouseRules.General = append(data.HouseRules.General, houseRule)
			}
		case "LocationSection":
			for _, locationDetail := range section.Section.SeeAllLocationDetails {
				seeAllLocationDetail := LocationDetail{
					Title:   locationDetail.Title,
					Content: locationDetail.Content.HTMLText,
				}
				data.LocationDescriptions = append(data.LocationDescriptions, seeAllLocationDetail)
			}
		case "PdpTitleSection":
			data.Title = section.Section.Title
		case "PdpHighlightsSection":
			for _, highliting := range section.Section.Highlights {
				highliting := Highlight{
					Title:    highliting.Title,
					Subtitle: highliting.Subtitle,
					Icon:     highliting.ICon,
				}
				data.Highlights = append(data.Highlights, highliting)
			}
		case "PdpDescriptionSection":
			data.Description = section.Section.HtmlDescription.HtmlText
		case "AmenitiesSection":
			for _, amenityGroupRaw := range section.Section.SeeAllAmenitiesGroups {
				amenityGroup := AmenityGroup{
					Title: amenityGroupRaw.Title,
				}
				for _, amenityRaw := range amenityGroupRaw.Amenity {
					amenity := Amenity{
						Title:     amenityRaw.Title,
						Subtitle:  amenityRaw.Subtitle,
						Icon:      amenityRaw.Icon,
						Available: amenityRaw.Available,
					}
					amenityGroup.Values = append(amenityGroup.Values, amenity)
				}
				data.Amenities = append(data.Amenities, amenityGroup)
			}
		}
	}

	return data
}
