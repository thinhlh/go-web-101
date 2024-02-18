package boot

import (
	"github.com/thinhlh/go-web-101/internal/core"
	"github.com/thinhlh/go-web-101/internal/server"
)

func Init() {
	core.InitDatabase()
	server.New()
}
