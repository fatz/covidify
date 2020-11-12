package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/namsral/flag"
	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const restaurants int = 30
const maxtables int = 50

var tables []string
var baseURL string

func init() {
	tables = GenTableList(restaurants, maxtables)
}

func getRandTable(tableList []string) string {
	return tableList[rand.Intn(len(tableList))]
}

type RandomVisitGen struct {
	requestCounter *prometheus.CounterVec

	Headers map[string]string
	BaseURL string
	Delay   int

	Name string
}

func NewRandomVisitGen(name, baseURL string, requestCounter *prometheus.CounterVec) *RandomVisitGen {
	r := new(RandomVisitGen)

	r.requestCounter = requestCounter

	r.Name = name
	r.Delay = 0
	r.BaseURL = baseURL

	return r
}

func (r *RandomVisitGen) success() {
	log.Infof("[%s] Sucessfull", r.Name)
	r.requestCounter.WithLabelValues("success", r.Name, "visit").Inc()
}

func (r *RandomVisitGen) failure(errStr string) {
	log.Warnf("[%s] Failure - %s", r.Name, errStr)
	r.requestCounter.WithLabelValues("failure", r.Name, "visit").Inc()
}

func (r *RandomVisitGen) visit() {
	client := resty.New()
	table := getRandTable(tables)
	v := NewFakeVisit(table)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeaders(r.Headers).
		SetBody(*v).
		Post(fmt.Sprintf("%s/visit", r.BaseURL))

	if err != nil {
		// log.Error(err)
		// boomer.RecordFailure("http", "visit", resp.Time().Microseconds(), err.Error())
		r.failure(err.Error())
		return
	}

	if resp.StatusCode() != 201 {
		errStr := fmt.Sprintf("Unexpected POST response code: %d Response: '%s'", resp.StatusCode(), string(resp.Body()))

		// log.Error(errStr)
		// boomer.RecordFailure("http", "visit", resp.Time().Microseconds(), errStr)
		r.failure(errStr)
		return
	}

	// boomer.RecordSuccess("http", "visit", resp.Time().Microseconds(), resp.Size())
	r.success()
}

func (r *RandomVisitGen) Run() {
	for {
		r.visit()
		time.Sleep(time.Duration(r.Delay) * time.Millisecond)
	}
}

func main() {
	instances := 4
	delay := 100
	url := "http://localhost:8080"
	hostHeader := ""
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "GENERATOR", 0)

	fs.StringVar(&url, "url", "http://localhost:8080", "url to be used")
	fs.StringVar(&hostHeader, "hostheader", "", "Change the Host header")
	fs.IntVar(&delay, "delay", 100, "Delay in ms")
	fs.IntVar(&instances, "instances", 4, "Load generator instances to start.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fs.PrintDefaults()
		os.Exit(1)
	}

	fs.Parse(os.Args[1:])

	headers := map[string]string{}

	if hostHeader != "" {
		headers["Host"] = hostHeader
	}

	gens := make([]*RandomVisitGen, instances)

	requestCounter := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "covidify_generator_request",
		Help: "The total number requests",
	},
		[]string{"state", "thread", "handler"})

	for i := 0; i < instances; i++ {
		name := fmt.Sprintf("proc%0d", i)
		gens[i] = NewRandomVisitGen(name, url, requestCounter)
		gens[i].Delay = delay
		gens[i].Headers = headers
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	if len(gens) > 1 {
		for i, g := range gens[1:] {
			log.Infof("Starting thread %d", i+1)
			go g.Run()
		}
	}

	log.Info("Starting default thread")
	gens[0].Run()
}
