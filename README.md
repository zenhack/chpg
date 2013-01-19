# Correct Horse Password Generator

## Background Reading

![http://xkcd.com/936/](http://imgs.xkcd.com/comics/password_strength.png
"To anyone who understands information theory and security and is in an
infuriating argument with someone who does not (possibly involving mixed
case), I sincerely apologize.")

([Source](http://xkcd.com/936/), in case anyone was unaware).

## Why?

Because you are not a good enough random number generator for the above to be
safe. And because most of the password generators I've seen generate
"troubador" passwords, or ones that are *harder* to remember.

## Details (not many)

This is just a simple password generator - It generates passwords like the one
praised in the comic.

## Building & installing

	go get github.com/zenhack/chpg

## Running

 - Make sure $GOPATH/bin is in your $PATH
 - Make sure you have a valid dictionary at /usr/share/dict/words, and type: 

	chpg

 - Or, make sure you have one /somewhere/else, and type:

	chpg -d /somewhere/else

 - This should work with windows style paths if you're using that, but I
   haven't tried it.

 - The -n flag changes the number of words used in the password.

## Notes

 - The dictonary should just be a newline-separated list of words. The file
   should end with a newline.
 - Using a dictionary that's large enough is is important. The general
   consensus is that Randall used a 2000-word dictionary in his calculations,
   which is fairly small; most dictionaries intended for other uses will be
   larger.
 - This uses the go standard library's cryptographic random number generator,
   Which slows things down a bit (Just barely perceptible on my laptop), and
   so this is probably not something you want to use in a performance sensitive
   context.

## License

MIT. See COPYING.
