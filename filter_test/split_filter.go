package filtertest

import (
	"errors"
	"strings"
)

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, errors.New("")
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
