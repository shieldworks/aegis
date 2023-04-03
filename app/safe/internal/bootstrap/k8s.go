/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package bootstrap

import (
	"context"
	"github.com/shieldworks/aegis/app/safe/internal/state"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/log"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func persistKeys(privateKey, publicKey string) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.FatalLn("Error creating client config: %v", err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.FatalLn("Error creating client: %v", err.Error())
	}

	data := make(map[string][]byte)
	keysCombined := privateKey + "\n" + publicKey
	data["KEY_TXT"] = ([]byte)(keysCombined)

	// Update the Secret in the cluster
	_, err = clientset.CoreV1().Secrets("aegis-system").Update(
		context.Background(),
		&v1.Secret{
			TypeMeta: metaV1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metaV1.ObjectMeta{
				Name:      env.SafeAgeKeySecretName(),
				Namespace: "aegis-system",
			},
			Data: data,
		},
		metaV1.UpdateOptions{
			TypeMeta: metaV1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
		},
	)

	if err != nil {
		log.FatalLn("error creating the secret", err.Error())
		return
	}

	log.InfoLn("Created the secret.")
	state.SetAgeKey(keysCombined)
	log.InfoLn("Registered the age key into memory.")
}
