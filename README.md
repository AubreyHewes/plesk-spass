# A Plesk Spam Assistant

A `postfix/pipe` for discarding spam rated above a certain threshold.

> The mail is actually _deferred_ and will eventually be discarded

## Parameters

|Parameter|Description|Default|
|---|---|---|
|-t|The threshold|50|
|-D|A dir to write all discarded mails to for later perusal or for spam/ham learning purposes||

## Examples

### Temporary (i.e. for testing et al)

> NOTE Plesk rewrites the following file when configuration is changed via the panel

Edit `/var/qmail/mailnames/DOMAIN/USER/.qmail`

And add the following

    |/usr/bin/spass -t X

Where `X` is the spam threshold to discard

> If forwarding add it before the forward!

## Building

    make

This creates `dist/spass` and `dist/spass-debug`

The `dist/spass` is a [crunched](https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/) binary (by about ~80%) 

## External build requirements

 * [upx](https://upx.github.io/) for the release build
