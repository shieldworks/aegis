/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package env

import (
	data "github.com/shieldworks/aegis/core/entity/data/v1"
	"os"
	"strconv"
	"strings"
	"time"
)

// SafeSecretBufferSize returns the buffer size for the Aegis Safe secret queue.
//
// The buffer size is determined by the environment variable
// AEGIS_SAFE_SECRET_BUFFER_SIZE.
//
// If the environment variable is not set, the default buffer size is 10.
// If the environment variable is set and can be parsed as an integer,
// it will be used as the buffer size.
// If the environment variable is set but cannot be parsed as an integer,
// the default buffer size is used.
func SafeSecretBufferSize() int {
	p := os.Getenv("AEGIS_SAFE_SECRET_BUFFER_SIZE")
	if p == "" {
		return 10
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 10
	}
	return l
}

// SafeK8sSecretBufferSize returns the buffer size for the Aegis Safe Kubernetes
// secret queue.
//
// The buffer size is determined by the environment variable
// AEGIS_SAFE_K8S_SECRET_BUFFER_SIZE.
//
// If the environment variable is not set, the default buffer size is 10.
// If the environment variable is set and can be parsed as an integer,
// it will be used as the buffer size.
// If the environment variable is set but cannot be parsed as an integer,
// the default buffer size is used.
func SafeK8sSecretBufferSize() int {
	p := os.Getenv("AEGIS_SAFE_K8S_SECRET_BUFFER_SIZE")
	if p == "" {
		return 10
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 10
	}
	return l
}

// SafeSecretDeleteBufferSize returns the buffer size for the Aegis Safe secret
// deletion queue.
//
// The buffer size is determined by the environment variable
// AEGIS_SAFE_SECRET_DELETE_BUFFER_SIZE.
//
// If the environment variable is not set, the default buffer size is 10.
// If the environment variable is set and can be parsed as an integer,
// it will be used as the buffer size.
// If the environment variable is set but cannot be parsed as an integer,
// the default buffer size is used.
func SafeSecretDeleteBufferSize() int {
	p := os.Getenv("AEGIS_SAFE_SECRET_DELETE_BUFFER_SIZE")
	if p == "" {
		return 10
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 10
	}
	return l
}

// SafeK8sSecretDeleteBufferSize returns the buffer size for the Aegis Safe
// Kubernetes secret deletion queue.
//
// The buffer size is determined by the environment variable
// AEGIS_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE.
//
// If the environment variable is not set, the default buffer size is 10.
// If the environment variable is set and can be parsed as an integer,
// it will be used as the buffer size.
// If the environment variable is set but cannot be parsed as an integer,
// the default buffer size is used.
func SafeK8sSecretDeleteBufferSize() int {
	p := os.Getenv("AEGIS_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE")
	if p == "" {
		return 10
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 10
	}
	return l
}

// SafeFipsCompliant returns a boolean indicating whether Aegis Safe should run in
// FIPS compliant mode. Note that this is not a guarantee that Aegis Safe will
// run in FIPS compliant mode, as it depends on the underlying base image.
// If you are using one of the official FIPS-complaint Aegis Docker images,
// then it will be FIPS-compliant. Check https://aegis.ist/docs/configuration/
// for more details.
func SafeFipsCompliant() bool {
	p := strings.ToLower(os.Getenv("AEGIS_SAFE_FIPS_COMPLIANT"))
	if p == "" {
		return false
	}

	return p == "true"
}

// SafeBackingStore returns the storage type for the data,
// as specified in the AEGIS_SAFE_BACKING_STORE environment variable.
// If the environment variable is not set, it defaults to "file".
// Any value that is not "file" will mean Aegis Safe will store
// its state in-memory
func SafeBackingStore() data.BackingStore {
	s := os.Getenv("AEGIS_SAFE_BACKING_STORE")
	if s == "" {
		return data.File
	}

	if data.BackingStore(s) == data.File {
		return data.File
	}

	return data.Memory
}

