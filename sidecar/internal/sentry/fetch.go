/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package sentry

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

func saveData(data string) {
	path := "/opt/aegis/secrets.json"

	// fmt.Println("path:", path)

	f, err := os.Create(path)
	if err != nil {
		// TODO: handle me.
		panic("poop!")
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString(data)
	if err != nil {
		// TODO: handle me
		panic("poop!")
	}

	// fmt.Println("wrote", n, "bytes.")

	err = w.Flush()
	if err != nil {
		// TODO: handle
		panic("poop")
	}
}

func fetchSecrets() {
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

	log.Println("WILL FETCH SECRETS!!!")

	//if !state.Bootstrapped() {
	//	return
	//}
	//
	//id := state.Id()
	//secret := state.Secret()
	//
	//fmt.Println(state.Id(), state.Secret(), state.SafeApiRoot())
	//
	//res, err := newSafeFetchEndpoint()(
	//	context.Background(),
	//	reqres.SecretFetchRequest{
	//		WorkloadId:     id,
	//		WorkloadSecret: secret,
	//	})
	//if err != nil {
	//	// TODO: handle me
	//	panic("handle me")
	//}
	//
	//sfr, ok := res.(reqres.SecretFetchResponse)
	//if !ok {
	//	// TODO: handle me
	//	panic("handle me!")
	//}
	//
	//data := sfr.Data
	//
	//// TODO: save data to /opt/aegis/secrets.json
	//// TODO: make the filename configurable.
	//// fmt.Println("data: '", data, "'")
	//
	// saveData(data)
}
