package boot

import (
	bootmanager "github.com/thinhlh/go-web-101/internal/core/boot_manager"
	"github.com/thinhlh/go-web-101/internal/server"
)

func Init() {
	bootmanager.Bootstrap(server.New())
}
