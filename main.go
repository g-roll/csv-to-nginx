package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	if !hasInputData() {
		fmt.Fprintln(os.Stderr, "No input data found on stdin")
		os.Exit(1)
	}

	reader := csv.NewReader(os.Stdin)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", "error reading CSV records", err)
		os.Exit(1)
	}

	for _, record := range records[1:] { // skip header row
		if len(record) != 2 {
			continue
		}

		from := record[0]
		to := record[1]

		nginxBlock := generateNginxBlock(from, to)
		fmt.Println(nginxBlock)
	}
}

func generateNginxBlock(from, to string) string {
	return fmt.Sprintf(`
location %s {
    rewrite ^%s$ %s permanent;
}
`, from, from, to)
}

func hasInputData() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
