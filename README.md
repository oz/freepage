# FreePage

Freepage sends SMS messages from e-mail messages, or the CLI.

This program uses French operator [Free mobile][free]'s little SMS API to send
text messages to your own phone.

Use this when you need to be paged by long-running scripts, alert systems, etc.

[free]: http://mobile.free.fr/

# Usage

Send a text with:

```
freepage -u USER -s SECRET -m "Hello, world!"
```

Pipe an e-mail to the program to get its subject and body sent:

```
freepage -u USER -s SECRET -m -
```

# License

See LICENSE file, or http://unlicense.org
