package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Slobodskoy/calc_go/internal/pkg/calc"
)

type CalcHandler struct{}

func (*CalcHandler) Calc(w http.ResponseWriter, r *http.Request) {
	request, err := parseCalcRequest(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
		return
	}
	result, err := calc.Calc(request.Expression)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
		return
	}
	fmt.Fprintf(w, "{\"result\": \"%f\"}", result)
}

func parseCalcRequest(r *http.Request) (*CalcRequest, error) {
	var request CalcRequest
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&request); err != nil {
		return nil, fmt.Errorf("decode err: %w", err)
	}
	return &request, nil
}
