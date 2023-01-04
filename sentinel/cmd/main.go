package main

import (
	"context"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"log"
)

func main() {
	// TODO:
	//
	// 1. Fetch workload SVID + bundle from SPIRE
	// 2. Get SafeSpiffeId from environment.
	// 3. Send a GET request to Safe (safe can know who you are from your spiffeid)
	// 4. parse and safe the returned data.

	// TODO: get this from environment.
	const socketPath = "unix:///spire-agent-socket/agent.sock"

	ctx := context.Background()

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)),
	)

	if err != nil {
		log.Println("Unable to create X509 source")
	} else {
		svid, err := source.GetX509SVID()
		if err != nil {
			// 2023/01/03 19:37:58 svid.id spiffe://aegis.z2h.dev/ns/default/sa/default/n/aegis-workload-demo-559877fd7d-92rcn
			log.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! could not get svid")
		}
		log.Println("svid.id", svid.ID)

		log.Println("Everything is awesome!", source)
	}
}
