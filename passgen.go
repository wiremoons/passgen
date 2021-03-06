//
//	passgen - an application to provide password suggestions based on
//	a collection of three letter English words.
//
// Copyright 2016 Simon Rowe <simon@wiremoons.com>.
// All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.
//
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// GLOBAL VARIABLES

// set the version of the app here
var appversion = "0.7.1"
var appname string

// below used by flag to store command line arguments given by the user
var numwords int
var numsuggestions int
var passcase bool
var helpMe bool
var quiet bool
var remove bool
var version bool

// init function always runs before main() so used here to
// set-up the required command line flag variables
func init() {
	// IntVar; StringVar; BoolVar options for flag
	// format required: variable, cmd line flag, initial value, description.
	flag.BoolVar(&passcase, "c", false, "\tUSE: '-c=true' provide mixed case passwords. Note: useful with -q only [DEFAULT: lowercase]")
	flag.BoolVar(&helpMe, "h", false, "\tUSE: '-h' display more detailed help about this program")
	flag.BoolVar(&quiet, "q", false, "\tUSE: '-q=true' to obtain just ONE password - no other screen output [DEFAULT: additional info output]")
	flag.BoolVar(&remove, "r", false, "\tUSE: '-r=true' remove password spaces. Note: useful with -q only [DEFAULT: with spaces]")
	flag.IntVar(&numsuggestions, "s", 3, "\tUSE: '-s #' where # is the number of password suggestions offered [DEFAULT: 3]")
	flag.BoolVar(&version, "v", false, "\tUSE: '-v=true.' display the application version [DEFAULT: false]")
	flag.IntVar(&numwords, "w", 3, "\tUSE: '-w #' where # is the number of three letter words to use [DEFAULT: 3]")
	appname = filepath.Base(os.Args[0])
}

func main() {
	// get the command line args passed to the program
	flag.Parse()

	// was the command line flag '-h' used?
	if helpMe {
		// call function to display information about the application
		printHelp()
		// call to display the standard command line usage info
		flag.Usage()
		// let user know we ran as expected
		fmt.Printf("\n\nAll is well.\n\n")
		// exit the application
		os.Exit(0)
	}

	// was the command line flag '-v' used?
	if version {
		// print app name called and version information
		fmt.Printf("\n Running %s version %s\n", appname, appversion)
		fmt.Printf(" Built with Go Complier '%s' on Golang version '%s'\n", runtime.Compiler, runtime.Version())
		fmt.Printf(" - Author's web site: http://www.wiremoons.com/\n")
		fmt.Printf(" - Source code for %s: https://github.com/wiremoons/passgen/\n", appname)
		fmt.Printf("\nAll is well\n")
		// exit the application
		os.Exit(0)
	}

	// check how many three letter words the user wants to include in
	// their new password?
	// if given a zero or negative value then reset to '3' the default
	if numwords <= 0 {
		numwords = 3
	}
	// check how many password suggestions the user wants to include?
	// if given a zero or negative value then reset to '3' the default
	if numsuggestions <= 0 {
		numsuggestions = 3
	}

	// create a seed from current time
	rand.Seed(time.Now().UTC().UnixNano())

	// quiet mode - so just output ONE password (ie -s 1) at whatever word
	// length for -w and nothing else. Also check for removal of spaces
	// and mixed case preference
	if quiet {
		// variable to hold quite password 'qpassword'
		var qpassword string
		// remove spaces in password if true on command line with -r
		if remove {
			qpassword = strings.Replace(getPassword(numwords), " ", "", -1)
		} else {
			// just get password with spaces
			qpassword = getPassword(numwords)
		}
		// check if mixed case password requested with -c
		if passcase {
			qpassword = mixedPassword(qpassword)
		}
		fmt.Printf("%s\n", qpassword)
		// done - so exit application
		os.Exit(0)
	}

	// default output is to include mixed case passwords and provide a
	// random number as well
	passcase = true
	// OK - so run as normal and display output
	fmt.Printf("\n\t\t\tTHREE WORD - PASSWORD GENERATOR\n\t\t\t¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯\n")
	fmt.Printf("» Number of three letter words available in the pool is: %d\n", (len(passmap)))
	fmt.Printf("» Number of three letter words to include in the suggested password is: %d\n", numwords)
	fmt.Printf("\t» Password character length will therefore be: %d\n", (numwords * 3))
	fmt.Printf("» Mixed case passwords to be provided: %s\n", strconv.FormatBool(passcase))
	fmt.Printf("» Offering %d suggested passwords for your consideration:\n\n", numsuggestions)

	// get password suggestion(s) based on number requested (numsuggestions),
	// and include specified number  of three letter words requested (numword)
	for ; numsuggestions > 0; numsuggestions-- {
		// defaultpass: passwords with spaces included between words
		defaultpass := getPassword(numwords)
		// nospacepass: passwords with NO spaces included between words
		nospacepass := strings.Replace(defaultpass, " ", "", -1)
		// get a mixed case password
		mixedcasepass := mixedPassword(nospacepass)
		//fmt.Printf("\t%s\n", getPassword(numwords))
		fmt.Printf("\t%s    %s    %s    %d\n", defaultpass, nospacepass, mixedcasepass, rand.Intn(100))
	}

	fmt.Printf("\nTo change the password suggestion output shown above, use the command line options.\n")
	fmt.Printf("Run the program as follows for more help:  %s -h\n", appname)
	fmt.Printf("\nAll is well\n")
}

// getPassword is used to return a suggested password
// the generated password will containing the number of three letter words
// requested in the call tn the function as an int 'numwords'.
// A string containing the generated password is return when the function
// completes.
func getPassword(numwords int) string {
	var passSuggestion string
	// get three letter word associated with random number:
	for ; numwords > 0; numwords-- {
		passSuggestion = passSuggestion + " " + (passmap[rand.Intn(len(passmap))])
	}
	// remove leading space from password string
	passSuggestion = strings.TrimLeft(passSuggestion, " ")
	// done - return password suggestion
	return passSuggestion
}

// mixedPassword converts a string passed to the function
// and returns a string converted to include mixed case
func mixedPassword(lcpassword string) string {
	// variable for new mixed case password
	var mcpassword string
	// for each letter in the password string - get a random number
	// if random number is even make letter uppercase
	for _, c := range lcpassword {
		dice := rand.Intn(100)
		//fmt.Printf("random number is: %d\n", dice)
		// if number is even
		if dice%2 == 0 {
			mcpassword = mcpassword + string(unicode.ToUpper(c))
		} else {
			mcpassword = mcpassword + string(c)
		}
	}
	// done - return password suggestion
	return mcpassword
}
