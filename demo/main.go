package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/CaliDog/certstream-go"
	"github.com/pedroegsilva/gotagthem/tagger"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	maxGoroutines := 50
	limit := make(chan struct{}, maxGoroutines)

	gofindthemRules := map[string][]string{
		"google": {
			`"Google"`,
		},
		"log": {
			`"log"`,
		},
		"lets encrypt": {
			`"Let's Encrypt" or "Lets Encrypt"`,
		},
		"duck": {
			`"pato" or "duck"`,
		},
		"great": {
			`"good" or "great" or "awesome" or "nice"`,
		},
	}

	rules := map[string][]string{
		"google":       {`"google:data.source.name" and "log:data.source.name"`},
		"lets encrypt": {`"lets encrypt:data.source.name" and "log:data.source.name"`},
		"unknown":      {`not "google:data.source.name" and not "lets encrypt:data.source.name"`},
		"greatweb":     {`"great:data.leaf_cert.subject.CN" or "duck:great:data.leaf_cert.subject.CN"`},
	}

	gfte, err := NewGoFindThemTagger(gofindthemRules)
	if err != nil {
		panic(err)
	}

	goFindtagger, err := tagger.NewTaggerWithRules([]tagger.StringTagger{gfte}, nil, nil, rules)
	if err != nil {
		panic(err)
	}

	ipTagger := NewDnsIpTagger()
	taggerIp := tagger.NewTagger([]tagger.StringTagger{ipTagger}, nil, nil)

	// The false flag specifies that we want heartbeat messages.
	stream, errStream := certstream.CertStreamEventStream(true)
	avrg := float64(0)
	count := 0
	for {
		select {
		case jq := <-stream:
			start := time.Now()
			message, err := jq.Interface()
			if err != nil {
				errr := errors.Wrap(err, "")
				log.Error().Stack().Err(errr).Msg("Error decoding jq string")
				continue
			}

			rawJson, err := json.Marshal(message)
			if err != nil {
				errr := errors.Wrap(err, "")
				log.Error().Stack().Err(errr).Msg("Error marshalling json")
				continue
			}

			res, err := goFindtagger.ProcessJson(string(rawJson), goFindtagger.GetFieldNames(), nil)
			if err != nil {
				errr := errors.Wrap(err, "")
				log.Error().Stack().Err(errr).Msg("Error processing json")
				continue
			}

			duration := time.Since(start)
			count++
			avrg = (float64(duration.Nanoseconds()) - avrg) / float64(count)
			if count%1000 == 0 {
				log.Info().Msg(fmt.Sprintf("avrg operation time from last 1000: %f nanoseconds", avrg))
				count = 0
				avrg = 0
			}

			for rule, expressions := range res {
				if rule != "greatweb" {
					continue
				}
				log.Info().Msg(fmt.Sprintf("buffer: %d/%d", len(limit), maxGoroutines))
				limit <- struct{}{} // allocate the resource
				go processResult(rule, expressions, string(rawJson), taggerIp, limit)
			}

		case err := <-errStream:
			errr := errors.Wrap(err, "")
			log.Error().Stack().Err(errr)
		}
	}
}

func processResult(rule string, expressions []string, rawJson string, taggerIp *tagger.Tagger, limit chan struct{}) {
	result := ""
	fieldsInfo, err := taggerIp.TagJson(rawJson, []string{"data.leaf_cert.subject.CN"}, nil)
	if err != nil {
		errr := errors.Wrap(err, "")
		log.Error().Stack().Err(errr).Msg("Error tagging ip")
		return
	}
	result += fmt.Sprintf("%s%s%s (expressions: %s)\n", Red, rule, Reset, expressions)
	for _, fieldInfo := range fieldsInfo {
		for _, info := range fieldInfo.Taggers {
			if len(info.Tags) > 0 {
				domain := info.RunData.(string)
				result += fmt.Sprintf("%s%s%s\n", Yellow, domain, Reset)
				for _, tag := range info.Tags {
					result += fmt.Sprintf("%s    %s%s\n", Yellow, tag, Reset)
				}
			}
		}
	}
	result += fmt.Sprintf("%s%s%s\n", Green, string(rawJson), Reset)
	fmt.Println(result)
	<-limit // release the resource
}
