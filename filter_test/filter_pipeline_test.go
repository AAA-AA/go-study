package filtertest

import (
	"fmt"
	"testing"
)

type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

func (f *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}

func TestStraightPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStraightPipeline("p1", spliter, converter, sum)
	ret, err := sp.Process("1,2,3")
	fmt.Printf("result %d", ret)
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("The expected is 6, but the actual is %d", ret)
	}
}
