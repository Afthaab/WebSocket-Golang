package handler

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := RenderPage(w, "home.jet", nil)
	if err != nil {
		log.Println("Error : ", err)
	}
}

//tmpl is the templete to render and data is the data needed to pass to the template
func RenderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println("Error : ", err)
		return err
	}
	err = view.Execute(w, data, nil) // the third variable is the context which has been ignored in this course
	if err != nil {
		log.Println("Error : ", err)
		return err
	}
	return nil
}

// WsResponse defines the response sent from WebSocket
type WsResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"messagetype"`
}

//Websocket connection struct
type WebSokcetConnnection struct {
	*websocket.Conn
}

// WsIncoming strcut
type WsPayload struct {
	Action   string               `json:"action"`
	Message  string               `json:"message"`
	Username string               `json:"username"`
	Conn     WebSokcetConnnection `json:"-"`
}

// WsEndpoint upgrades the socket connection
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error : ", err)
	}
	log.Println("Client Side Connection has been established")
	var res WsResponse
	res.Message = `<em><small>Connected to server</small></em>`
	err = ws.WriteJSON(res)
	if err != nil {
		log.Println("Error: ", err)
	}

}
