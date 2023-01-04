package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"log"
	"os"
)

// TODO: get this from environment.
const (
	socketPath = "unix:///spire-agent-socket/agent.sock"
	serverUrl  = "https://aegis-safe.aegis-system.svc.cluster.local:8443/"
)

func main() {
	parser := argparse.NewParser(
		"aegis",
		"Assigns secrets to workloads.",
	)

	workload := parser.String(
		"w", "workload",
		&argparse.Options{
			Required: true,
			Help:     "name of the workload (i.e. the '$name' segment of its ClusterSPIFFEID ('spiffe://trustDomain/workload/$name/…'))",
		},
	)

	secret := parser.String(
		"s", "secret",
		&argparse.Options{
			Required: true,
			Help:     "the secret to store for the workload",
		},
	)

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	log.Println("workload", workload, "secret", secret)

	//// Finally print the collected string
	//fmt.Println("workload", workload, "secret", secret)
	//
	//log.Println("Welcome to sentinel")
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	//log.Println("will create svid")
	//
	//source, err := workloadapi.NewX509Source(
	//	ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)),
	//)
	//
	//if err != nil {
	//	log.Println("Unable to create X509 source")
	//} else {
	//	svid, err := source.GetX509SVID()
	//	if err != nil {
	//		// 2023/01/03 19:37:58 svid.id spiffe://aegis.z2h.dev/ns/default/sa/default/n/aegis-workload-demo-559877fd7d-92rcn
	//		log.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! could not get svid")
	//	}
	//	log.Println("svid.id", svid.ID)
	//
	//	log.Println("Everything is awesome!", source)
	//}
	//defer func(source *workloadapi.X509Source) {
	//	err := source.Close()
	//	if err != nil {
	//		// TODO: handle me
	//	}
	//}(source)
	//
	//// Allowed SPIFFE ID
	//// serverID := spiffeid.RequireFromString("spiffe://example.org/server")
	//// spiffe://aegis.z2h.dev/ns/{{ .PodMeta.Namespace }}/sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}
	//
	//authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
	//	ids := id.String()
	//
	//	if strings.HasPrefix(ids, "spiffe://aegis.z2h.dev/ns/aegis-system/sa/aegis-safe/n/") {
	//		return nil
	//	}
	//
	//	return errors.New("I don’t know you, and it’s crazy")
	//})
	//
	//// Create a `tls.Config` to allow mTLS connections, and verify that presented certificate has SPIFFE ID `spiffe://example.org/server`
	//tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	//client := &http.Client{
	//	Transport: &http.Transport{
	//		TLSClientConfig: tlsConfig,
	//	},
	//}
	//
	//p, err := url.JoinPath(serverUrl, "/v1/secret")
	//if err != nil {
	//	// TODO: handle this
	//	return
	//}
	//
	//sr := v1.SecretUpsertRequest{
	//	WorkloadId: "aegis-workload-demo",
	//	Value:      `{"u": "root", "p": "toppyTopSecret", "realm": "narnia"}`,
	//}
	//
	//md, err := json.Marshal(sr)
	//if err != nil {
	//	// TODO: handle me
	//	log.Println("handle me")
	//	return
	//}
	//
	//r, err := client.Post(p, "application/json", bytes.NewBuffer(md))
	//if err != nil {
	//	log.Fatalf("Error connecting to %q: %v", serverUrl, err)
	//}
	//
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//		// TODO: handle me.
	//		log.Println("handle me")
	//	}
	//}(r.Body)
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	log.Fatalf("Unable to read body: %v", err)
	//}
	//
	//p, err = url.JoinPath(serverUrl, "/v1/fetch")
	//if err != nil {
	//	// TODO: handle this
	//	return
	//}
	//
	//// This is only for testing and will normally be rejected due to svid mismatch.
	//
	//sfr := v1.SecretFetchRequest{}
	//
	//md, err = json.Marshal(sfr)
	//if err != nil {
	//	// TODO: handle me
	//	log.Println("handle me")
	//	return
	//}
	//
	//r, err = client.Post(p, "application/json", bytes.NewBuffer(md))
	//if err != nil {
	//	log.Fatalf("Error connecting to %q: %v", serverUrl, err)
	//}
	//
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//		// TODO: handle me.
	//		log.Println("handle me")
	//	}
	//}(r.Body)
	//body, err = io.ReadAll(r.Body)
	//if err != nil {
	//	log.Fatalf("Unable to read body: %v", err)
	//}
	//
	//log.Printf("%s", body)
}
