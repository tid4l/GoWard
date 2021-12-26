package server

import (
	"main/adminpanel"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	HashedPassword      = ""
	impersonatedWebpage = ""
	PanelLinks          = []string{"", ""}
	loginTime           = "N/A"
	upgrader            = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// Attach functions to the proxy.
func AttachFunctions(prox *mux.Router) {
	prox.HandleFunc(PanelLinks[1], internalPageHandler)
	prox.HandleFunc(PanelLinks[0], panelLoginPageHandler)
	prox.HandleFunc("/", indexPageHandler)
	prox.HandleFunc("/ws", websocketHandler)

	prox.HandleFunc("/login", loginFuncHandler).Methods("POST")
	prox.HandleFunc("/logout", logoutFuncHandler).Methods("POST")
}

// Login function handler
func loginFuncHandler(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	pass := req.FormValue("password")

	redirectTarget := PanelLinks[0]
	if name != "" && pass != "" {
		if name == "admin" && adminpanel.PasswordMatch(HashedPassword, pass, adminpanel.Salt) {
			Logger.Printf("[WARNING] Successful admin panel login.")
			redirectTarget = PanelLinks[1]
			adminpanel.WebData.LastLogin = loginTime
			loginTime = time.Now().UTC().Format("01-02-2006 15:04:05 UCT")
			adminpanel.SetSession(name, w)
		} else {
			Logger.Printf("Admin panel login attempt. Username: %s / Password: %s", name, pass)
		}
	}
	http.Redirect(w, req, redirectTarget, 302)
	t, err := template.New("hello").Parse(adminpanel.LoginPage)
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

// Logout function handler
func logoutFuncHandler(response http.ResponseWriter, request *http.Request) {
	adminpanel.ClearSession(response)
	http.Redirect(response, request, PanelLinks[0], 302)
}

// Login page handler function
func panelLoginPageHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("hello").Parse(adminpanel.LoginPage)
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
	Logger.Printf("[WARNING] New admin panel request from %s. Was this you?\n", req.RemoteAddr)
}

// Default index handler function
func indexPageHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("hello").Parse(impersonatedWebpage)
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
	Logger.Printf("New web request from %s\n", req.RemoteAddr)
	adminpanel.WebData.RequestCounter++
}

// Internal Admin panel handler function
func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := adminpanel.GetUserName(request)
	t, err := template.New("hello").Parse(adminpanel.InternalPage)
	if err != nil {
		panic(err)
	}
	if userName != "" {
		adminpanel.WebData.RequestHost = request.Host
		t.Execute(response, &adminpanel.WebData)
	} else {
		http.Redirect(response, request, PanelLinks[0], 302)
	}
}

// Admin panel websocket handler function
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			Logger.Println(err)
		}
		return
	}
	Logger.Println("Web socket served")
	go adminpanel.WsWriter(ws)
}
