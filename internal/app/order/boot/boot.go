package boot

import (
	"github.com/thinhlh/go-web-101/internal/app/order/server"
	bootmanager "github.com/thinhlh/go-web-101/internal/core/boot_manager"
)

func Init() {
	bootmanager.Bootstrap(server.New())
}
