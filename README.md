# ezpass

<https://xkcd.com/936/>

![demo](https://github.com/user-attachments/assets/ccdee5f4-04ad-42a3-b0c3-7584a0d1ed0a)

Replicate the following in a zero dependency statically linked go binary:

```bash
grep "^[^']\{3,5\}$" /usr/share/dict/words | shuf -n4 | paste -sd'.' -
```

```bash
CGO_ENABLED=0 go install github.com/noxsios/ezpass
```

```bash
$ ezpass -h
-n, --number-of-words int   Number of words in resulting password. (default 4)
-d, --delimiter string      Delimiter between words. (default ".")
-v, --version               Print the version number of ezpass and exit.
-h, --help                  Print this message and exit.
```
