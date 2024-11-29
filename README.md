# ezpass

<https://xkcd.com/936/>

First goal is to replicate the following in a statically linked go binary:

```bash
grep "^[^']\{3,5\}$" /usr/share/dict/words | shuf -n4 | paste -sd'.' -
```

```bash
go install github.com/noxsios/ezpass
```

```bash
$ ezpass -h
  -d string
      Delimiter between words (default ".")
  -h	Print this message and exit.
  -n int
      Number of words in resulting password (default 4)
```
