[![NewBSD License](http://img.shields.io/badge/license-NewBSD-orange.svg?style=flat-square)](http://opensource.org/licenses/BSD-3-Clause)

## THREE WORD - PASSWORD GENERATOR

## Application Summary

A simple application written to generate a random password created
from a pool of English three letter words.

### About

This application will generate password suggestions based on a pool of
several hundred three letter English words. The words are selected
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

For more information see the related blog post here:
[2014-12-09-Three-Letter-Word-Passwords](http://www.wiremoons.com/posts/2014-12-09-Three-Letter-Word-Passwords)

### Application Usage

The program is run from a command prompt&mdash;so on Windows using
Powershell or the Command Prompt, and on Linux or MacOSX in a Terminal
window using a shell such as bash. When the progam is run without any
command line arguments, it displays the following output:

```
			THREE WORD - PASSWORD GENERATOR
			¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯
• Number of three letter words available in the pool is: 573
• Number of three letter words to include in the suggested password is: 3
	• Password character length will therefore be: 9
• Mixed case passwords to be provided: false
• Offering 3 suggested passwords for your consideration:

	 din wry ran
	 ova ram lit
	 ski yea koa

All is well
```
The command line argument available are shown as below, when you run: `passgen --help`: 
```
Usage of ./passgen:
  -c=false:     USE: '-c=true' to get mixed case passwords [DEFAULT: lowercase]
  -h=false:     USE: '-h' to provide more detailed help about this program
  -s=3:         USE: '-s no.' where no. is the number of password suggestions offered [DEFAULT: 3]
  -w=3:         USE: '-w no.' where no. is the number of three letter words to use [DEFAULT: 3]
  -q=false:     USE: '-q=true' to obtain just ONE password - no other screen output [DEFAULT: additonal info output]
```
The command line options are explained in more detail below:
- **-c** : 'c' stands for 'case'. by default the suggested passwords are shown in lowercase only. Run with `-c=true` to get mixed case suggestions
- **-h** : 'h' stands for 'help'. if run with the `-h` option a screen of help text will be displayed
- **-s** : 's' stands for 'suggestion'. By default three passwords will be suggested. Change by adding a different number, so `-s 5` would provide five passwords. 
- **-w** : 'w' stands for 'word'. By default the suggested passwords consist of three x three letter words, so 9 characters in length. If you wanted a longer password length, you can chnage the number of words&mdash;so using `-w 4` would provide four words instead, giving a password length of 12 characters.
- **-q** : 'q' stands for 'quiet'. This options only outputs ONE password (optionally at the length specified with -w) and no other text, so useful for using with command line pipes.

## To Do

The following enhancements are planned - see the source code also for any other *TODO* suggestions:

- TODO - maybe check for newer version and update if needed?
- TODO - add -c mixed case output mode
- TODO - add option to remove spaces between three letter words in output

## License

The program is licensed under the "New BSD License" or "BSD 3-Clause
License". A copy of the license is available
[here](https://github.com/wiremoons/passgen/blob/master/License.txt).

## OTHER INFORMATION

- Latest version is kept on GitHub here: [Wiremoons GitHub Pages](https://github.com/wiremoons)
- The program is written in Go - more information here: [Go](http://www.golang.org/)
- The program was written by Simon Rowe, licensed under [New BSD License](http://opensource.org/licenses/BSD-3-Clause)


