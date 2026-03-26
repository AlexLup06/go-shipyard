package render

import (
	"net/http"

	"github.com/a-h/templ"
)

type Renderer func(w http.ResponseWriter, r *http.Request, status int, c templ.Component) error

type Assets map[string]string

func (a Assets) Static(name string) string {
	if a == nil {
		return "/static/" + name
	}
	if v, ok := a[name]; ok {
		return "/static/" + v
	}
	return "/static/" + name
}
