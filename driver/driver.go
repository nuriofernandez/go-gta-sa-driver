package driver

import (
	"github.com/xXNurioXx/go-gta-sa-memory-reader/gtaMemoryAccessor"
	"github.com/xXNurioXx/go-gta-sa-memory-reader/gtaSa"
)

type Driver struct {
	GtaSa  gtaSa.GtaSa
	gtaApi gtaMemoryAccessor.GtaApi
}

func New() Driver {
	gtaSa := gtaSa.GtaSa{}
	gtaSa.Hook()
	return Driver{
		gtaSa,
		gtaSa.GtaApi,
	}
}
