//GoWard
//By chdav

package main

import (
	"flag"
	"fmt"
	"main/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

var (
	hostProxy = make(map[string][]string)
	prox      = mux.NewRouter()

	//initialize flags
	targetURLInput = flag.String("target", "", "Optional.\nSpecify a target URL to impersonate and use. If none specificed, default will be used.\n\tEx: -target=https://www.somewebsite.co/")
	passwordInput  = flag.String("password", "", "Required.\nSpecify the password for the admin panel.\n\tEx: -password=pass")
	proxiesInput   = flag.Int("proxies", 0, "Required.\nSpecify the number of proxies.\n\tEx: -proxies=3")
)

// Automatically runs upon program start.
func init() {

	programCloseHandler()

	// Parse command line arguments.
	flag.Parse()
	if *passwordInput == "" || *proxiesInput == 0 {
		fmt.Printf("Missing arguments.\n\n")
		flag.Usage()
		os.Exit(1)
	}

	banner()

	// Read in user inputted proxies.
	for proxies := 1; proxies <= *proxiesInput; proxies++ {
		var headerInput string
		var hostInput string
		fmt.Printf("Enter header for proxy %d: ", proxies)
		fmt.Scanln(&headerInput)
		fmt.Printf("Enter IP followed by port for proxy %d (I.E. http://IP:PORT): ", proxies)
		fmt.Scanln(&hostInput)
		hostProxy[headerInput] = []string{hostInput, "down"}
		fmt.Print("\n")
	}

	server.Logger.Println("Server started.")
	fmt.Printf("\nServer started. For more verbose output, see log file: %s\n", server.LogFileName)

	//Initialize the server and proxies.
	prox = server.InitializeServer(passwordInput, hostProxy, prox, targetURLInput)
}

// Print banner.
func banner() {
	banner := `
               
	           _        _      
	          | \__/\__/ |
	   ___    |  '.||.'  |             _
	  / _ \___|__/ || \__|__ _ _ __ __| |
	 / /_\/ _ \--\ || /--/ _' | '__/ _' |
	/ /_\\ (_) \  \||/  / (_| | | | (_| |
	\____/\___/ \  ||  / \__,_|_|  \__,_|
	             '.||.'


			GoWard (v0.1)
						
		`

	fmt.Println(banner)
}

// Handles interrupts from system and Ctrl+C; gracefully exits program.
func programCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Exiting program...")
		server.Logger.Printf("Program exited.\n\n")
		server.LogFile.Close()
		os.Exit(0)
	}()
}

func main() {
	// Start server and proxies.
	server.Logger.Println("Starting proxy...")
	_ = http.ListenAndServe(":80", prox)
}
