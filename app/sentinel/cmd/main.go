/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/shieldworks/aegis/app/sentinel/internal/safe"
	"os"
)

func main() {
	parser := argparse.NewParser("aegis", "Assigns secrets to workloads.")
	list := parser.Flag("l", "list", &argparse.Options{
		Required: false, Default: false, Help: "lists all registered workloads.",
	})
	useKubernetes := parser.Flag("k", "use-k8s", &argparse.Options{
		Required: false, Default: false,
		Help: "update an associated Kubernetes secret upon save. " +
			"Overrides AEGIS_SAFE_USE_KUBERNETES_SECRETS.",
	})
	namespace := parser.String("n", "namespace", &argparse.Options{
		Required: false, Default: "default",
		Help: "the namespace of the Kubernetes Secret to create.",
	})
	backingStore := parser.String("b", "store", &argparse.Options{
		Required: false,
		Help: "backing store type (file|memory|cluster). " +
			"Overrides AEGIS_SAFE_BACKING_STORE.",
	})
	workload := parser.String("w", "workload", &argparse.Options{
		Required: false,
		Help: "name of the workload (i.e. the '$name' segment of its " +
			"ClusterSPIFFEID ('spiffe://trustDomain/workload/$name/…')).",
	})
	secret := parser.String("s", "secret", &argparse.Options{
		Required: false,
		Help:     "the secret to store for the workload.",
	})
	template := parser.String("t", "template", &argparse.Options{
		Required: false,
		Help:     "the template used to transform the secret stored.",
	})
	format := parser.String("f", "format", &argparse.Options{
		Required: false,
		Help: "the format to display the secrets in." +
			" Has effect only when `-t` is provided. " +
			"Valid values: yaml, json, and none. Defaults to none.",
	})

	encrypt := parser.Flag("e", "encrypt", &argparse.Options{
		Required: false,
		Help: "returns an encrypted version of the secret if used with `-s`; " +
			"decrypts the secret before registering it to the workload if used " +
			"with `-s` and `-w`.",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if list != nil && *list == true {
		safe.Get()
		return
	}

	if (workload == nil || *workload == "") && (encrypt == nil || !*encrypt) {
		fmt.Println("Please provide a workload name.")
		fmt.Println("")
		fmt.Println("type `aegis -h` (without backticks) and press return for help.")
		fmt.Println("")
		return
	}

	if secret == nil || *secret == "" {
		fmt.Println("Please provide a secret.")
		fmt.Println("")
		fmt.Println("type `aegis -h` (without backticks) and press return for help.")
		fmt.Println("")
		return
	}

	if namespace == nil || *namespace == "" {
		*namespace = "default"
	}

	workloadP := ""
	if workload != nil {
		workloadP = *workload
	}

	secretP := ""
	if secret != nil {
		secretP = *secret
	}

	namespaceP := ""
	if namespace != nil {
		namespaceP = *namespace
	}

	backingStoreP := ""
	if backingStore != nil {
		backingStoreP = *backingStore
	}

	useK8sP := false
	if useKubernetes != nil {
		useK8sP = *useKubernetes
	}

	templateP := ""
	if template != nil {
		templateP = *template
	}

	formatP := ""
	if format != nil {
		formatP = *format
	}

	encryptP := false
	if encrypt != nil {
		encryptP = *encrypt
	}

	safe.Post(
		workloadP, secretP, namespaceP, backingStoreP, useK8sP,
		templateP, formatP, encryptP,
	)
}
