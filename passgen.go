/*
	passgen - an application to provide password suggestions based on a collection of
	three letter English words.

	- by default, three passwords are suggested each time the application is run;
	- the number of three letter words used to create a password can be changed, from default of three;
	- optional upper and lower-case combinations can be offered with a command line switch.

	Created: 20 Nov 2014 - initial program written
        Updated: 09 Dec 2014 - added quiet mode output added, and published to GitHub
        Updated: 12 Dec 2014 - added space removal option -r, fixed leading space GitHub Issue #1 
                                 and -v option for version output

       TODO - maybe check for newer version and update if needed?
       TODO - add -c mixed case output mode
	
*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"strings"
)

// GLOBAL VARIABLES
//
// set the version of the app here

var appversion string = "0.4"

// below used by flag for command line args
var numwords int
var numsuggestions int
var passcase bool
var helpMe bool
var quiet bool
var remove bool
var version bool

// var passmap holds a map of strings each containing a three letter word, each with a numeric key
// if new words are required - just add them to the end of the list
var passmap = map[int]string{1: "axe", 2: "azo", 3: "baa", 4: "bad", 5: "bag", 6: "bah", 7: "bam", 8: "ban", 9: "bar",
	10: "bat", 11: "bay", 12: "bed", 13: "bee", 14: "beg", 15: "bet", 16: "bey", 17: "bib", 18: "bid", 19: "big",
	20: "bin", 21: "bio", 22: "bit", 23: "boa", 24: "bob", 25: "bod", 26: "bog", 27: "boo", 28: "bop", 29: "bot",
	30: "bow", 31: "box", 32: "boy", 33: "bra", 34: "bro", 35: "bub", 36: "bud", 37: "bug", 38: "bum", 39: "bun",
	40: "bus", 41: "but", 42: "buy", 43: "bye", 44: "cab", 45: "cad", 46: "cam", 47: "can", 48: "cap", 49: "car",
	50: "cat", 51: "caw", 52: "cee", 53: "cha", 54: "chi", 55: "cob", 56: "cod", 57: "cog", 58: "con", 59: "coo",
	60: "cop", 61: "cot", 62: "cow", 63: "cox", 64: "coy", 65: "cry", 66: "cub", 67: "cud", 68: "cue", 69: "cup",
	70: "cur", 71: "cut", 72: "dab", 73: "dad", 74: "dag", 75: "dam", 76: "day", 77: "dee", 78: "den", 79: "dew",
	80: "dib", 81: "did", 82: "die", 83: "dig", 84: "dim", 85: "din", 86: "dip", 87: "doe", 88: "dog", 89: "don",
	90: "doo", 91: "dop", 92: "dot", 93: "dry", 94: "dub", 95: "dud", 96: "due", 97: "dug", 98: "duh", 99: "dun",
	100: "duo", 101: "dux", 102: "dye", 103: "ear", 104: "eat", 105: "ebb", 106: "eel", 107: "egg", 108: "ego", 109: "eke",
	110: "elf", 111: "elk", 112: "elm", 113: "emo", 114: "emu", 115: "end", 116: "eon", 117: "era", 118: "erg", 119: "err",
	120: "eve", 121: "ewe", 122: "eye", 123: "fab", 124: "fad", 125: "fag", 126: "fan", 127: "far", 128: "far", 129: "fat",
	130: "fax", 131: "fay", 132: "fed", 133: "fee", 134: "fen", 135: "few", 136: "fey", 137: "fez", 138: "fib", 139: "fie",
	140: "fig", 141: "fin", 142: "fir", 143: "fit", 144: "fix", 145: "fly", 146: "fob", 147: "foe", 148: "fog", 149: "fon",
	150: "fop", 151: "for", 152: "fox", 153: "fry", 154: "fun", 155: "fur", 156: "gab", 157: "gag", 158: "gak", 159: "gal",
	160: "gap", 161: "gas", 162: "gaw", 163: "gay", 164: "gee", 165: "gel", 166: "gem", 167: "get", 168: "gig", 169: "gil",
	170: "gin", 171: "git", 172: "gnu", 173: "gob", 174: "god", 175: "goo", 176: "got", 177: "gum", 178: "gun", 179: "gut",
	180: "guy", 181: "gym", 182: "had", 183: "hag", 184: "hal", 185: "ham", 186: "has", 187: "hat", 188: "hay", 189: "hem",
	190: "hen", 191: "her", 192: "hew", 193: "hex", 194: "hey", 195: "hid", 196: "him", 197: "hip", 198: "his", 199: "hit",
	200: "hoe", 201: "hog", 202: "hop", 203: "hot", 204: "how", 205: "hoy", 206: "hub", 207: "hue", 208: "hug", 209: "hug",
	210: "huh", 211: "hum", 212: "hut", 213: "ice", 214: "ick", 215: "icy", 216: "ilk", 217: "ill", 218: "imp", 219: "ink",
	220: "inn", 221: "ion", 222: "ire", 223: "irk", 224: "ism", 225: "its", 226: "jab", 227: "jag", 228: "jah", 229: "jak",
	230: "jam", 231: "jap", 232: "jar", 233: "jaw", 234: "jay", 235: "jem", 236: "jet", 237: "Jew", 238: "jib", 239: "jig",
	240: "job", 241: "joe", 242: "jog", 243: "jon", 244: "jot", 245: "joy", 246: "jug", 247: "jus", 248: "jut", 249: "keg",
	250: "key", 251: "kid", 252: "kin", 253: "kit", 254: "koa", 255: "kob", 256: "koi", 257: "lab", 258: "lad", 259: "lag",
	260: "lap", 261: "law", 262: "lax", 263: "lay", 264: "lea", 265: "led", 266: "leg", 267: "lei", 268: "let", 269: "lew",
	270: "lid", 271: "lie", 272: "lip", 273: "lit", 274: "lob", 275: "log", 276: "loo", 277: "lop", 278: "lot", 279: "low",
	280: "lug", 281: "lux", 282: "lye", 283: "mac", 284: "mad", 285: "mag", 286: "man", 287: "map", 288: "mar", 289: "mat",
	290: "maw", 291: "max", 292: "may", 293: "men", 294: "met", 295: "mic", 296: "mid", 297: "mit", 298: "mix", 299: "mob",
	300: "mod", 301: "mog", 302: "mom", 303: "mon", 304: "moo", 305: "mop", 306: "mow", 307: "mud", 308: "mug", 309: "mum",
	310: "nab", 311: "nag", 312: "nap", 313: "nee", 314: "neo", 315: "net", 316: "new", 317: "nib", 318: "nil", 319: "nip",
	320: "nit", 321: "nix", 322: "nob", 323: "nod", 324: "nog", 325: "nor", 326: "not", 327: "now", 328: "nub", 329: "nun",
	330: "nut", 331: "oaf", 332: "oak", 333: "oar", 334: "oat", 335: "odd", 336: "ode", 337: "off", 338: "oft", 339: "ohm",
	340: "oil", 341: "old", 342: "ole", 343: "one", 344: "opt", 345: "orb", 346: "ore", 347: "our", 348: "out", 349: "out",
	350: "ova", 351: "owe", 352: "owl", 353: "own", 354: "pac", 355: "pad", 356: "pal", 357: "pan", 358: "pap", 359: "par",
	360: "pat", 361: "paw", 362: "pax", 363: "pay", 364: "pea", 365: "pee", 366: "peg", 367: "pen", 368: "pep", 369: "per",
	370: "pet", 371: "pew", 372: "pic", 373: "pie", 374: "pig", 375: "pin", 376: "pip", 377: "pit", 378: "pix", 379: "ply",
	380: "pod", 381: "pog", 382: "poi", 383: "poo", 384: "pop", 385: "pot", 386: "pow", 387: "pox", 388: "pro", 389: "pry",
	390: "pub", 391: "pud", 392: "pug", 393: "pun", 394: "pup", 395: "pus", 396: "put", 397: "pyx", 398: "qat", 399: "qua",
	400: "quo", 401: "rad", 402: "rag", 403: "ram", 404: "ran", 405: "rap", 406: "rat", 407: "raw", 408: "ray", 409: "red",
	410: "rib", 411: "rid", 412: "rig", 413: "rim", 414: "rip", 415: "rob", 416: "roc", 417: "rod", 418: "roe", 419: "rot",
	420: "row", 421: "rub", 422: "rue", 423: "rug", 424: "rum", 425: "run", 426: "rut", 427: "rye", 428: "sac", 429: "sad",
	430: "sag", 431: "sap", 432: "sat", 433: "saw", 434: "sax", 435: "say", 436: "sea", 437: "sec", 438: "see", 439: "set",
	440: "sew", 441: "sex", 442: "she", 443: "shy", 444: "sic", 445: "sim", 446: "sin", 447: "sip", 448: "sir", 449: "sis",
	450: "sit", 451: "six", 452: "ski", 453: "sky", 454: "sly", 455: "sob", 456: "sod", 457: "som", 458: "son", 459: "sop",
	460: "sot", 461: "sow", 462: "soy", 463: "spa", 464: "spy", 465: "sty", 466: "sub", 467: "sue", 468: "sum", 469: "sun",
	470: "sun", 471: "sup", 472: "tab", 473: "tad", 474: "tag", 475: "tam", 476: "tan", 477: "tap", 478: "tar", 479: "tat",
	480: "tax", 481: "tea", 482: "tee", 483: "ten", 484: "the", 485: "tic", 486: "tie", 487: "til", 488: "tin", 489: "tip",
	490: "tit", 491: "toe", 492: "toe", 493: "tom", 494: "ton", 495: "too", 496: "top", 497: "tot", 498: "tow", 499: "toy",
	500: "try", 501: "tub", 502: "tug", 503: "tui", 504: "tut", 505: "two", 506: "ugh", 507: "uke", 508: "ump", 509: "urn",
	510: "use", 511: "van", 512: "vat", 513: "vee", 514: "vet", 515: "vex", 516: "via", 517: "vie", 518: "vig", 519: "vim",
	520: "voe", 521: "vow", 522: "wad", 523: "wag", 524: "wan", 525: "war", 526: "was", 527: "wax", 528: "way", 529: "web",
	530: "wed", 531: "wee", 532: "wen", 533: "wet", 534: "who", 535: "why", 536: "wig", 537: "win", 538: "wit", 539: "wiz",
	540: "woe", 541: "wog", 542: "wok", 543: "won", 544: "woo", 545: "wow", 546: "wry", 547: "wye", 548: "yak", 549: "yam",
	550: "yap", 551: "yaw", 552: "yay", 553: "yea", 554: "yen", 555: "yep", 556: "yes", 557: "yet", 558: "yew", 559: "yip",
	560: "you", 561: "yow", 562: "yum", 563: "yup", 564: "zag", 565: "zap", 566: "zed", 567: "zee", 568: "zen", 569: "zig",
	570: "zip", 571: "zit", 572: "zoa", 573: "zoo"}

// init() function - always runs before main() - used here to set-up required command line flag variables
//
func init() {
	// IntVar; StringVar; BoolVar all required: variable, cmd line flag, initial value, description used by flag.Usage() on error / help

	flag.BoolVar(&passcase, "c", false, "\tUSE: '-c=true' to get mixed case passwords [DEFAULT: lowercase]")
	flag.BoolVar(&helpMe, "h", false, "\tUSE: '-h' to provide more detailed help about this program")
	flag.BoolVar(&quiet, "q", false, "\tUSE: '-q=true' to obtain just ONE password - no other screen output [DEFAULT: additional info output]")
	flag.BoolVar(&remove, "r", false, "\tUSE: '-r=true' to remove spaces in suggested passwords [DEFAULT: with spaces]")
	flag.IntVar(&numsuggestions, "s", 3, "\tUSE: '-s no.' where no. is the number of password suggestions offered [DEFAULT: 3]")
	flag.BoolVar(&version, "v", false, "\tUSE: '-v=true.' display the application version [DEFAULT: false]")
	flag.IntVar(&numwords, "w", 3, "\tUSE: '-w no.' where no. is the number of three letter words to use [DEFAULT: 3]")
}

// PROGRAM MAIN
//
func main() {
	// get the command line args passed to the program
	flag.Parse()
	//
	// see what command line flags we got from user:
	//
	// check if the user just wanted to know more information by using the command line flag '-h'
	if helpMe {
		// call function to display information about the application
		printHelp()
		// call to display the standard command lines usage info
		flag.Usage()
		// let user know we ran as expected
		fmt.Println("\n\nAll is well.\n")
		// exit the application
		os.Exit(-3)
	}

	// check if the user just wanted to know the version using the command line flag '-v'
	if version {
		// print app name called and version information
		fmt.Printf("%s version %s\n",os.Args[0], appversion)
		// exit the application
		os.Exit(-4)
	}



	// check how many three letter words the user wants to include in there password?
	// if given a zero or negative value - reset to '3' the default
	if numwords <= 0 {
		numwords = 3
	}
	// check how many password suggestions the user wants to include?
	// if given a zero or negative value - reset to '3' the default
	if numsuggestions <= 0 {
		numsuggestions = 3
	}

	// create a seed from current time
	rand.Seed(time.Now().UTC().UnixNano())

	// did the user request '-q' for quiet mode?
	if quiet {
		// quiet mode - so just output ONE password (ie -s 1) at whatever length for -w and nothing else
		fmt.Printf("%s\n", getPassword(numwords))
		// done - so exit application
		os.Exit(0)
	}


	// OK - so run as normal and display output
	fmt.Printf("\n\t\t\tTHREE WORD - PASSWORD GENERATOR\n\t\t\t¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯\n")
	fmt.Printf("• Number of three letter words available in the pool is: %d\n", (len(passmap)))
	fmt.Printf("• Number of three letter words to include in the suggested password is: %d\n", numwords)
	fmt.Printf("\t• Password character length will therefore be: %d\n", (numwords * 3))
	fmt.Printf("• Mixed case passwords to be provided: %s\n", strconv.FormatBool(passcase))
	fmt.Printf("• Offering %d suggested passwords for your consideration:\n\n", numsuggestions)
	//}
	// get password suggestion(s) based on number requested (numsuggestions), and to include specified number of three letter
	// word (numword)
	for ; numsuggestions > 0; numsuggestions-- {
		fmt.Printf("\t%s\n", getPassword(numwords))
	}
	// END OF MAIN()
	fmt.Printf("\nAll is well\n")
}

// Function to return a suggested password - containing number of three letter words requested
func getPassword(numwords int) string {
	var passSuggestion string
	// get three letter word associated with random number:
	for ; numwords > 0; numwords-- {
		passSuggestion = passSuggestion + " " + (passmap[rand.Intn(len(passmap))])

	}
	// remove leading space from password string
	passSuggestion = strings.TrimLeft(passSuggestion, " ")
	// if remove spaces is true on command line with -r
	if remove {
		passSuggestion = strings.Replace(passSuggestion, " ", "", -1)
	}
	return passSuggestion
}

// Function to print out some basic help information for the user
func printHelp() {
	helptext := `
	THREE WORD - PASSWORD GENERATOR
	¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯

	About
	¯¯¯¯¯
	This application will generate password suggestions based on a pool of
	several hundred three letter English words. The words are selected from 
	the pool randomly, and then displayed to the screen so you can choose one
	for use as a very secure password.

	It is important that you combine the three letter words together to form
	a single string of characters (without the spaces) - to obtain a password
	with a minimum length on 9 characters. Longer combinations are stronger, but
	unfortunately not all sites accept really long passwords still.

	You can of course add digits/numbers to your password also, and punctuation
	characters too if you wish - but it would be wiser to keep the password
	simple, and easy to remember, but change it more frequently instead, using a
	fresh newly generated one every few weeks.

	Are These Passwords Secure?
	¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯
	While the passwords generated look far too simple and easy to be secure, they 
	are in fact very secure, and difficult to crack. Just because they look simple 
	to a human - it doesn't mean they are simple to work out in a computer. They are 
	in fact quite hard to work out for a computer, as they are random, not a single 
	dictionary word, or a single common name, with perhaps number substitutions and 
	other common 'complex' combinations - or some combination of these with things. 
	It is a common misconception that a password has to be 'complex' to be any good.

	Unfortunately we have been led to believe that the more complex a password is - 
	the better and more secure it will be - which is in fact wrong.

	In fact a longer password, that can more easily be remember and therefore changed
	more frequently as a consequence, actually offers a far greater degree of security.

	For more information and explanations of this, please see the web pages 
	included below under 'References'. There are plenty of other expert sources on 
	the Internet also, that will explain the benefits and security of using a randomly 
	generated three word (or more) combination password. Just remember - your password
	must be at least nine characters in total - or longer if possible. You can of course 
	always add additional punctuation, should you wish!

	So How Many Possible Passwords Are There?
	¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯
	There are over 500 three letter words in the pool that can be chosen from, and 
	assuming you use a three of these words combined, that provide 500^3 (500 to power 
	of 3) possible combinations - of which one is your password.  

	So - 500 x 500 x 500 = 125,000,000 (one hundred and twenty five million) possibilities.

	Maybe that doesn't sound like a lot - but if you could check 20 of them every second, 
	24 hours a day, you would need roughly 60 days to get through them all!

	If you use the mixed case option (upper and lower case) - then number increases further 
	of course - and you can still add numbers, and/or punctuation characters if you wish too.

	Or just increase you password length to 12 characters, so use four of the three letter 
	words, and you end up with 62,500,000,000 (sixty two billion five hundred million) 
	possibilities - and that just lower case letters only. 


	References
	¯¯¯¯¯¯¯¯¯¯
	Thomas Baekdal - The Usability of Passwords - FAQ
	 - http://www.baekdal.com/insights/the-usability-of-passwords-faq
	Steve Gibson - GRC 'How Big is Your Haystack?'
	 - https://www.grc.com/haystack.htm

	`
	// now output the above to screen
	fmt.Println(helptext)
}
