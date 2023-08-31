![Version](https://img.shields.io/badge/version-0.0.0-orange.svg)

# Maps Utils

A basic golang package for demonstration purpose. Package currently contains only one function:

---

## Installation

Go to your project root, where `go.mod` file exists, than grab the library via:

```bash
go get github.com/seymasa/stringutils
```

---

## Usage

```go
package main

import (
	"fmt"
	"github.com/seymasa/stringutils"
	"log"
)

func main() {
	m1 := map[string]int{
		"key1": 12,
	}

	m2 := map[string]int{
		"key1": 12,
	}
	isEqual, err := mapsutils.DeepEqualMap(m1, m2)
	if err != nil {
		log.Fatal(err)
	}
	if isEqual {
		fmt.Println("Maps is equal")
	}
}
```

---

## Contributor(s)

* [seymasa](https://github.com/seymasa) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/seymasa/stringutils/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'add some functionality'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

---

## License

This project is licensed under GPLv3

---
