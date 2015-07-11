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
        Updated: 21 Mar 2015 -  [v0.5] increased the three letter word pool from
										573 words to 1,312 words. and fixed the errors identified by
										the 'Go Report Card' here: http://goreportcard.com/report/wiremoons/passgen
        Updated: 20 June 2015 - typo fixes kindly reported by GitHub User @RickCogley.
				Updated: 11 July 2015 - [v0.6] changed default output to show multiple suggestions
										(ie with & without spaces & mixed case) plus as a
										consequence made '-r' only relevant to '-q' option

       TODO - maybe check for newer version and update if needed?
       TODO - add -c mixed case output mode

*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// GLOBAL VARIABLES
//
// set the version of the app here
var appversion = "0.6"
var appname string

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
// Word list acquired from the 'Association of British Scrabble Players' (ABSP), from the web page
// is here: http://www.absp.org.uk/words/3lw.shtml
var passmap = map[int]string{
	// start one digit  block below
	1: "aah", 2: "aal", 3: "aas", 4: "aba", 5: "abb",
	6: "abo", 7: "abs", 8: "aby", 9: "ace", 10: "ach",
	// start two digit  block below
	11: "act", 12: "add", 13: "ado", 14: "ads", 15: "adz",
	16: "aff", 17: "aft", 18: "aga", 19: "age", 20: "ago",
	21: "ags", 22: "aha", 23: "ahi", 24: "ahs", 25: "aia",
	26: "aid", 27: "ail", 28: "aim", 29: "ain", 30: "air",
	31: "ais", 32: "ait", 33: "aka", 34: "ake", 35: "ala",
	36: "alb", 37: "ale", 38: "alf", 39: "all", 40: "alp",
	41: "als", 42: "alt", 43: "alu", 44: "ama", 45: "ame",
	46: "ami", 47: "amp", 48: "amu", 49: "ana", 50: "and",
	51: "ane", 52: "ani", 53: "ann", 54: "ans", 55: "ant",
	56: "any", 57: "ape", 58: "apo", 59: "app", 60: "apt",
	61: "arb", 62: "arc", 63: "ard", 64: "are", 65: "arf",
	66: "ark", 67: "arm", 68: "ars", 69: "art", 70: "ary",
	71: "ash", 72: "ask", 73: "asp", 74: "ass", 75: "ate",
	76: "ats", 77: "att", 78: "aua", 79: "aue", 80: "auf",
	81: "auk", 82: "ava", 83: "ave", 84: "avo", 85: "awa",
	86: "awe", 87: "awk", 88: "awl", 89: "awn", 90: "axe",
	91: "aye", 92: "ays", 93: "ayu", 94: "azo", 95: "baa",
	96: "bac", 97: "bad", 98: "bag", 99: "bah", 100: "bal",
	// start three digit  block below
	101: "bam", 102: "ban", 103: "bap", 104: "bar", 105: "bas",
	106: "bat", 107: "bay", 108: "bed", 109: "bee", 110: "beg",
	111: "bel", 112: "ben", 113: "bes", 114: "bet", 115: "bey",
	116: "bez", 117: "bib", 118: "bid", 119: "big", 120: "bin",
	121: "bio", 122: "bis", 123: "bit", 124: "biz", 125: "boa",
	126: "bob", 127: "bod", 128: "bog", 129: "boh", 130: "boi",
	131: "bok", 132: "bon", 133: "boo", 134: "bop", 135: "bor",
	136: "bos", 137: "bot", 138: "bow", 139: "box", 140: "boy",
	141: "bra", 142: "bro", 143: "brr", 144: "bru", 145: "bub",
	146: "bud", 147: "bug", 148: "bum", 149: "bun", 150: "bur",
	151: "bus", 152: "but", 153: "buy", 154: "bye", 155: "bys",
	156: "caa", 157: "cab", 158: "cad", 159: "cag", 160: "cam",
	161: "can", 162: "cap", 163: "car", 164: "cat", 165: "caw",
	166: "cay", 167: "caz", 168: "cee", 169: "cel", 170: "cep",
	171: "cha", 172: "che", 173: "chi", 174: "cid", 175: "cig",
	176: "cis", 177: "cit", 178: "cly", 179: "cob", 180: "cod",
	181: "cog", 182: "col", 183: "con", 184: "coo", 185: "cop",
	186: "cor", 187: "cos", 188: "cot", 189: "cow", 190: "cox",
	191: "coy", 192: "coz", 193: "cru", 194: "cry", 195: "cub",
	196: "cud", 197: "cue", 198: "cum", 199: "cup", 200: "cur",
	201: "cut", 202: "cuz", 203: "cwm", 204: "dab", 205: "dad",
	206: "dae", 207: "dag", 208: "dah", 209: "dak", 210: "dal",
	211: "dam", 212: "dan", 213: "dap", 214: "das", 215: "daw",
	216: "day", 217: "deb", 218: "dee", 219: "def", 220: "deg",
	221: "dei", 222: "del", 223: "den", 224: "dev", 225: "dew",
	226: "dex", 227: "dey", 228: "dib", 229: "did", 230: "die",
	231: "dif", 232: "dig", 233: "dim", 234: "din", 235: "dip",
	236: "dis", 237: "dit", 238: "div", 239: "dob", 240: "doc",
	241: "dod", 242: "doe", 243: "dof", 244: "dog", 245: "doh",
	246: "dol", 247: "dom", 248: "don", 249: "doo", 250: "dop",
	251: "dor", 252: "dos", 253: "dot", 254: "dow", 255: "doy",
	256: "dry", 257: "dso", 258: "dub", 259: "dud", 260: "due",
	261: "dug", 262: "duh", 263: "dui", 264: "dun", 265: "duo",
	266: "dup", 267: "dux", 268: "dye", 269: "dzo", 270: "ean",
	271: "ear", 272: "eas", 273: "eat", 274: "eau", 275: "ebb",
	276: "ech", 277: "eco", 278: "ecu", 279: "edh", 280: "eds",
	281: "eek", 282: "eel", 283: "een", 284: "eff", 285: "efs",
	286: "eft", 287: "egg", 288: "ego", 289: "ehs", 290: "eik",
	291: "eke", 292: "eld", 293: "elf", 294: "elk", 295: "ell",
	296: "elm", 297: "els", 298: "elt", 299: "eme", 300: "emo",
	301: "ems", 302: "emu", 303: "end", 304: "ene", 305: "eng",
	306: "ens", 307: "eon", 308: "era", 309: "ere", 310: "erf",
	311: "erg", 312: "erk", 313: "erm", 314: "ern", 315: "err",
	316: "ers", 317: "ess", 318: "est", 319: "eta", 320: "eth",
	321: "euk", 322: "eve", 323: "evo", 324: "ewe", 325: "ewk",
	326: "ewt", 327: "exo", 328: "eye", 329: "faa", 330: "fab",
	331: "fad", 332: "fae", 333: "fag", 334: "fah", 335: "fan",
	336: "fap", 337: "far", 338: "fas", 339: "fat", 340: "faw",
	341: "fax", 342: "fay", 343: "fed", 344: "fee", 345: "feg",
	346: "feh", 347: "fem", 348: "fen", 349: "fer", 350: "fes",
	351: "fet", 352: "feu", 353: "few", 354: "fey", 355: "fez",
	356: "fib", 357: "fid", 358: "fie", 359: "fig", 360: "fil",
	361: "fin", 362: "fir", 363: "fit", 364: "fix", 365: "fiz",
	366: "flu", 367: "fly", 368: "fob", 369: "foe", 370: "fog",
	371: "foh", 372: "fon", 373: "fop", 374: "for", 375: "fou",
	376: "fox", 377: "foy", 378: "fra", 379: "fro", 380: "fry",
	381: "fub", 382: "fud", 383: "fug", 384: "fum", 385: "fun",
	386: "fur", 387: "gab", 388: "gad", 389: "gae", 390: "gag",
	391: "gak", 392: "gal", 393: "gam", 394: "gan", 395: "gap",
	396: "gar", 397: "gas", 398: "gat", 399: "gau", 400: "gaw",
	401: "gay", 402: "ged", 403: "gee", 404: "gel", 405: "gem",
	406: "gen", 407: "geo", 408: "ger", 409: "get", 410: "gey",
	411: "ghi", 412: "gib", 413: "gid", 414: "gie", 415: "gif",
	416: "gig", 417: "gin", 418: "gio", 419: "gip", 420: "gis",
	421: "git", 422: "gju", 423: "gnu", 424: "goa", 425: "gob",
	426: "god", 427: "goe", 428: "gon", 429: "goo", 430: "gor",
	431: "gos", 432: "got", 433: "gov", 434: "gox", 435: "goy",
	436: "gub", 437: "gue", 438: "gul", 439: "gum", 440: "gun",
	441: "gup", 442: "gur", 443: "gus", 444: "gut", 445: "guv",
	446: "guy", 447: "gym", 448: "gyp", 449: "had", 450: "hae",
	451: "hag", 452: "hah", 453: "haj", 454: "ham", 455: "han",
	456: "hao", 457: "hap", 458: "has", 459: "hat", 460: "haw",
	461: "hay", 462: "heh", 463: "hem", 464: "hen", 465: "hep",
	466: "her", 467: "hes", 468: "het", 469: "hew", 470: "hex",
	471: "hey", 472: "hic", 473: "hid", 474: "hie", 475: "him",
	476: "hin", 477: "hip", 478: "his", 479: "hit", 480: "hmm",
	481: "hoa", 482: "hob", 483: "hoc", 484: "hod", 485: "hoe",
	486: "hog", 487: "hoh", 488: "hoi", 489: "hom", 490: "hon",
	491: "hoo", 492: "hop", 493: "hos", 494: "hot", 495: "how",
	496: "hox", 497: "hoy", 498: "hub", 499: "hue", 500: "hug",
	501: "huh", 502: "hui", 503: "hum", 504: "hun", 505: "hup",
	506: "hut", 507: "hye", 508: "hyp", 509: "ice", 510: "ich",
	511: "ick", 512: "icy", 513: "ide", 514: "ids", 515: "iff",
	516: "ifs", 517: "igg", 518: "ilk", 519: "ill", 520: "imp",
	521: "ing", 522: "ink", 523: "inn", 524: "ins", 525: "ion",
	526: "ios", 527: "ire", 528: "irk", 529: "ish", 530: "ism",
	531: "iso", 532: "ita", 533: "its", 534: "ivy", 535: "iwi",
	536: "jab", 537: "jag", 538: "jai", 539: "jak", 540: "jam",
	541: "jap", 542: "jar", 543: "jaw", 544: "jay", 545: "jee",
	546: "jet", 547: "jeu", 548: "jew", 549: "jib", 550: "jig",
	551: "jin", 552: "jiz", 553: "job", 554: "joe", 555: "jog",
	556: "jol", 557: "jor", 558: "jot", 559: "jow", 560: "joy",
	561: "jud", 562: "jug", 563: "jun", 564: "jus", 565: "jut",
	566: "kab", 567: "kae", 568: "kaf", 569: "kai", 570: "kak",
	571: "kam", 572: "kas", 573: "kat", 574: "kaw", 575: "kay",
	576: "kea", 577: "keb", 578: "ked", 579: "kef", 580: "keg",
	581: "ken", 582: "kep", 583: "ket", 584: "kex", 585: "key",
	586: "khi", 587: "kid", 588: "kif", 589: "kin", 590: "kip",
	591: "kir", 592: "kis", 593: "kit", 594: "koa", 595: "kob",
	596: "koi", 597: "kon", 598: "kop", 599: "kor", 600: "kos",
	601: "kow", 602: "kue", 603: "kye", 604: "kyu", 605: "lab",
	606: "lac", 607: "lad", 608: "lag", 609: "lah", 610: "lam",
	611: "lap", 612: "lar", 613: "las", 614: "lat", 615: "lav",
	616: "law", 617: "lax", 618: "lay", 619: "lea", 620: "led",
	621: "lee", 622: "leg", 623: "lei", 624: "lek", 625: "lep",
	626: "les", 627: "let", 628: "leu", 629: "lev", 630: "lew",
	631: "lex", 632: "ley", 633: "lez", 634: "lib", 635: "lid",
	636: "lie", 637: "lig", 638: "lin", 639: "lip", 640: "lis",
	641: "lit", 642: "lob", 643: "lod", 644: "log", 645: "loo",
	646: "lop", 647: "lor", 648: "los", 649: "lot", 650: "lou",
	651: "low", 652: "lox", 653: "loy", 654: "lud", 655: "lug",
	656: "lum", 657: "lur", 658: "luv", 659: "lux", 660: "luz",
	661: "lye", 662: "lym", 663: "maa", 664: "mac", 665: "mad",
	666: "mae", 667: "mag", 668: "mak", 669: "mal", 670: "mam",
	671: "man", 672: "map", 673: "mar", 674: "mas", 675: "mat",
	676: "maw", 677: "max", 678: "may", 679: "med", 680: "mee",
	681: "meg", 682: "meh", 683: "mel", 684: "mem", 685: "men",
	686: "mes", 687: "met", 688: "meu", 689: "mew", 690: "mho",
	691: "mib", 692: "mic", 693: "mid", 694: "mig", 695: "mil",
	696: "mim", 697: "mir", 698: "mis", 699: "mix", 700: "miz",
	701: "mna", 702: "moa", 703: "mob", 704: "moc", 705: "mod",
	706: "moe", 707: "mog", 708: "moi", 709: "mol", 710: "mom",
	711: "mon", 712: "wit", 713: "moo", 714: "mop", 715: "mor",
	716: "mos", 717: "mot", 718: "mou", 719: "mow", 720: "moy",
	721: "moz", 722: "mud", 723: "mug", 724: "mum", 725: "mun",
	726: "mus", 727: "mut", 728: "mux", 729: "myc", 730: "nab",
	731: "nae", 732: "nag", 733: "nah", 734: "nam", 735: "nan",
	736: "nap", 737: "nas", 738: "nat", 739: "naw", 740: "nay",
	741: "neb", 742: "ned", 743: "nee", 744: "nef", 745: "neg",
	746: "nek", 747: "nep", 748: "net", 749: "new", 750: "nib",
	751: "nid", 752: "nie", 753: "nil", 754: "nim", 755: "nip",
	756: "nis", 757: "nit", 758: "nix", 759: "nob", 760: "nod",
	761: "nog", 762: "noh", 763: "nom", 764: "non", 765: "noo",
	766: "nor", 767: "nos", 768: "not", 769: "now", 770: "nox",
	771: "noy", 772: "nth", 773: "nub", 774: "nun", 775: "nur",
	776: "nus", 777: "nut", 778: "nye", 779: "nys", 780: "oaf",
	781: "oak", 782: "oar", 783: "oat", 784: "oba", 785: "obe",
	786: "obi", 787: "obo", 788: "obs", 789: "oca", 790: "och",
	791: "oda", 792: "odd", 793: "ode", 794: "ods", 795: "oes",
	796: "off", 797: "oft", 798: "ohm", 799: "oho", 800: "ohs",
	801: "oik", 802: "oil", 803: "ois", 804: "oka", 805: "oke",
	806: "old", 807: "ole", 808: "olm", 809: "oms", 810: "one",
	811: "ono", 812: "ons", 813: "ony", 814: "oof", 815: "ooh",
	816: "oom", 817: "oon", 818: "oop", 819: "oor", 820: "oos",
	821: "oot", 822: "ope", 823: "ops", 824: "opt", 825: "ora",
	826: "orb", 827: "orc", 828: "ord", 829: "ore", 830: "orf",
	831: "ors", 832: "ort", 833: "ose", 834: "oud", 835: "ouk",
	836: "oup", 837: "our", 838: "ous", 839: "out", 840: "ova",
	841: "owe", 842: "owl", 843: "own", 844: "owt", 845: "oxo",
	846: "oxy", 847: "oye", 848: "oys", 849: "pac", 850: "pad",
	851: "pah", 852: "pal", 853: "pam", 854: "pan", 855: "pap",
	856: "par", 857: "pas", 858: "pat", 859: "pav", 860: "paw",
	861: "pax", 862: "pay", 863: "pea", 864: "pec", 865: "ped",
	866: "pee", 867: "peg", 868: "peh", 869: "pel", 870: "pen",
	871: "pep", 872: "per", 873: "pes", 874: "pet", 875: "pew",
	876: "phi", 877: "pho", 878: "pht", 879: "pia", 880: "pic",
	881: "pie", 882: "pig", 883: "pin", 884: "pip", 885: "pir",
	886: "pis", 887: "pit", 888: "piu", 889: "pix", 890: "plu",
	891: "ply", 892: "poa", 893: "pod", 894: "poh", 895: "poi",
	896: "pol", 897: "pom", 898: "poo", 899: "pop", 900: "pos",
	901: "pot", 902: "pow", 903: "pox", 904: "poz", 905: "pre",
	906: "pro", 907: "pry", 908: "psi", 909: "pst", 910: "pub",
	911: "pud", 912: "pug", 913: "puh", 914: "pul", 915: "pun",
	916: "pup", 917: "pur", 918: "pus", 919: "put", 920: "puy",
	921: "pya", 922: "pye", 923: "pyx", 924: "qat", 925: "qis",
	926: "qua", 927: "qin", 928: "rad", 929: "rag", 930: "rah",
	931: "rai", 932: "raj", 933: "ram", 934: "ran", 935: "rap",
	936: "ras", 937: "rat", 938: "rav", 939: "raw", 940: "rax",
	941: "ray", 942: "reb", 943: "rec", 944: "red", 945: "ree",
	946: "ref", 947: "reg", 948: "reh", 949: "rei", 950: "rem",
	951: "ren", 952: "reo", 953: "rep", 954: "res", 955: "ret",
	956: "rev", 957: "rew", 958: "rex", 959: "rez", 960: "rho",
	961: "rhy", 962: "ria", 963: "rib", 964: "rid", 965: "rif",
	966: "rig", 967: "rim", 968: "rin", 969: "rip", 970: "rit",
	971: "riz", 972: "rob", 973: "roc", 974: "rod", 975: "roe",
	976: "rok", 977: "rom", 978: "roo", 979: "rot", 980: "row",
	981: "rub", 982: "ruc", 983: "rud", 984: "rue", 985: "rug",
	986: "rum", 987: "run", 988: "rut", 989: "rya", 990: "rye",
	991: "sab", 992: "sac", 993: "sad", 994: "sae", 995: "sag",
	996: "sai", 997: "sal", 998: "sam", 999: "san", 1000: "sap",
	// start four digit  block below
	1001: "sar", 1002: "sat", 1003: "sau", 1004: "sav", 1005: "saw",
	1006: "sax", 1007: "say", 1008: "saz", 1009: "scd", 1010: "sea",
	1011: "sec", 1012: "sed", 1013: "see", 1014: "seg", 1015: "sei",
	1016: "sel", 1017: "sen", 1018: "ser", 1019: "set", 1020: "sew",
	1021: "sex", 1022: "sey", 1023: "sez", 1024: "sha", 1025: "she",
	1026: "shh", 1027: "shy", 1028: "sib", 1029: "sic", 1030: "sif",
	1031: "sik", 1032: "sim", 1033: "sin", 1034: "sip", 1035: "sir",
	1036: "sis", 1037: "sit", 1038: "six", 1039: "ska", 1040: "ski",
	1041: "sky", 1042: "sly", 1043: "sma", 1044: "sny", 1045: "sob",
	1046: "soc", 1047: "sod", 1048: "sog", 1049: "soh", 1050: "sol",
	1051: "som", 1052: "son", 1053: "sop", 1054: "sos", 1055: "sot",
	1056: "sou", 1057: "sov", 1058: "sow", 1059: "sox", 1060: "soy",
	1061: "soz", 1062: "spa", 1063: "spy", 1064: "sri", 1065: "sty",
	1066: "sub", 1067: "sud", 1068: "sue", 1069: "sug", 1070: "sui",
	1071: "suk", 1072: "sum", 1073: "sun", 1074: "sup", 1075: "suq",
	1076: "sur", 1077: "sus", 1078: "swy", 1079: "sye", 1080: "syn",
	1081: "tab", 1082: "tad", 1083: "tae", 1084: "tag", 1085: "tai",
	1086: "taj", 1087: "tak", 1088: "tam", 1089: "tan", 1090: "tao",
	1091: "tap", 1092: "tar", 1093: "tas", 1094: "tat", 1095: "tau",
	1096: "tav", 1097: "taw", 1098: "tax", 1099: "tay", 1100: "tea",
	1101: "tec", 1102: "ted", 1103: "tee", 1104: "tef", 1105: "teg",
	1106: "tel", 1107: "ten", 1108: "tes", 1109: "tet", 1110: "tew",
	1111: "tex", 1112: "the", 1113: "tho", 1114: "thy", 1115: "tic",
	1116: "tid", 1117: "tie", 1118: "tig", 1119: "tik", 1120: "til",
	1121: "tin", 1122: "tip", 1123: "tis", 1124: "tit", 1125: "tix",
	1126: "toc", 1127: "tod", 1128: "toe", 1129: "tog", 1130: "tom",
	1131: "ton", 1132: "too", 1133: "top", 1134: "tor", 1135: "tot",
	1136: "tow", 1137: "toy", 1138: "try", 1139: "tsk", 1140: "tub",
	1141: "tug", 1142: "tui", 1143: "tum", 1144: "tun", 1145: "tup",
	1146: "tut", 1147: "tux", 1148: "twa", 1149: "two", 1150: "twp",
	1151: "tye", 1152: "tyg", 1153: "udo", 1154: "uds", 1155: "uey",
	1156: "ufo", 1157: "ugh", 1158: "ugs", 1159: "uke", 1160: "ule",
	1161: "ulu", 1162: "umm", 1163: "ump", 1164: "ums", 1165: "umu",
	1166: "uni", 1167: "uns", 1168: "upo", 1169: "ups", 1170: "urb",
	1171: "urd", 1172: "ure", 1173: "urn", 1174: "urp", 1175: "use",
	1176: "uta", 1177: "ute", 1178: "uts", 1179: "utu", 1180: "uva",
	1181: "vac", 1182: "vae", 1183: "vag", 1184: "van", 1185: "var",
	1186: "vas", 1187: "vat", 1188: "vau", 1189: "vav", 1190: "vaw",
	1191: "vee", 1192: "veg", 1193: "vet", 1194: "vex", 1195: "via",
	1196: "vid", 1197: "vie", 1198: "vig", 1199: "vim", 1200: "vin",
	1201: "vis", 1202: "vly", 1203: "voe", 1204: "vol", 1205: "vor",
	1206: "vow", 1207: "vox", 1208: "vug", 1209: "vum", 1210: "wab",
	1211: "wad", 1212: "wae", 1213: "wag", 1214: "wai", 1215: "wan",
	1216: "wap", 1217: "war", 1218: "was", 1219: "wat", 1220: "waw",
	1221: "wax", 1222: "way", 1223: "web", 1224: "wed", 1225: "wee",
	1226: "wem", 1227: "wen", 1228: "wet", 1229: "wex", 1230: "wey",
	1231: "wha", 1232: "who", 1233: "why", 1234: "wig", 1235: "win",
	1236: "wis", 1237: "wit", 1238: "wiz", 1239: "woe", 1240: "wof",
	1241: "wog", 1242: "wok", 1243: "won", 1244: "woo", 1245: "wop",
	1246: "wos", 1247: "wot", 1248: "wow", 1249: "wox", 1250: "wry",
	1251: "wud", 1252: "wus", 1253: "wye", 1254: "wyn", 1255: "xis",
	1256: "yad", 1257: "yae", 1258: "yag", 1259: "yah", 1260: "yak",
	1261: "yam", 1262: "yap", 1263: "yar", 1264: "yaw", 1265: "yay",
	1266: "yea", 1267: "yeh", 1268: "yen", 1269: "yep", 1270: "yes",
	1271: "yet", 1272: "yew", 1273: "yex", 1274: "ygo", 1275: "yid",
	1276: "yin", 1277: "yip", 1278: "yob", 1279: "yod", 1280: "yok",
	1281: "yom", 1282: "yon", 1283: "you", 1284: "yow", 1285: "yug",
	1286: "yuk", 1287: "yum", 1288: "yup", 1289: "yus", 1290: "zag",
	1291: "zap", 1292: "zas", 1293: "zax", 1294: "zea", 1295: "zed",
	1296: "zee", 1297: "zek", 1298: "zel", 1299: "zep", 1300: "zex",
	1301: "zho", 1302: "zig", 1303: "zin", 1304: "zip", 1305: "zit",
	1306: "ziz", 1307: "zoa", 1308: "zol", 1309: "zoo", 1310: "zos",
	1311: "zuz", 1312: "zzz"}

