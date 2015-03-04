# gopastebin3
This version of Pastebin uses PrismJS syntax highlighting. I tried to figure Bootstrap out. Logging out invalidates the cookie. I used MongoBD for storing the pastes. 

# Dependencies

<code>go get github.com/dgrijalva/jwt-go</code><br>
<code>go get code.google.com/p/goauth2/oauth</code><br>
<code>go get gopkg.in/mgo.v2</code><br>
<code>go get gopkg.in/mgo.v2/bson</code>

# Encryption

Use openssl to make a private key called "demo.rsa" and a related public key called "demo.rsa.pub". Put them in the "static" directory.
