package main

import "fmt"

// PrintHelp function prints out some basic help information for the user
// that is diaplyed on the command line.
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
	While the passwords generated look far too simple and easy to be secure, the are
	in fact very secure, and difficult to crack. Just because they look simple to a
	human - it doesn't mean they are simple to work out using a computer. They are
	in fact quite hard to work out for a computer. The reason for this is that they
	are randomly generated, not a single dictionary word, or a single common name.
	This makes the password harder to 'find' as it is not commonly known.

	It is a common misconception that a password has to be 'complex' to be any good.
	Unfortunately we have been led to believe that the more complex a password 
	is - the better and more secure it will be - which is in fact wrong.

	In fact a longer password, that can more easily be remembered, and therefore 
	changed more frequently as a consequence, actually offers a far greater degree 
	of security.

	For more information and explanations of this, please see the web pages included
	below under 'References'. There are plenty of expert sources on the Internet
	also, that will explain the benefits and security of using a randomly generated
	three word (or more) combination password. Just remember - your password must
	be at least nine characters in total - or longer if possible. You can of
	course	always add additional punctuation, should you wish!

	So How Many Possible Passwords Are There?
	¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯
	There are over 1,000 three letter words in the pool that can be chosen from, and
	assuming you use three of these words combined, that provide 1,000^3 (1,000 to
	power of 3) possible combinations - of which one is your password.

	So - 1,000 x 1,000 x 1,000 = 1,000,000,000 (one billion) possibilities.

	If you use the mixed case option (upper and lower case) - then number increases
	further of course - and you can still add numbers, and/or punctuation characters
	if you wish too.

	Or just increase you password length to 12 characters, so use four of the three
	letter words, and you end up with 1,000,000,000,000 (one thousand billion)
	possibilities -and that is just lower case letters only.

	References
	¯¯¯¯¯¯¯¯¯¯
	Thomas Baekdal - The Usability of Passwords - FAQ
	 - http://www.baekdal.com/insights/the-usability-of-passwords-faq
	Steve Gibson - GRC 'How Big is Your Haystack?'
	 - https://www.grc.com/haystack.htm
	Application 'passgen' - author's web site
	 - http://www.wiremoons.com/

	`
	// now output the above to screen
	fmt.Println(helptext)
}
