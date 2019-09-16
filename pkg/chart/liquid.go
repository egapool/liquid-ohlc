package chart

import (
	"fmt"
	"strconv"
	"time"

	"github.com/egapool/go-liquid"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

type Agent struct {
	client *liquid.Liquid
}

type ExecutionsResponse struct {
	Executions liquid.Executions
}

func (e *ExecutionsResponse) First() (execution liquid.Execution) {
	return e.Executions[0]
}

func (e *ExecutionsResponse) Last() (execution liquid.Execution) {
	return e.Executions[len(e.Executions)-1]
}

func (e *ExecutionsResponse) TrimEnd() (response ExecutionsResponse) {
	last := e.Last()
	// last.CreatedAtと同じ秒数になるindexを探して、それ以前を返す
	for i, ex := range e.Executions {
		if ex.CreatedAt == last.CreatedAt {
			return ExecutionsResponse{Executions: e.Executions[:i]}
		}
	}
	return *e
}

func SaveOHLC(startFrom time.Time) (err error) {
	liquid := liquid.New(API_KEY, API_SECRET)
	liquid.SetDebug(true)
	fmt.Println(startFrom)

	//
	//file, err := os.OpenFile("sample.csv", os.O_WRONLY|os.O_CREATE, 0600)
	//defer file.Close()
	db := NewConnect()
	defer db.Close()
	//writer := csv.NewWriter(file)

	start := startFrom.Unix()
	minute := 0
	var ohlc OHLC
	for {
		executions, err := liquid.GetExecutions("BTCJPY", start, 1000, 1)
		if err != nil {
			return err
		}
		res := ExecutionsResponse{Executions: executions}
		fmt.Println(len(res.Executions))
		res = res.TrimEnd()
		fmt.Println(len(res.Executions))
		for _, ex := range res.Executions {
			this := time.Unix(ex.CreatedAt, 0)
			if minute != 0 && this.Minute() != minute {
				ohlc.Range = 60
				ohlc.Timestamp = time.Date(this.Year(), this.Month(), this.Day(), this.Hour(), minute, 0, 0, this.Location()).Add(time.Minute).Unix()

				// save
				fmt.Println(ohlc)
				// writer.Write(ohlc.ToArray())
				db.Insert(ohlc)

				// reset
				ohlc = OHLC{}
			}
			minute = this.Minute()
			price, _ := strconv.ParseFloat(ex.Price, 64)
			quantity, _ := strconv.ParseFloat(ex.Quantity, 64)
			ohlc.Update(price, quantity, ex.TakerSide)
		}
		// writer.Flush()
		last := res.Last()
		start = last.CreatedAt
		time.Sleep(3000 * time.Millisecond)
	}
}

func GetExecutions(startFrom time.Time) {}
