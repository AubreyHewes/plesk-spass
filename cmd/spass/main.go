// CLI spass
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

var (
	version = "dev"
	// Exit codes from <sysexits.h>
	EX_TEMPFAIL = 75
)

func main() {
	//flag.Usage = usage
	var threshold = flag.Int("t", 50, "set discard threshold")
	var showVersion = flag.Bool("v", false, "prints current version")
	var writeToDirectory = flag.String("D", "", "write discarded mails to this directory")
	flag.Parse()

	if *showVersion {
		fmt.Printf("spass version %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
		os.Exit(2)
	}

	threshold64 := float64(*threshold)

	inHeaders := true
	isWritingToDirectory := *writeToDirectory != ""

	if err := syscall.SetNonblock(0, true); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)

	//_, _ = fmt.Fprintf(os.Stderr, "DEBUG: %s\n", "testing spam score");
	buffer := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		// the first empty line is between the headers and the body
		if inHeaders && len(strings.TrimSpace(line)) == 0 {
			inHeaders = false
			if isWritingToDirectory {
				buffer = append(buffer, line)
			}
		}

		if inHeaders && isWritingToDirectory {
			buffer = append(buffer, line)
		}

		if inHeaders && strings.HasPrefix(line, "X-Spam-Status: Yes") {
			//_, _ = fmt.Fprintf(os.Stderr, "DEBUG: %s\n", line);
			score := regexp.MustCompile(`^.*score=(.*?)\s.*$`).ReplaceAllString(line, `$1`)
			scoreFloat64, _ := strconv.ParseFloat(score, 32)
			if scoreFloat64 >= threshold64 {

				//
				// Command  exit  status  codes  are  expected  to  follow the conventions
				// defined in <sysexits.h>.  Exit status 0 means normal successful completion.
				//
				// In the case of a non-zero exit status, a limited amount of command output
				// is reported in an delivery status notification.   When  the  output
				// begins  with  a  4.X.X  or  5.X.X enhanced status code, the status code
				// takes precedence over the non-zero exit status (Postfix version 2.3 and
				// later).
				//
				// Thus, see /usr/include/sysexits.h for appropriate exit status codes.
				// To defer mail, EX_TEMPFAIL (75) comes to mind.

				_, _ = fmt.Fprintf(os.Stderr, "discarded spam %.2f/%.2f\n", scoreFloat64, threshold64)
				os.Exit(EX_TEMPFAIL)
			}
		}

		if !inHeaders && isWritingToDirectory {
			buffer = append(buffer, line)
		}
	}

	if isWritingToDirectory {
		// TODO write to directory
		fmt.Println(strings.Join(buffer, "\n"))
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
	}
}
