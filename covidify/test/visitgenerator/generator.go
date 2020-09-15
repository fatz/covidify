package visitgenerator

import "github.com/myzhan/boomer"

const restaurants int = 30
const maxtables int = 50

var tables []string

func init() {
	tables = GenTableList(restaurants, maxtables)
}

func visit() {

}

func report() {

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
