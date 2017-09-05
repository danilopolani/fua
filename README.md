# Go FUA
![GitHub version](https://badge.fury.io/gh/danilopolani%2Fgosc.svg)
[![GoDoc](https://godoc.org/github.com/danilopolani/gosc?status.svg)](https://godoc.org/github.com/danilopolani/gosc)
[![GoReport](https://goreportcard.com/badge/github.com/danilopolani/gosc)](https://goreportcard.com/report/github.com/danilopolani/gosc)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go FUA, which stands for **F**ake **U**ser **A**gent, allows you to retrieve a fake user agent or to generate a random one from a given group (or from all).

### Installation

To install Go FUA, just run this command from your terminal: `go get github.com/danilopolani/go-fua`.  
Then, include it in your project: `import "github.com/danilopolani/go-fua"`.

### Generate a random user agent

User agents are grouped by their environment: [desktop](#desktop), [mobile](#mobile), [tablet](#tablet), [console](#game-consoles), [ereader](#e-readers) and [bot](#bots). You can see the available user agents clicking on the group name or going to the end of the document.  
  
To generate a random UA, you have to use the `Random()` function passing the your group as parameter or `all` to catch everything.

```go
package main

import (
  "fmt"
  "github.com/danilopolani/go-fua"
)

func main() {
  rand1 := fua.Random("all") // One from desktop, tablet, mobile, console, ereader and bot
  rand2 := fua.Random("console") // A random console
  
  fmt.Println(rand1)
  fmt.Println(rand2)
}
```

### Get a specific User Agent

If you don't want to generate a random UA, you can retrieve a specific one with its alias from the `UA` public variable. You can see a list of alias scrolling down.   
  
The path to access a user agent is: `UA.<group>.<alias>`.

```go
package main

import (
  "fmt"
  "github.com/danilopolani/go-fua"
)

func main() {
  // Retrieve googlebot
  fmt.Println(fua.UA["bot"]["google"]) // Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)
  // Retrieve HTC browser
  fmt.Println(fua.UA["mobile"]["htc"]) // Mozilla/5.0 (Linux; Android 6.0; HTC One M9 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36
}
```

### List of groups and alias

#### Desktop
Group: `desktop`  
  
- **Alias**: `edge`  
  **O.S.**: Windows 10  
  **Browser**: Microsoft Edge  
- **Alias**: `chromebook`  
  **O.S.**: Chromebook  
  **Browser**: Google Chrome  
- **Alias**: `safari`  
  **O.S.**: Mac OSX  
  **Browser**: Safari  
- **Alias**: `win_chrome`  
  **O.S.**: Windows 7  
  **Browser**: Google Chrome  
- **Alias**: `firefox`  
  **O.S.**: Ubuntu  
  **Browser**: Firefox

#### Tablet
Group: `tablet`  
  
- **Alias**: `pixel`  
  **O.S.**: Android on Google Pixel C  
  **Browser**: Chrome  
- **Alias**: `xperia`  
  **O.S.**: Android on Xperia Z4 Tablet 
  **Browser**: Google Chrome  
- **Alias**: `nvidia`  
  **O.S.**: Android on Nvidia Shield Tablet  
  **Browser**: Chrome  
- **Alias**: `samsung`  
  **O.S.**: Android on Samsung Galaxy Tab A  
  **Browser**: Samsung Browser  
- **Alias**: `kindle`  
  **O.S.**: Android on Amazon Kindle Fire HDX 7  
  **Browser**: Chrome  
- **Alias**: `lg`  
  **O.S.**: Android on LG G Pad 7.0  
  **Browser**: Chrome

#### Mobile
Group: `mobile`  
  
- **Alias**: `s6`  
  **O.S.**: Android on Samsung Galaxy S6  
  **Browser**: Chrome  
- **Alias**: `s6_edge`  
  **O.S.**: Android on Galaxy S6 Edge Plus 
  **Browser**: Chrome  
- **Alias**: `lumia`  
  **O.S.**: Windows Phone on Lumia  
  **Browser**: Microsoft Edge  
- **Alias**: `nexus`  
  **O.S.**: Android on Nexus 6P  
  **Browser**: Chrome  
- **Alias**: `xperia`  
  **O.S.**: Android on Sony Xperia Z5  
  **Browser**: Chrome  
- **Alias**: `htc`  
  **O.S.**: Android on HTC One M9  
  **Browser**: Chrome

#### Game Consoles
Group: `console`  
  
- **Alias**: `wii`  
  **Platform**: Wii U  
  **Browser**: Nintendo Browser  
- **Alias**: `xbox`  
  **Platform**: Xbox One 
  **Browser**: Microsoft Edge  
- **Alias**: `ps4`  
  **Platform**: PlayStation 4  
  **Browser**: -  
- **Alias**: `psv`  
  **Platform**: PlayStation Vita  
  **Browser**: Silk  
- **Alias**: `ds`  
  **Platform**: Nintendo 3DS  
  **Browser**: -

#### E-readers
Group: `ereader`  
  
- **Alias**: `kindle4`  
  **Platform**: Kindle 4  
  **Browser**: Kindle?  
- **Alias**: `kindl3`  
  **Platform**: Kindle 3  
  **Browser**: Kindle?

#### Bots
Group: `bot`  
  
- **Alias**: `google`  
  **Bot**: Google Bot    
- **Alias**: `bing`  
  **Bot**: Bing bot    
- **Alias**: `yahoo`  
  **Bot**: Yahoo bot
