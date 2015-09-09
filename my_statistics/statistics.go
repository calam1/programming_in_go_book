package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head><title>Statistics</title>
	<body><h3>Statistics</h3>
	<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
	<label for="numbers">Numbers (comma or space-separated):</label><br/>
	<input type="text" name="numbers" size="30"><br/>
	<input type="submit" value="Calculate">
	</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
	stddev  float64
	modes   []float64
}

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers) // create and initialize the struct
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

// this function just breaks down the comma separated numbers passed in the form into a slice/array
func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1) // form is passing value over as one string, so we will replace the commas with spaces so that was can use strings.Fields to split up by spaces and iterate the numeric values
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x) // add the converted float64 into the slice
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false
	}
	return numbers, "", true
}

// Creates a table of data
func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
	<tr><th colspan="2">Results</th></tr>
	<tr><td>Numbers</td><td>%v</td></tr>
	<tr><td>Count</td><td>%d</td></tr>
	<tr><td>Mean</td><td>%f</td></tr>
	<tr><td>Median</td><td>%f</td></tr>
	<tr><td>Standard Deviation</td><td>%f</td></tr>
	<tr><td>Mode</td><td>%v</td></tr>
	</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.stddev, stats.modes)
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	stats.stddev = stdDev(numbers, stats.mean)
	stats.modes = mode(numbers)

	return stats
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle] // if odd list of numbers return the middle number
	if len(numbers)%2 == 0 {  //if even list of numbers
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func stdDev(numbers []float64, mean float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += math.Pow(number-mean, 2)
	}
	variance := total / float64(len(numbers)-1)
	return math.Sqrt(variance)
}

func mode(numbers []float64) []float64 {
	m := make(map[float64]int)
	modes := make([]float64, 0)

	for _, number := range numbers {
		if value, ok := m[number]; ok {
			m[number] = value + 1
		} else {
			m[number] = 1
		}
	}

	for key, value := range m {
		if value > 1 {
			modes = append(modes, key)
		}
	}
	//fmt.Printf("%v", modes)

	return modes
}
