//go:build !windows && !macos
// +build !windows,!macos

package platform

import (
	"github.com/lavren1974/rssfeeds/src/server"
)

func Start(s *server.Server) {
	s.Start()
}
