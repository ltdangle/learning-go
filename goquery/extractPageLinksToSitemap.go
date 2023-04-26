package main

// SiteMap data structure
type SiteMap struct {
	Links []*SiteLink
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
	Scheme  string
	Host    string
	SiteMap *SiteMap
}

func NewSiteMapper(scheme string, host string) *SiteMapper {
	return &SiteMapper{
		Scheme:  scheme,
		Host:    host,
		SiteMap: &SiteMap{},
	}
}

func (m *SiteMapper) addPathToSiteMap(host string, path string) {
	if m.Host != host {
		return
	} else {
		if m.pathExists(path) == false {
			m.SiteMap.Links = append(m.SiteMap.Links, &SiteLink{Path: path, Status: LinkStatus{Visited: false}})
		}
	}
}
func (m *SiteMapper) pathExists(path string) bool {
	for _, v := range m.SiteMap.Links {
		if v.Path == path {
			return true
		}
	}
	return false
}

func (m *SiteMapper) uncrawledLinksRemain() bool {
	for _, link := range m.SiteMap.Links {
		if link.Status.Visited == false {
			return true
		}
	}
	return false
}

func (m *SiteMapper) nextUncrawledLink() *SiteLink {
	for _, link := range m.SiteMap.Links {
		if link.Status.Visited == false {
			return link
		}
	}
	return nil
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
	for url := range hostLinks.Urls {
		siteMap.Links = append(siteMap.Links, &SiteLink{Path: url, Status: LinkStatus{Visited: false}})

	}
}
