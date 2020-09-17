package main

import (
	"fmt"
	"math/rand"
	"github.com/myzhan/boomer"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"

)

const restaurants int = 30
const maxtables int = 50

var tables []string
var baseURL string

func init() {
	tables = GenTableList(restaurants, maxtables)
	baseURL = "http://127.0.0.1:8080"
}

func getRandTable(tableList []string) string {
	return tableList[rand.Intn(len(tableList))]
}

func visit() {
	client :=  resty.New()
	table := getRandTable(tables)
	v := NewFakeVisit(table)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(*v).
		Post(fmt.Sprintf("%s/visit", baseURL))


	if err != nil {
		log.Error(err)
		boomer.RecordFailure("http", "visit", resp.Time().Microseconds(), err.Error())
		return
	}

	if resp.StatusCode() != 201 {
		errStr := fmt.Sprintf("Unexpected POST response code: %d Response: '%s'", resp.StatusCode(), string(resp.Body()))

		log.Error(errStr)
		boomer.RecordFailure("http", "visit", resp.Time().Microseconds(), errStr)
		return
	}

	boomer.RecordSuccess("http", "visit", resp.Time().Microseconds(), resp.Size())
}

func report() {
	// nothing yet
}

func main() {
	task1 := &boomer.Task{
		Name:   "visit",
		Weight: 96,
		Fn:     visit,
	}

	// 4% infected
	task2 := &boomer.Task{
		Name:   "report",
		Weight: 4,
		Fn:     report,
	}

	boomer.Run(task1, task2)
}
