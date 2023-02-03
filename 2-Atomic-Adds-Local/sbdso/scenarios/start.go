package scenarios

import (
	"backend/config"
	"backend/server"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func Start(servers []config.Node, scenario string, zctx *zmq.Context, bdso_networks map[string][]config.Node) {
	// all servers act normal
	if scenario == "NORMAL" {
		for i := 0; i < config.N; i++ {
			s := server.CreateServer(servers[i], servers, zctx, bdso_networks)
			go Normal_listener_task(s)
		}
	}
	// f mutes
	if scenario == "MUTES" {
		// f mute
		for i := 0; i < config.F; i++ {
			go Mute_listener_task(servers[i], servers, zctx)
			fmt.Println("Mute: ", servers[i])
		}
		// n-f normal
		for i := config.F; i < config.N; i++ {
			s := server.CreateServer(servers[i], servers, zctx, bdso_networks)
			go Normal_listener_task(s)
			fmt.Println("Normal: ", servers[i])
		}
	}
	// f/2 act correctly, the rest act maliciously
	if scenario == "HALF&HALF" {

	}
}
