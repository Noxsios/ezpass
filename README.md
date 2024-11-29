# ezpass

<https://xkcd.com/936/>

![demo](https://github.com/user-attachments/assets/ccdee5f4-04ad-42a3-b0c3-7584a0d1ed0a)

Replicate the following in a zero dependency statically linked go binary:

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
