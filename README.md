# BingHomepage.Go

BingHomepage.Go library is powered by BingHomepageAPI which provides Bing's current homepage image details which include URL for image, Copyright information, and a Copyright link.

## Usage

### Installing library

To get this library, run `go get`

```sh
go get github.com/BingHomepage/BingHomepage.Go
```

### Importing

In the import section of your Go source, include `github.com/BingHomepage/BingHomepage.Go`

```go
import "github.com/BingHomepage/BingHomepage.Go"
```

### Calling `Get` function

To get information, you need to call a function defined in the library named `Get`. It returns two values, `BingHomepage` structure and `error` details.

`Get` take one argument which is two letter country code of Bing Region.

```go
data, err := BingHomepage.Get("US")
if err != nil {
	panic(err)
}
```

#### BingHomepage Structure

```go
type BingHomepage struct {
	URL           string  `json:"url"`
	Copyright     string  `json:"copyright"`
	CopyrightLink string  `json:"copyrightlink"`
}
```

## Example

```go
package main

import (
	"fmt"
	"github.com/BingHomepage/BingHomepage.Go"
)

func main() {
	data, err := BingHomepage.Get("US")
	if err != nil { panic(err) }
	fmt.Printf("%+v\n", data)
	fmt.Printf("URL\t\t=> %s\nCopyright\t=> %s\nCopyrightLink\t=> %s\n", data.URL, data.Copyright, data.CopyrightLink)
}
```

Expected output

```
&{URL:https://bing.com/az/hprichbg/rb/CometMoth_EN-US9387578049_1920x1080.jpg Copyright:Comet moth in Ranomafana National Park, Madagascar (© Robin Hoskyns/Minden Pictures) CopyrightLink:http://www.bing.com/search?q=comet+moth&form=hpcapt&filters=HpDate:%2220180721_0700%22}
URL             => https://bing.com/az/hprichbg/rb/CometMoth_EN-US9387578049_1920x1080.jpg
Copyright       => Comet moth in Ranomafana National Park, Madagascar (© Robin Hoskyns/Minden Pictures)
CopyrightLink   => http://www.bing.com/search?q=comet+moth&form=hpcapt&filters=HpDate:%2220180721_0700%22
```
