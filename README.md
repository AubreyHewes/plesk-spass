# Postfix Spamassassin Assistant

A `postfix/pipe` for discarding spam rated above a certain threshold.

> The mail is actually _deferred_ and will eventually be discarded

## Parameters

|Parameter|Description|Default|
|---|---|---|
|-t|The threshold|50|
|-D|A dir to write all discarded mails to for later perusal or for spam/ham learning purposes (WIP)||

## Examples

### Temporary (i.e. for testing et al)

> NOTE Plesk rewrites the following file when configuration is changed via the panel

Edit `/var/qmail/mailnames/DOMAIN/USER/.qmail`

And add the following

    |/usr/bin/spass

> If forwarding add it before the forward!

## Building

    make

This creates `dist/spass` and `dist/spass-debug`

The `dist/spass` is a [crunched](https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/) binary (by about ~80%) 

## External build requirements

 * [upx](https://upx.github.io/) for the release build

### TODO

 * [ ] integrate upx into a local build process (possibly just build via docker)

## Testing (WIP)

### Development

`cat` a test fixture to the binary

For the pipe the following is valid;
 - no output would be [carry on delivering](https://en.wikipedia.org/wiki/Carry_On_(franchise)#Carry_On_films)
 - otherwise it would be deferred with the output message
 
        discarded spam SCORE/THRESHOLD

### Production

Send a [GTUBE mail](https://spamassassin.apache.org/gtube/) to a configured recipient

## Pronunciation

[IPA](https://en.wiktionary.org/wiki/Wiktionary:International_Phonetic_Alphabet)<sup>([key](https://en.wiktionary.org/wiki/Appendix:English_pronunciation))</sup> [/spæz/](https://en.wiktionary.org/wiki/spaz#Pronunciation)
