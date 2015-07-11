[![NewBSD License](http://img.shields.io/badge/license-NewBSD-orange.svg?style=flat-square)](http://opensource.org/licenses/BSD-3-Clause)

## THREE WORD - PASSWORD GENERATOR
#### (Go language version)

## Application Summary

A simple command line application written to generate a random password created
from a pool of English three letter words.

### About

This application will generate password suggestions based on a pool of
over 1,000 three letter English words. The words are selected
from the pool randomly, and then displayed to the screen so you can
choose one for use as a very secure password.

It is important that you combine the three letter words together to
form a single string of characters (without the spaces)&mdash;to
obtain a password with a minimum length on 9 characters. Longer
combinations are stronger, but unfortunately not all sites or computer
systems accept really long passwords still.

You can of course add digits/numbers to your password also, and
punctuation characters too if you wish&mdash;but it would be wiser to
keep the password simple, and easy to remember, but change it more
frequently instead, using a fresh newly generated one every few weeks.

### Updates and News

**Updated: 11 July 2015** v0.5 update includes the addition of mix case output.

**Updated: 20 June 2015** typo fixes kindly reported by GitHub user @RickCogley.

**Update 21 March 2015:** the larger pool of three letter words used in
  the c language version (instead of Go) available here:
  [sugpass](https://github.com/wiremoons/sugpass) has now been added to this
  application also.

**Update 22 Jan 2015:** there is a newer version of this program with
  a larger pool of three letter words. The application is written in
  the c language (instead of Go) and is available here:
  [sugpass](https://github.com/wiremoons/sugpass)

For more information see the related blog posts here:

- [2014-12-09-Three-Letter-Word-Passwords](http://www.wiremoons.com/posts/2014-12-09-Three-Letter-Word-Passwords)
- [2015-01-22-Easy-Passwords-Everywhere](http://www.wiremoons.com/posts/2015-01-22-Easy-Passwords-Everywhere)

### Application Usage

The program is run from a command prompt&mdash;so on Windows using
Powershell or the Command Prompt, and on Linux or MacOSX in a Terminal
window using a shell such as bash. When the progam is run without any
command line arguments, it displays the following output:

```

                  THREE WORD - PASSWORD GENERATOR
                  ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯
• Number of three letter words available in the pool is: 1312
• Number of three letter words to include in the suggested password is: 3
        • Password character length will therefore be: 9
• Mixed case passwords to be provided: true
• Offering 3 suggested passwords for your consideration:

        | kos ads ods |   | kosadsods |   | KOsADSoDs |
        | elf gal nas |   | elfgalnas |   | elfgaLNAs |
        | lye meh bee |   | lyemehbee |   | lYeMehBeE |

To change the password suggestion output shown above, use the command line options.
Run the program as follows for more help:  passgen.exe -h

All is well
```
If you need to generate a 12 character password randomly created from three
letter words&mdash;perhaps as part of a command line pipe, you could use this
following options:
```
./passgen -q -r -w 4
jeealtmimwha
```
Or as above - but include random mixed case in the password also:
```
./passgen -q -r -w 4 -c
paxGoADEvbIt
```
or just generate a random block of output from the three letter words, and
include mixed case:
```
./passgen -q -r -w 200 -c
ONyKORMirpLunoGOPtYePpUygJuROWsonSoxBaTLACarkDiGdaLDuNdSOLYmfiXhAYyaDPASDUo
FagthegOovLyPasDALhoeEmSwrYHAJlaGCaWSUNSalPulISmUTuCoPtoOGASlowlEumacFAAEnd
wEtYAKaInGIfIDENixvAReRSaFtFaadAganISUQmUddElAYSFeHmAwDOmshABUswedTetVexDzO
LiglEpUrNYoNGaTKosOSEAlBdAMpULHOGERMTskkiPdIStsKCUTAysRoEawNganIFFThyARtgoX
aLFDIfLAtNAssOvhOytUTuEYLeeLoXrETGIDnAtRagORDtOpdIFcadDoeZuztOElYMErEYeHwaT
weeNothONdeWpiAEaUDodSEgKiSUDSmiSpsTSisRoceHsReClODPIgUSEPEWfesmOcjiGnOtoUR
oPsmAareXDoGMAmaiSTidJEthotgNutHYMoNMoEmOIBAlauawiTorcsUbnubDOTTIgVETbOSbOP
SALcoxLumEcuaFToOFuLEkIfYOBohsAbaGUSfuNgOOHonZaSSAZWAEtoCfeWREDCABPosvAVaah
```


The full list of command line arguments available are shown below, and also when you run: `passgen -help`:
```
Usage of ./passgen:
  -c=false:     USE: '-c=true' to get mixed case passwords. Note: useful with -q only [DEFAULT: lowercase]
  -h=false:     USE: '-h' to provide more detailed help about this program
  -q=false:     USE: '-q=true' to obtain just ONE password - no other screen output [DEFAULT: additional info output]
  -r=false:     USE: '-r=true' to remove spaces. Note: useful with -q only [DEFAULT: with spaces]
  -s=3:         USE: '-s no.' where no. is the number of password suggestions offered [DEFAULT: 3]
  -v=false:     USE: '-v=true.' display the application version [DEFAULT: false]
  -w=3:         USE: '-w no.' where no. is the number of three letter words to use [DEFAULT: 3]
```
The command line options are explained in more detail below:
- **-c** : 'c' stands for 'case'. Used to get mixed case passwords. **Note:** useful with `-q` only
- **-h** : 'h' stands for 'help'. if run with the `-h` option a screen of help text will be displayed
- **-s** : 's' stands for 'suggestion'. By default three passwords will be suggested. Change by adding a different number, so `-s 5` would provide five passwords.
- **-w** : 'w' stands for 'word'. By default the suggested passwords consist of three x three letter words, so 9 characters in length. If you wanted a longer password length, you can chnage the number of words&mdash;so using `-w 4` would provide four words instead, giving a password length of 12 characters.
- **-q** : 'q' stands for 'quiet'. This option only outputs ONE password (optionally at the length specified with -w) and no other text, so useful for using with command line pipes. Use with option `-r` to also remove spaces in the password and the `-c` options to obtain a mixed case password suggestion.
- **-r** : 'r' stands for 'remove'. This options removes any spaces from the password suggestions that are output **Note:** useful with `-q` only
- **-v** : 'v' stands for 'version'. This options only outputs the version of the application

### Downloading the Application

Pre-compiled binaries are available from the Release page below. These are provide for Windows (32bit and 64bit), MacOSX (64bit), and Linux (64bit):

- [passgen Release 0.5](https://github.com/wiremoons/passgen/releases/tag/0.5)
- [passgen Release 0.4](https://github.com/wiremoons/passgen/releases/tag/0.4)


### Compiling the Program

Assuming you already have Go installed and set-up on your computer—you
only need to download the single source file '`passgen.go`'. This can
then be built using the command below, assuming the `passgen.go` file
is in you current directory:

```
go build ./passgen.go
```

There is also a `Makefile` that I use on a computer running Linux to
cross compile the program for Linux (64 bit version), Windows (32 bit
and 64 bit versions), FreeBSD (64 bit version), and Mac OS X (64 bit version).
This can be done (assuming you have your computer and Go set-up correctly) by also
downloading the 'Makefile', and then entering:
```
make all
```

## To Do

The following enhancements are planned - see the source code also for any other *TODO* suggestions:

- TODO - maybe check for newer version and update if needed?

## License

The program is licensed under the "New BSD License" or "BSD 3-Clause
License". A copy of the license is available
[here](https://github.com/wiremoons/passgen/blob/master/LICENSE).

## OTHER INFORMATION

- Latest version is kept on GitHub here: [Wiremoons GitHub Pages](https://github.com/wiremoons)
- The program is written in Go - more information here: [Go](http://www.golang.org/)
- The program was written by Simon Rowe, licensed under [New BSD License](http://opensource.org/licenses/BSD-3-Clause)
