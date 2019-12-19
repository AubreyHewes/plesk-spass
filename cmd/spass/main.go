// CLI spass
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var (
	version = "dev"
)

func main() {
	//flag.Usage = usage
	var threshold = flag.Int("t", 15, "help message for threshold")
	var showVersion = flag.Bool("v", false, "prints current roxy version")
	flag.Parse()

	if *showVersion {
		fmt.Printf("spass version %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
		os.Exit(2)
	}

	threshold64 := float64(*threshold)

	inHeaders := true
	headers := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	//_, _ = fmt.Fprintf(os.Stderr, "DEBUG: %s\n", "testing spam score");
	for scanner.Scan() {
		line := scanner.Text()

		if inHeaders && len(strings.TrimSpace(line)) == 0 {
			inHeaders = false
			headers = append(headers, "X-Spam-Sieve: 1")
			fmt.Println(strings.Join(headers, "\n"))
		}

		if inHeaders {
			headers = append(headers, line)
		}

		if inHeaders && strings.HasPrefix(line, "X-Spam-Status: Yes") {
			//_, _ = fmt.Fprintf(os.Stderr, "DEBUG: %s\n", line);
			score := regexp.MustCompile(`^.*score=(.*?)\s.*$`).ReplaceAllString(line, `$1`)
			scoreFloat64, _ := strconv.ParseFloat(score, 32)
			if scoreFloat64 >= threshold64 {
				_, _ = fmt.Fprintf(os.Stderr, "5.7.1 deleted spam %.2f/%.2f\n", scoreFloat64, threshold64)
				os.Exit(1)
			}
		}

		if !inHeaders {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
