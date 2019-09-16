package chart

import (
	"fmt"
	"testing"
	"time"

	liquid "github.com/egapool/go-liquid"
)

func TestExample(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	startFrom, err := time.ParseInLocation("20060102150405", "20190908225000", loc)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	err = SaveOHLC(startFrom)
	if err != nil {
		fmt.Println(err)
	}
}

func TestTrimEnd(t *testing.T) {
	res := ExecutionsResponse{
		Executions: []liquid.Execution{
			liquid.Execution{CreatedAt: 1},
			liquid.Execution{CreatedAt: 1},
			liquid.Execution{CreatedAt: 1},
			liquid.Execution{CreatedAt: 2},
			liquid.Execution{CreatedAt: 2},
			liquid.Execution{CreatedAt: 3},
			liquid.Execution{CreatedAt: 3},
		},
	}
	actual := res.TrimEnd().Executions
	expected := []liquid.Execution{
		liquid.Execution{CreatedAt: 1},
		liquid.Execution{CreatedAt: 1},
		liquid.Execution{CreatedAt: 1},
		liquid.Execution{CreatedAt: 2},
		liquid.Execution{CreatedAt: 2},
	}
	if len(actual) != len(expected) {
		t.Fatalf("Each length will got: %v\nwant: %v", len(actual), len(expected))
	}
	for i, e := range actual {
		if e.CreatedAt != expected[i].CreatedAt {
			t.Fatalf("#%v got: %v\nwant: %v", i, e, expected[i])
		}
	}
}
