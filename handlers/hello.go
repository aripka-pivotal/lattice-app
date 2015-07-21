package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aripka-pivotal/lattice-app/helpers"
)

type Hello struct {
	Time time.Time
}

func (p *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index, _ := helpers.FetchIndex()
	appName := "APP 1"
	colors := helpers.FetchColors()

	styledColoredTemplate.Execute(w, ColoredBody{Body: fmt.Sprintf(`
<div class="hello">
	ALattice - %s
</div>

<div class="my-index">My Index Is</div>

<div class="index">%d</div>
<div class="mid-color">Uptime: %s</div>
<div class="bottom-color"></div>
    `, appName, index, time.Since(p.Time)), Colors:colors})
}
