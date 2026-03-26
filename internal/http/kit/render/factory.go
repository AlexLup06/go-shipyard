package render

import (
	"__MODULE__/internal/http/kit/httpctx"
	"net/http"

	"github.com/a-h/templ"
)

func New(a Assets) Renderer {
	return func(w http.ResponseWriter, r *http.Request, status int, component templ.Component) error {
		// Make assets available to templ via context:
		r = r.WithContext(httpctx.WithAssets(r.Context(), a))

		// (Optional but good) Ensure HTML content type
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		}

		w.WriteHeader(status)
		return component.Render(r.Context(), w)
	}
}
