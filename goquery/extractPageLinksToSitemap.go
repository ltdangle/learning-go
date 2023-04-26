package main

// SiteMap data structure
type SiteMap struct {
	Links []SiteLink
}

type SiteLink struct {
	Path   string
	Status LinkStatus
}

type LinkStatus struct {
	Visited bool
}

// SiteMapper object.
type SiteMapper struct {
	Host    string
	SiteMap *SiteMap
}

func NewSiteMapper(host string) *SiteMapper {
	return &SiteMapper{
		Host:    host,
		SiteMap: &SiteMap{},
	}
}

func (m *SiteMapper) addLinkToSiteMap(host string, path string) {
	if m.Host != host {
		return
	} else {
		m.SiteMap.Links = append(m.SiteMap.Links, SiteLink{Path: path, Status: LinkStatus{Visited: false}})
	}
}

// TODO: move to SiteMapper.
func extractPageLinksToSitemap(siteMap *SiteMap, page Page, host string) {
	var hostLinks *HostLinks
	for _, value := range page.Links {
		if value.Host == host {
			hostLinks = &value
			break
		}
	}
	// Host links not found, return empty value.
	if hostLinks == nil {
		return
	}
	for url, _ := range hostLinks.Urls {
		siteMap.Links = append(siteMap.Links, SiteLink{Path: url, Status: LinkStatus{Visited: false}})

	}
}