// init() function - always runs before main() - used here to set-up required command line flag variables
//
func init() {
	// IntVar; StringVar; BoolVar all required: variable, cmd line flag, initial value, description used by flag.Usage() on error / help
	flag.BoolVar(&passcase, "c", false, "\tUSE: '-c=true' to get mixed case passwords. Note: useful with -q only [DEFAULT: lowercase]")
	flag.BoolVar(&helpMe, "h", false, "\tUSE: '-h' to provide more detailed help about this program")
	flag.BoolVar(&quiet, "q", false, "\tUSE: '-q=true' to obtain just ONE password - no other screen output [DEFAULT: additional info output]")
	flag.BoolVar(&remove, "r", false, "\tUSE: '-r=true' to remove spaces. Note: useful with -q only [DEFAULT: with spaces]")
	flag.IntVar(&numsuggestions, "s", 3, "\tUSE: '-s no.' where no. is the number of password suggestions offered [DEFAULT: 3]")
	flag.BoolVar(&version, "v", false, "\tUSE: '-v=true.' display the application version [DEFAULT: false]")
	flag.IntVar(&numwords, "w", 3, "\tUSE: '-w no.' where no. is the number of three letter words to use [DEFAULT: 3]")
	appname = filepath.Base(os.Args[0])
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
		fmt.Printf("\n\nAll is well.\n\n")
		// exit the application
		os.Exit(0)
	}

	// check if the user just wanted to know the version using the command line flag '-v'
	if version {
		// print app name called and version information
		fmt.Printf("\n Running %s version %s\n", appname, appversion)
		fmt.Printf(" - Author's web site: http://www.wiremoons.com/\n")
		fmt.Printf(" - Source code for %s: https://github.com/wiremoons/passgen/\n", appname)
		fmt.Printf("\nAll is well\n")
		// exit the application
		os.Exit(0)
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

	// quiet mode - so just output ONE password (ie -s 1) at whatever word
	// length for -w and nothing else. Also check for remove space and mixed case
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
			qpassword = getMixPassword(qpassword)
		}
		fmt.Printf("%s\n", qpassword)
		// done - so exit application
		os.Exit(0)
	}

	// default output is to include mixed case passwords as well
	passcase = true
	// OK - so run as normal and display output
	fmt.Printf("\n\t\t\tTHREE WORD - PASSWORD GENERATOR\n\t\t\t¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯\n")
	fmt.Printf("• Number of three letter words available in the pool is: %d\n", (len(passmap)))
	fmt.Printf("• Number of three letter words to include in the suggested password is: %d\n", numwords)
	fmt.Printf("\t• Password character length will therefore be: %d\n", (numwords * 3))
	fmt.Printf("• Mixed case passwords to be provided: %s\n", strconv.FormatBool(passcase))
	fmt.Printf("• Offering %d suggested passwords for your consideration:\n\n", numsuggestions)
	//}
	// get password suggestion(s) based on number requested (numsuggestions), and include specified number
	// of three letter words requested (numword)
	for ; numsuggestions > 0; numsuggestions-- {
		// defaultpass: passwords with spaces included between words
		defaultpass := getPassword(numwords)
		// nospacepass: passwords with NO spaces included between words
		nospacepass := strings.Replace(defaultpass, " ", "", -1)
		// get a mixed case password
		mixedcasepass := getMixPassword(nospacepass)
		//fmt.Printf("\t%s\n", getPassword(numwords))
		fmt.Printf("\t| %s |   | %s |   | %s |\n", defaultpass, nospacepass, mixedcasepass)
	}
	// END OF MAIN()
	fmt.Printf("\nTo change the password suggestion output shown above, use the command line options.\n")
	fmt.Printf("Run the program as follows for more help:  %s -h\n", appname)
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
	// done - return password suggestion
	return passSuggestion
}

