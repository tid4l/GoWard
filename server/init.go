package server

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"main/adminpanel"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gorilla/mux"
)

var proxies = make(map[string]*httputil.ReverseProxy)

// NOTE: This is the string array that GoWard will use to select a
// random web page to impersonate. Just add URLs to this
// array for them to be chosen.
// Example: []string {
//       "https://www.webpagetoimpersonate.com/",
//       "https://www.anotherpagetoimpersonate.com/"}
var DefaultTargetURLs = []string{
	"",
	"",
	""}

func InitializeServer(password *string, hostProxy map[string][]string, prox *mux.Router, targetURL *string) *mux.Router {
	impersonatedWebpage = impersonateWebpage(targetURL)
	PanelLinks = panelLinkGenerator()
	HashedPassword = adminpanel.HashPassword(*password, adminpanel.Salt)

	// Initialize Misc variables
	adminpanel.WebData.TotalHosts = 0
	adminpanel.WebData.RequestCounter = 0
	adminpanel.WebData.OnlineHosts = 0

	for k, v := range hostProxy {
		remote, err := url.Parse(v[0])
		if err != nil {
			Logger.Fatal("Unable to parse proxy target")
		}
		proxies[k] = httputil.NewSingleHostReverseProxy(remote)
		adminpanel.WebData.TotalHosts++
	}

	// Instantiate proxies
	for host, proxy := range proxies {
		prox.Host(host).Handler(proxy)
	}

	// Initialize web functions
	AttachFunctions(prox)

	// Create channel, begin healthchecks with goroutine, pass result of healthcheck through channel.
	channel := make(chan int, 1)
	go healthCheck(hostProxy, channel)
	go func() {
		for {
			adminpanel.WebData.OnlineHosts = <-channel
		}
	}()

	return prox
}

// Impersonates a webpage and save to html template.
func impersonateWebpage(targetURL *string) string {
	var (
		response *http.Response
		url      string
		err      error
	)
	if *targetURL != "" {
		url = *targetURL
	} else {
		b := make([]byte, 1)
		rand.Read(b)
		n := int(b[0]) % len(DefaultTargetURLs)
		url = DefaultTargetURLs[n]
	}
	response, err = http.Get(url)
	Logger.Printf("Serving impersonated webpage: %s", url)
	fmt.Printf("Serving impersonated webpage: %s\n", url)
	if err != nil {
		Logger.Fatal(err)
	}
	defer response.Body.Close()
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	return pageContent
}

// Randomly generate the links for the admin panel login and
func panelLinkGenerator() []string {
	var links = []string{"", ""}
	// Randomly generate admin panel logon link

	for i := 0; i < 4; i++ {
		b := make([]byte, 1)
		rand.Read(b)
		v := (int(b[0]) % (4 + i)) + 6
		result := ""
		for j := 0; j < v; j++ {
			b := make([]byte, 1)
			rand.Read(b)
			randomI := int(b[0]) % 2
			var randomC rune
			b = make([]byte, 1)
			rand.Read(b)
			if randomI == 1 {
				randomC = 'A' + rune(int(b[0])%26)
			} else {
				randomC = 'a' + rune(int(b[0])%26)
			}
			result = result + string(randomC)
		}
		links[0] = links[0] + "/" + result
	}

	//Randomly generate internal admin panel link
	for i := 0; i < 3; i++ {
		b := make([]byte, 1)
		rand.Read(b)
		randomI := int(b[0]) % 2
		var randomC rune
		b = make([]byte, 1)
		rand.Read(b)
		if randomI == 1 {
			randomC = 'A' + rune(int(b[0])%26)
		} else {
			randomC = 'a' + rune(int(b[0])%26)
		}
		links[1] = links[1] + string(randomC)
	}
	links[1] = links[0] + "/" + links[1]

	Logger.Printf("Admininstration panel can be remotely accessed at " + links[0])
	fmt.Printf("Admininstration panel can be remotely accessed at " + links[0] + "\n")
	return links
}

// Checks the health of the C2 servers on the backend.
func healthCheck(proxies map[string][]string, channel chan int) {
	t := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-t.C:
			Logger.Printf("Starting health check...")
			count := 0
			for a, b := range proxies {
				status := "down"
				remote, _ := url.Parse(b[0])
				alive := isBackendAlive(remote)
				if !alive {
					status = "down"
					proxies[a] = []string{b[0], "down"}
				} else {
					status = "up"
					proxies[a] = []string{b[0], "up"}
				}
				Logger.Printf("%s [%s]\n", b[0], status)
				if b[1] == "up" {
					count++
				}
			}
			fmt.Printf("Backend health status: %d of %d online\n", count, adminpanel.WebData.TotalHosts)
			channel <- count
			Logger.Println("Health check completed")
		}
	}
}

// Determines if the backend is reachable.
func isBackendAlive(u *url.URL) bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", u.Host, timeout)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}
