// Package fua generates a fake user agent
package fua

import (
	"math/rand"
	"time"
)

// UA contains all the user agents grouped by environment
var UA = map[string]map[string]string{
	"desktop": map[string]string{
		"edge":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246", // Windows 10 Edge
		"chromebook": "Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36",               // Chromebook Chrome
		"safari":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9",            // OSX Safari
		"win_chrome": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36",                   // Windows 7 Chrome
		"firefox":    "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1",                                                  // Linux Firefox
	},
	"tablet": map[string]string{
		"pixel":   "Mozilla/5.0 (Linux; Android 7.0; Pixel C Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/52.0.2743.98 Safari/537.36",               // Google Pixel C
		"xperia":  "Mozilla/5.0 (Linux; Android 6.0.1; SGP771 Build/32.2.A.0.253; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/52.0.2743.98 Safari/537.36",        // Xperia Z4 Tablet
		"nvidia":  "Mozilla/5.0 (Linux; Android 5.1.1; SHIELD Tablet Build/LMY48C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Safari/537.36",                       // Nvidia Shield Tablet
		"samsung": "Mozilla/5.0 (Linux; Android 5.0.2; SAMSUNG SM-T550 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/3.3 Chrome/38.0.2125.102 Safari/537.36", // Samsung Galaxy Tab A
		"kindle":  "Mozilla/5.0 (Linux; Android 4.4.3; KFTHWI Build/KTU84M) AppleWebKit/537.36 (KHTML, like Gecko) Silk/47.1.79 like Chrome/47.0.2526.80 Safari/537.36",            // Amazon Kindle Fire HDX 7
		"lg":      "Mozilla/5.0 (Linux; Android 5.0.2; LG-V410/V41020c Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/34.0.1847.118 Safari/537.36",        // LG G Pad 7.0
	},
	"mobile": map[string]string{
		"s6":      "Mozilla/5.0 (Linux; Android 6.0.1; SM-G920V Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",                          // Galaxy S6
		"s6_edge": "Mozilla/5.0 (Linux; Android 5.1.1; SM-G928X Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.83 Mobile Safari/537.36",                          // Galaxy S6 Edge Plus
		"lumia":   "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/13.10586", // Lumia 950
		"nexus":   "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 6P Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.83 Mobile Safari/537.36",                          // Nexus 6P
		"xperia":  "Mozilla/5.0 (Linux; Android 6.0.1; E6653 Build/32.2.A.0.253) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",                       // Xperia Z5
		"htc":     "Mozilla/5.0 (Linux; Android 6.0; HTC One M9 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",                          // HTC One M9
	},
	"console": map[string]string{
		"wii":  "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.4.2.12 NintendoBrowser/4.3.1.11264.US",                                              // Wii U
		"xbox": "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/13.10586", // Xbox One
		"ps4":  "Mozilla/5.0 (PlayStation 4 3.11) AppleWebKit/537.73 (KHTML, like Gecko)",                                                                                      // PlayStation 4
		"psv":  "Mozilla/5.0 (PlayStation Vita 3.61) AppleWebKit/537.73 (KHTML, like Gecko) Silk/3.2",                                                                          // PlayStation Vita
		"ds":   "Mozilla/5.0 (Nintendo 3DS; U; ; en) Version/1.7412.EU",                                                                                                        // 3DS
	},
	"ereader": map[string]string{
		"kindle4": "Mozilla/5.0 (X11; U; Linux armv7l like Android; en-us) AppleWebKit/531.2+ (KHTML, like Gecko) Version/5.0 Safari/533.2+ Kindle/3.0+", // Kindle 4
		"kindle3": "Mozilla/5.0 (Linux; U; en-US) AppleWebKit/528.5+ (KHTML, like Gecko, Safari/528.5+) Version/4.0 Kindle/3.0 (screen 600x800; rotate)", // Kindle 3
	},
	"bot": map[string]string{
		"google": "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",            // Google
		"bing":   "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",             // Bing
		"yahoo":  "Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)", // Yahoo
	},
}

// Random returns a random user agent from the choosen group
func Random(group string) string {
	var uas []string

	if group != "all" {
		g, exists := UA[group]
		if !exists {
			return Random("all")
		}

		uas = values(g)
	} else {
		for _, g := range UA {
			uas = append(uas, values(g)...)
		}
	}

	rand.Seed(time.Now().Unix())
	return uas[rand.Intn(len(uas))]
}

func values(m map[string]string) []string {
	var r []string

	for _, v := range m {
		r = append(r, v)
	}

	return r
}
