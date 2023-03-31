package cdn

import (
	"net/url"
)

func (c *CDN) Sign(from string) (signed *url.URL, err error) {
	if parsed, pe := url.Parse(from); nil != pe {
		err = pe
	} else if _domain, ee := c.lookupDomain(parsed.Host); nil != ee {
		err = ee
	} else {
		err = _domain.sign(parsed)
	}

	return
}
