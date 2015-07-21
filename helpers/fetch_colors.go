package helpers

import (
	"github.com/aripka-pivotal/lattice-app/structs"
)

func FetchColors()[]structs.RGB {
	colors := make([]structs.RGB, 333)
        colors[0] = structs.RGB{ Red: `75`, Green: `0`, Blue: `130`}
        colors[1] = structs.RGB{ Red: `75`, Green: `12`, Blue: `143`}
        colors[2] = structs.RGB{ Red: `111`, Green: `45`, Blue: `180`}
	return colors
}
