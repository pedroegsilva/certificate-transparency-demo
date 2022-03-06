package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	requestSleep   = time.Second * 2
	requestTimeout = time.Second * 10
	cacheTTl       = time.Second * 30
)

// DnsIpTagger is a string tagger that uses the gofindthem library to tag.
type DnsIpTagger struct {
	resolver *net.Resolver
	cache    *CacheIpLocWithExpir
}

// NewDnsIpTagger initializes the DnsIpTagger
func NewDnsIpTagger() *DnsIpTagger {
	return &DnsIpTagger{
		resolver: net.DefaultResolver,
		cache:    NewCacheIpLocWithExpir(cacheTTl, requestSleep, requestTimeout),
	}
}

// IsValid all non empty texts are valid for the GoFindThemTagger
func (dit *DnsIpTagger) IsValid(data string) bool {
	//TODO do proper validation of a domain
	return data != ""
}

// GetTags gets the tags for the given data and returns the on the runData the expressions
// that were matched by their tags
func (dit *DnsIpTagger) GetTags(data string) (tags []string, runData interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	ips, err := dit.resolver.LookupIP(ctx, "ip4", data)
	if err != nil {
		log.Debug().Msg(fmt.Sprintf("failed dns %s", data))
		return nil, nil, nil
	}

	for _, ip := range ips {
		ipStr := ip.String()
		// location := "unknown"
		location, err := dit.cache.GetIpLocation(ipStr)
		if err != nil {
			location = "unknown"
		}
		tags = append(tags, fmt.Sprintf("%s (%s)", ipStr, location))
	}

	return tags, data, nil
}

// GetName returns the string 'gofindthem'
func (dit *DnsIpTagger) GetName() string {
	return "iptagger"
}
