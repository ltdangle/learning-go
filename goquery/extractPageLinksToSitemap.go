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

func extractPageLinksToSitemap(p Page, host string) SiteMap {
	var hostLinks *HostLinks
	for _, value := range p.Links {
		if value.Host == host {
			hostLinks = &value
			break
		}
	}
	// Host links not found, return empty value.
	if hostLinks == nil {
		return SiteMap{}
	}
	var sitemap = SiteMap{}
	for url, _ := range hostLinks.Urls {
		if url == "" {
			url = "/"
		}
		sitemap.Links = append(sitemap.Links, SiteLink{Url: url, Status: LinkStatus{Visited: false}})

	}
	return sitemap
}
