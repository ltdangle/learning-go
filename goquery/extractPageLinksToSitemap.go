package main

type SiteMap struct {
	Links []SiteLink
}

type SiteLink struct {
	Url    string
	Status LinkStatus
}

type LinkStatus struct {
	Visited bool
}

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
		siteMap.Links = append(siteMap.Links, SiteLink{Url: url, Status: LinkStatus{Visited: false}})

	}
}
