# grasshopper

Get GitHub contribution

# Usage

```
$ gh-contributions gomachan46
username,from,to,total,currentStreak,longestStreak
gomachan46,2015-08-16,2016-08-16,887,32,176
```

```
$ gh-contributions gomachan46 foo bar
username,from,to,total,currentStreak,longestStreak
gomachan46,2015-08-16,2016-08-16,887,32,176
foo,2015-08-16,2016-08-16,1,0,1
bar,2015-08-16,2016-08-16,2,0,2
```

# Requirements

* go

# Installation

```
$ go get github.com/gomachan46/gh-contributions/cmd/gh-contributions
```

# Lisence

MIT

# Author

Go Nakanishi (a.k.a gomachan46)