// Function to return a string converted to mixed case
func getMixPassword(lcpassword string) string {
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

// Function to print out some basic help information for the user
func printHelp() {
	helptext := `
	THREE WORD - PASSWORD GENERATOR
	¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯

	About
	¯¯¯¯¯
	This application will generate password suggestions based on a pool of
	over 1,000 three letter English words. The words are selected from
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
	¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯
	While the passwords generated look far too simple and easy to be secure, they
	are in fact very secure, and difficult to crack. Just because they look simple
	to a human - it doesn't mean they are simple to work out using a computer. They
	are in fact quite hard to work out for a computer. The reason for this is that
	they are randomly generated, not a single dictionary word, or a single common
	name. This makes the password harder to 'find' as it is not commonly known.
	It is a common misconception that a password has to be 'complex' to be any good.

	Unfortunately we have been led to believe that the more complex a password is -
	the better and more secure it will be - which is in fact wrong.

	In fact a longer password, that can more easily be remember, and therefore changed
	more frequently as a consequence, actually offers a far greater degree of security.

	For more information and explanations of this, please see the web pages
	included below under 'References'. There are plenty of expert sources on
	the Internet also, that will explain the benefits and security of using a randomly
	generated three word (or more) combination password. Just remember - your password
	must be at least nine characters in total - or longer if possible. You can of course
	always add additional punctuation, should you wish!

	So How Many Possible Passwords Are There?
	¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯
	There are over 1,000 three letter words in the pool that can be chosen from, and
	assuming you use three of these words combined, that provide 1,000^3 (1,000 to power
	of 3) possible combinations - of which one is your password.

	So - 1,000 x 1,000 x 1,000 = 1,000,000,000 (one billion) possibilities.

	If you use the mixed case option (upper and lower case) - then number increases further
	of course - and you can still add numbers, and/or punctuation characters if you wish too.

	Or just increase you password length to 12 characters, so use four of the three letter
	words, and you end up with 1,000,000,000,000 (one thousand billion) possibilities -
	and that is just lower case letters only.


	References
	¯¯¯¯¯¯¯¯¯¯
	Thomas Baekdal - The Usability of Passwords - FAQ
	 - http://www.baekdal.com/insights/the-usability-of-passwords-faq
	Steve Gibson - GRC 'How Big is Your Haystack?'
	 - https://www.grc.com/haystack.htm
	Application 'passgen' - authors web site
	 - http://www.wiremoons.com/

	`
	// now output the above to screen
	fmt.Println(helptext)
}
