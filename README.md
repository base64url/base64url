# base64url

This is a simple commandline utility aimed to mimic the behavior of coreutils/base64 using the URL and Filename Safe Alphabet as specified in [RFC 4648](https://tools.ietf.org/html/rfc4648).

# Usage

Encode a String 

    echo foobar | base64url

Decode a String

    echo Zm9vYmFy | base64url -d
