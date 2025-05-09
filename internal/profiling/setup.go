package profiling

import (
	"os"

	"github.com/grafana/pyroscope-go"
)

func SetupPyroscope() error {
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "dinheiros-api",

		// replace this with the address of pyroscope server
		ServerAddress: os.Getenv("GRAFANA_PYROSCOPE"),

		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
	})

	return err
}
