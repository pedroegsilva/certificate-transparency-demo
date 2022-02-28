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
	}

	rules := map[string][]string{
		"google":       {`"google:data.source.name" and "log:data.source.name"`},
		"lets encrypt": {`"lets encrypt:data.source.name" and "log:data.source.name"`},
		"unknown":      {`not "google:data.source.name" and not "lets encrypt:data.source.name"`},
	}

	gfte, err := NewGoFindThemTagger(gofindthemRules)
	if err != nil {
		panic(err)
	}

	stringTaggers := []tagger.StringTagger{gfte}
	intTaggers := []tagger.IntTagger{}
	floatTaggers := []tagger.FloatTagger{}

	tagger, err := tagger.NewTaggerWithRules(stringTaggers, intTaggers, floatTaggers, rules)
	if err != nil {
		panic(err)
	}

	// The false flag specifies that we want heartbeat messages.
	stream, errStream := certstream.CertStreamEventStream(true)
	for {
		select {
		case jq := <-stream:
			message, err := jq.Interface()
			if err != nil {
				errr := errors.Wrap(err, "")
				log.Error().Stack().Err(errr).Msg("Error decoding jq string")
				continue
			}

			rawJson, err := json.Marshal(message)
			if err != nil {
				errr := errors.Wrap(err, "")
				log.Error().Stack().Err(errr).Msg("Error decoding jq string")
				continue
			}

			res, err := tagger.ProcessJson(string(rawJson), tagger.GetFieldNames(), nil)
			if err != nil {
				errr := errors.Wrap(err, "")
				log.Error().Stack().Err(errr).Msg("Error decoding jq string")
				continue
			}
			for rule, expressions := range res {
				if rule == "google" {
					continue
				}
				fmt.Printf("%s%s%s (expressions: %s)\n", Red, rule, Reset, expressions)
				fmt.Println(Green, string(rawJson), Reset)
			}

		case err := <-errStream:
			errr := errors.Wrap(err, "")
			log.Error().Stack().Err(errr)
		}
	}

}
