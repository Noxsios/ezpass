# ezpass

<https://xkcd.com/936/>

First goal is to replicate the following in a statically linked go binary:

```bash
grep "^[^']\{3,5\}$" /usr/share/dict/words | shuf -n4 | paste -sd'.' -
```

