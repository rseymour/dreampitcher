package main

import (
	"bytes"
	"code.google.com/p/gcfg"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/rseymour/markov"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func GetApi() *anaconda.TwitterApi {
	type Config struct {
		Api struct {
			ConsumerKey       string
			ConsumerSecret    string
			AccessToken       string
			AccessTokenSecret string
		}
	}
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
		panic(err)
	}
	anaconda.SetConsumerKey(cfg.Api.ConsumerKey)
	anaconda.SetConsumerSecret(cfg.Api.ConsumerSecret)
	api := anaconda.NewTwitterApi(cfg.Api.AccessToken, cfg.Api.AccessTokenSecret)
	return api
}

func GetVerb() string {
	verbs := []string{"i dreamed",
		"i wish",
		"i hope"}
	rand.Seed(time.Now().UTC().UnixNano())
	return verbs[rand.Intn(len(verbs))]
}

func main() {
	api := GetApi()
	v := url.Values{}
	v.Set("count", "200")
	var sourceText bytes.Buffer
	var verb string
	re := regexp.MustCompile(`@.*?\W`)
	for {
		if sourceText.Len() > 24000 {
			sourceText.Reset()
		}
		verb = GetVerb()
		search_result, err := api.GetSearch("\""+verb+"\" -RT", v)
		if err != nil {
			panic(err)
		}
		for _, tweet := range search_result.Statuses {
			sourceText.WriteString(strings.ToLower(tweet.Text))
			sourceText.WriteString(" ")
		}
		sample := re.ReplaceAllString(sourceText.String(), " ")
		for i := 0; i < 1200; i++ {
			sample = strings.SplitAfterN(sample, " ", 2)[1]
			fmt.Println(markov.Generate(26, 2, sample))
			time.Sleep(75000 * time.Microsecond)
		}
	}
}
