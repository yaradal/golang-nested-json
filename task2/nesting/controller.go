package nesting

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func NewController(s Service) http.Handler {
	t := &controller{svc: s}

	router := chi.NewRouter()
	router.Post("/json", t.CreateLeavesFromLoop)
	return router
}

type controller struct {
	svc Service
}

func (t *controller) CreateLeavesFromLoop(w http.ResponseWriter, r *http.Request) {
	// Since it's not specified in the instructions, we take the arguments as a comma separated list.
	argsStr := r.URL.Query().Get("nesting_levels")
	args := strings.Split(argsStr, ",")
	// We remove the first "useless" argument in case the param is empty
	if len(args) == 1 && args[0] == "" {
		args = []string{}
	}

	items := []map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		http.Error(w, fmt.Sprintf("error while decoding the json: %s", err), http.StatusBadRequest)
		return
	}

	outputMap, err := t.svc.CreateLeavesFromLoop(args, items)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while nesting the input file: %s", err), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(outputMap); err != nil {
		http.Error(w, fmt.Sprintf("error while encoding the response: %s", err), http.StatusInternalServerError)
		return
	}
}
