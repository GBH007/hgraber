package webServer

import (
	"app/service/webServer/base"
	"app/service/webServer/rendering"
	"net/http"
)

func (ws *WebServer) routeTitleList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := struct {
			Count  int `json:"count"`
			Offset int `json:"offset,omitempty"`
		}{}

		ctx := r.Context()

		err := base.ParseJSON(r, &request)
		if err != nil {
			base.WriteJSON(ctx, w, http.StatusBadRequest, err)
			return
		}

		data := ws.Storage.GetTitles(ctx, request.Offset, request.Count)

		base.WriteJSON(ctx, w, http.StatusOK, rendering.TitlesFromStorage(data))
	})
}