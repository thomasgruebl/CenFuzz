package https_fuzzer

import (
	"log"

	"github.com/censoredplanet/CenFuzz/config"
	"github.com/censoredplanet/CenFuzz/util"
)

type ServernameAlternate struct{}

func (s *ServernameAlternate) Init(all bool) []*RequestWord {
	var requestWords []*RequestWord
	var requestWord *RequestWord
	retries := 0
	if !all {
		for i := 0; i < config.NumberOfProbesPerTest; i++ {
			ServerNameAlternate := util.GenerateServerNameAlternatives()
			requestWord = &RequestWord{
				Servername: ServerNameAlternate,
			}
			if containsRequestWord(requestWords, requestWord) {
				i--
				retries += 1
				if retries >= 10 {
					log.Println("[ServernameAlternate.Init] Could not find a new random value after 10 retries. Breaking.")
					break
				}
			} else {
				requestWords = append(requestWords, requestWord)
				retries = 0
			}
		}
	} else {
		servernameAllAlternatives := util.GenerateAllServerNameAlternatives()
		for _, servername := range servernameAllAlternatives {
			requestWords = append(requestWords, &RequestWord{Servername: servername})
		}
	}
	return requestWords
}

func (s *ServernameAlternate) Fuzz(target string, hostname string, requestWord RequestWord) (interface{}, interface{}, interface{}) {
	return MakeConnection(target, hostname, requestWord)
}
