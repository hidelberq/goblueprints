package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/hidelbreq/goblueprints/chapter4/synonyms/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	if apiKey == "" {
		log.Fatalln("環境変数 BHT_APIKEY が空です")
	}
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%qの類語検索に失敗しました: %v\n", word, err)
		}

		if len(syns) == 0 {
			log.Fatalf("%sには類義語はありませんでした", word)
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
