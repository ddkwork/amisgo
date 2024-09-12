package comp

import (
	js "encoding/json"
	"fmt"
	"io"
	"net/http"
)

var innerApiID = -1

func getInnerApiID() int {
	innerApiID++
	return innerApiID
}

func getRoute() string {
	return fmt.Sprintf("/__amisgo__%d", getInnerApiID())
}

func serveApi(action func(map[string]any)) string {
	route := getRoute()
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		input, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		m := map[string]any{}
		err = js.Unmarshal(input, &m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		action(m)
	})
	return route
}

func serveData(getter func() any) string {
	route := getRoute()
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		data, err := js.Marshal(getter())
		if err != nil {
			resp := Response{Status: 1, Msg: err.Error()}
			w.Write(resp.Json())
			return
		}
		w.Write(data)
	})
	return route
}
