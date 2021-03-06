package main

import (
	"flag"
	"net/http"

	"github.com/YutakaHorikawa/gows/config"
	"github.com/YutakaHorikawa/gows/server"
	"github.com/YutakaHorikawa/gows/ws"
)

func main() {
	flag.Parse()
	router := server.NewRouter()
	config := config.NewConfig()
	hm := ws.NewHubManager(config.Hub.Worker)

	hm.RunAllHub()

	router.HandleFunc("/room/{id:[1-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := server.Vars(r)
		roomId := vars["id"]
		hub := hm.GetHubByRoomid(roomId)
		if hub == nil {
			hub = hm.GetHub()
		}
		ws.ServeWs(hub, w, r, roomId)
	})

	server.ListenServer(config.Server.Host, config.Server.Port, router)
}
