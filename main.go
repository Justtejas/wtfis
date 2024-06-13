package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Definition struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
	Example    string   `json:"example"`
}

type Phonetic struct {
	Text      string  `json:"text"`
	Audio     string  `json:"audio"`
	SourceURL string  `json:"sourceUrl"`
	License   License `json:"license"`
}

type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Meaning struct {
	Definitions  []Definition `json:"definitions"`
	PartOfSpeech string       `json:"partOfSpeech"`
}
type Response struct {
	Word       string     `json:"word"`
	Phonetic   string     `json:"phonetic"`
	Phonetics  []Phonetic `json:"phonetics"`
	Meaning    []Meaning  `json:"meanings"`
	License    License    `json:"license"`
	SourceURLs []string   `json:"sourceUrls"`
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Usage : wtfis word")
		fmt.Println("Usage : wtfis \"word word\"")
		os.Exit(1)
	}
	wordToSearch := os.Args[1:2]
	apiToReq :=  fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", wordToSearch[0])
	res, err := handleRequest(apiToReq)	
	if err != nil {
		panic(err)
	}
	printDefinition(res)      
}

func handleRequest(url string) ([]Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var jsonData string
	for scanner.Scan() {
		jsonData += scanner.Text()
	}

	var res []Response
	err = json.Unmarshal([]byte(jsonData), &res)
	if err != nil {
		fmt.Println("Cannot find the meaning of the given word.")
	}
	return res, nil
}

func printDefinition(data []Response) {
	for _, res := range data {
		fmt.Printf("word : %s\n", res.Word)
		for _, meaning := range res.Meaning {
			for _, definition := range meaning.Definitions {
				fmt.Printf("Definition : %s\n", definition.Definition)
				if definition.Example != "" {
					fmt.Printf("Example : %s\n", definition.Example)
				}
				if len(definition.Antonyms) != 0 {
					for _, antonym := range definition.Antonyms {
						fmt.Printf(" Antonym : %s ", antonym)
					}
				}
				if len(definition.Synonyms) != 0 {
					for _, synonym := range definition.Synonyms {
						fmt.Printf(" Synonym : %s ", synonym)
					}
				}
				fmt.Println()
			}
		}
	}
}