// SafeUseKubernetesSecrets returns a boolean indicating whether to create a
// plain text Kubernetes secret for the workloads registered. There are two
// things to note about this approach:
//
// 1. By design, and for security the original kubernetes `Secret` should exist,
// and it should be initiated to a default data as follows:
//
//	data:
//	  # '{}' (e30=) is a special placeholder to tell Safe that the Secret
//	  # is not initialized. DO NOT remove or change it.
//	  KEY_TXT: "e30="
//
// 2. This approach is LESS secure, and it is meant to be used for LEGACY
// systems where directly using the Safe Sidecar or Safe SDK are not feasible.
// It should be left as a last resort.
//
// If the environment variable is not set or its value is not "true",
// the function returns false. Otherwise, the function returns true.
func SafeUseKubernetesSecrets() bool {
	p := os.Getenv("AEGIS_SAFE_USE_KUBERNETES_SECRETS")
	if p == "" {
		return false
	}
	if strings.ToLower(p) == "true" {
		return true
	}
	return false
}

// SafeSecretBackupCount retrieves the number of backups to keep for Aegis Safe
// secrets. If the environment variable AEGIS_SAFE_SECRET_BACKUP_COUNT is not
// set or is not a valid integer, the default value of 3 will be returned.
func SafeSecretBackupCount() int {
	p := os.Getenv("AEGIS_SAFE_SECRET_BACKUP_COUNT")
	if p == "" {
		return 3
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 3
	}
	return l
}

// SafeManualKeyInput returns a boolean indicating whether to use manual
// cryptographic key input for Aegis Safe, instead of letting it bootstrap
// automatically. If the environment variable is not set or its value is
// not "true", the function returns false. Otherwise, the function returns true.
func SafeManualKeyInput() bool {
	p := os.Getenv("AEGIS_SAFE_MANUAL_KEY_INPUT")
	if p == "" {
		return false
	}
	if strings.ToLower(p) == "true" {
		return true
	}
	return false
}

// SafeDataPath returns the path to the safe data directory.
// The path is determined by the AEGIS_SAFE_DATA_PATH environment variable.
// If the environment variable is not set, the default path "/data" is returned.
func SafeDataPath() string {
	p := os.Getenv("AEGIS_SAFE_DATA_PATH")
	if p == "" {
		p = "/data"
	}
	return p
}

// SafeAgeKeyPath returns the path to the safe age key directory.
// The path is determined by the AEGIS_CRYPTO_KEY_PATH environment variable.
// If the environment variable is not set, the default path "/key/key.txt"
// is returned.
func SafeAgeKeyPath() string {
	p := os.Getenv("AEGIS_CRYPTO_KEY_PATH")
	if p == "" {
		p = "/key/key.txt"
	}
	return p
}

// SafeBootstrapTimeout returns the allowed time for Aegis Safe to wait
// before killing the pod to retrieve an SVID, in time.Duration.
// The interval is determined by the AEGIS_SAFE_BOOTSTRAP_TIMEOUT environment
// variable, with a default value of 30000 milliseconds if the variable is not
// set or if there is an error in parsing the value.
func SafeBootstrapTimeout() time.Duration {
	p := os.Getenv("AEGIS_SAFE_BOOTSTRAP_TIMEOUT")
	if p == "" {
		p = "30000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 30000 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}

// SafeAgeKeySecretName returns the name of the environment variable that holds
// the Aegis Safe age key. The value is retrieved using the
// "AEGIS_CRYPTO_KEY_NAME" environment variable. If this variable is
// not set or is empty, the default value "aegis-safe-age-key" is returned.
func SafeAgeKeySecretName() string {
	p := os.Getenv("AEGIS_CRYPTO_KEY_NAME")
	if p == "" {
		p = "aegis-safe-age-key"
	}
	return p
}

// SafeSecretNamePrefix returns the prefix to be used for the names of secrets that
// Aegis Safe stores, when it is configured to persist the secret in the Kubernetes
// cluster as Kubernetes `Secret` objects.
//
// The prefix is retrieved using the "AEGIS_SAFE_SECRET_NAME_PREFIX"
// environment variable. If this variable is not set or is empty, the default
// value "aegis-secret-" is returned.
func SafeSecretNamePrefix() string {
	p := os.Getenv("AEGIS_SAFE_SECRET_NAME_PREFIX")
	if p == "" {
		p = "aegis-secret-"
	}
	return p
}
