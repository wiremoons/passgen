[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hyperium/hyper/master/LICENSE)

## THREE WORD - PASSWORD GENERATOR

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

The application outputs mixed case passwords also. The switching of certain password letters to upper case is also done randomly, so as to enhance the overall entropy of the resultant password. 

Additionally with each password suggested, a randomly generated number is provided, which can be include it in the password you select from the outputs, should you wish.

### Application Usage

The program is run from a command prompt&mdash;so on Windows using
Powershell or cmd.exe, and on Linux or MacOSX in a Terminal
window using a shell such as bash. 

When the program is run without any command line arguments, it displays the following output:

```

                        THREE WORD - PASSWORD GENERATOR
                        ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯
• Number of three letter words available in the pool is: 1312
• Number of three letter words to include in the suggested password is: 3
        • Password character length will therefore be: 9
• Mixed case passwords to be provided: true
• Offering 3 suggested passwords for your consideration:

        off due lar    offduelar    OFfdUelAr    50
        lep poa fid    leppoafid    lEpPoAfid    19
        mos box sei    mosboxsei    moSbOXSEi    15

To change the password suggestion output shown above, use the command line options.
Run the program as follows for more help:  passgen.exe -h

All is well
```
Using the command line parameters, you can also use `passgen` in the following ways:

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
- **-w** : 'w' stands for 'word'. By default the suggested passwords consist of three x three letter words, so 9 characters in length. If you wanted a longer password length, you can change the number of words&mdash;so using `-w 4` would provide four words instead, giving a password length of 12 characters.
- **-q** : 'q' stands for 'quiet'. This option only outputs ONE password (optionally at the length specified with -w) and no other text, so useful for using with command line pipes. Use with option `-r` to also remove spaces in the password and the `-c` options to obtain a mixed case password suggestion.
- **-r** : 'r' stands for 'remove'. This options removes any spaces from the password suggestions that are output **Note:** useful with `-q` only
- **-v** : 'v' stands for 'version'. This options only outputs the version of the application

### Downloading the Application

Pre-compiled binaries are available from the Release page below. These are provide for Windows (32bit and 64bit), MacOSX (64bit), and Linux (64bit):

- [passgen Release 0.7.0](https://github.com/wiremoons/passgen/releases/tag/0.7.0)
- [passgen Pre-Release 0.5](https://github.com/wiremoons/passgen/releases/tag/0.5)
- [passgen Pre-Release 0.4](https://github.com/wiremoons/passgen/releases/tag/0.4)

Final compiled binary versions for Linux, Windows, macOS and FreeBSD are available in the `binaries/` sub folder
of this GitHub repo. The functionality is the same as the *0.7.0* release. Future versions of the application will 
have an adjusted appearance - so please one of these binaries if you dislike the newer version.

### Compiling the Program

Assuming you already have Go installed and set-up on your computer—you
can download and install with the go command:

```
go get -u -v github.com/wiremoons/passgen
```

There is also a `Makefile` that I use to cross compile the program for Linux (32 and 64 bit versions for Intel 
and ARM/Raspberry Pi), Windows (32 bit and 64 bit versions for Intel and 64 bit ARM), FreeBSD (64 bit version), 
and macOS (64 bit versions for Intel and Apple Silicon arm64).

This can be done (assuming you have `make` installed your computer, and Go set-up correctly) by downloading the source zip file or Git cloning the `passgen` repo with the command:

```
git clone https://github.com/wiremoons/passgen/
```

After you have a copy of the source code, you can use the 'Makefile'. Read the Makefile for the options - but to get a copy of all the binary version run:

```
make all
```

You can also run the command: `make help` for further assistance.

## To Do

The following enhancements are planned in the future:

- TODO - maybe check for newer version and update if needed?
- TODO - add support for tags and better versioning info on compile
- TODO - update the interface with colour output options etc.


## License

This program is licensed under the "MIT License". See
http://opensource.org/licenses/mit for more details.

A copy of the license is available
[here](https://github.com/wiremoons/passgen/blob/master/LICENSE).

## OTHER INFORMATION

- Latest version is kept on GitHub here: [Wiremoons GitHub Pages](https://github.com/wiremoons)
- The program is written in Go - more information here: [Go](http://www.golang.org/)
- The program was written by Simon Rowe, and is licensed under [The MIT License (MIT)](http://opensource.org/licenses/mit)
