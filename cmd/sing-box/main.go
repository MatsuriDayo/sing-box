//go:build !generate

package boxmain

import "github.com/sagernet/sing-box/log"

func Main() {
	if err := mainCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
