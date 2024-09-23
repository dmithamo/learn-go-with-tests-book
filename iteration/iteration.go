package iteration

import (
	"fmt"
)

func RepeatPhrase(phrase string, times int) string {
	repeated := phrase
	for i := 0; i < times-1; i++ {
		repeated += fmt.Sprintf("%s", phrase)
	}

	return repeated
}
