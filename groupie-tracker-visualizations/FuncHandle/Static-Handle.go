package Handle

import "net/http"

func StaticHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r,"./static/Style.css")
}