// +build !release

package client

var (
	rootCrtPaths = []string{"assets/client/tls/ngrokroot.crt",
		"assets/client/tls/letsencrypt.crt",
		"assets/client/tls/startcom.crt",
		"assets/client/tls/snakeoilca.crt",
	}
)

func useInsecureSkipVerify() bool {
	return true
}
