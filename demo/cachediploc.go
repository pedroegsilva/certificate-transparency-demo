package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type CacheIpLocWithExpir struct {
	ipByCacheElement map[string]CacheElement
	maxTTL           time.Duration
	client           *http.Client
	leaser           chan time.Time
	requestSleep     time.Duration
}

type CacheElement struct {
	IpLocation string
	LastLookup time.Time
}

func NewCacheIpLocWithExpir(maxTTL time.Duration, requestSleep time.Duration, requestTimeout time.Duration) *CacheIpLocWithExpir {
	leaser := make(chan time.Time, 1)
	client := &http.Client{
		Timeout: requestTimeout,
	}
	leaser <- time.Now()
	return &CacheIpLocWithExpir{
		ipByCacheElement: make(map[string]CacheElement),
		maxTTL:           maxTTL,
		client:           client,
		leaser:           leaser,
		requestSleep:     requestSleep,
	}
}

// uses https://about.ip2c.org/#about
func (ci *CacheIpLocWithExpir) GetIpLocation(ip string) (location string, err error) {
	if cacheElem, ok := ci.ipByCacheElement[ip]; ok {
		if time.Since(cacheElem.LastLookup) < ci.maxTTL {
			cacheElem.LastLookup = time.Now()
			return cacheElem.IpLocation, nil
		}
	}

	location, err = ci.getIpLocation(ip)
	if err != nil {
		location = "unknown"
	}

	ci.ipByCacheElement[ip] = CacheElement{
		IpLocation: location,
		LastLookup: time.Now(),
	}

	return
}

func (ci *CacheIpLocWithExpir) CleanCache() {
	for ip, cacheElem := range ci.ipByCacheElement {
		if time.Since(cacheElem.LastLookup) > ci.maxTTL {
			delete(ci.ipByCacheElement, ip)
		}
	}
}

func (ci *CacheIpLocWithExpir) getIpLocation(ip string) (location string, err error) {
	lastRequest := <-ci.leaser

	// verify if the ip was already looked up while waiting for the "lock"
	if cacheElem, ok := ci.ipByCacheElement[ip]; ok {
		if time.Since(cacheElem.LastLookup) < ci.maxTTL {
			cacheElem.LastLookup = time.Now()
			ci.leaser <- lastRequest
			return cacheElem.IpLocation, nil
		}
	}

	elapsedTime := time.Since(lastRequest)
	if elapsedTime < ci.requestSleep {
		log.Debug().Msg(fmt.Sprintf("wating %d ms to make the request\n", (ci.requestSleep - elapsedTime).Milliseconds()))
		time.Sleep(ci.requestSleep - elapsedTime)
	}
	now := time.Now()
	ci.leaser <- now
	log.Debug().Msg(fmt.Sprintf("requesting at %s", now.String()))
	resp, err := ci.client.Get(fmt.Sprintf("https://ip2c.org/?ip=%s", ip))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	bodySplit := strings.Split(string(body), ";")
	location = bodySplit[len(bodySplit)-1]
	return
}
