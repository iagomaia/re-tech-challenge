package adapters

import (
	"encoding/json"
	contractDomain "github.com/iagomaia/re-tech-challenge/internal/presentation/contracts"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
	presentation "github.com/iagomaia/re-tech-challenge/internal/presentation/protocols"
)

func AdaptRoute[T interface{}](controller presentation.IHandler, params *[]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := ExtractRequest(r, new(T), params)
		if err != nil {
			handleError(err, w, r)
			return
		}

		response, err := controller.Handle(request)
		if err != nil {
			handleError(err, w, r)
			return
		}
		w.WriteHeader(response.Status)
		for k, v := range response.Headers {
			w.Header().Set(k, v)
		}
		json.NewEncoder(w).Encode(response.Body)
	}
}

func ExtractRequest(r *http.Request, reqBody any, paramNames *[]string) (*presentation.HttpRequest, error) {
	if reqBody != nil && r.Body != nil {
		err := json.NewDecoder(r.Body).Decode(reqBody)
		if err != nil && err != io.EOF {
			cErr := utils.CustomError{
				Status:        http.StatusInternalServerError,
				Message:       "Error parsing request",
				OriginalError: err,
			}
			return nil, cErr
		}
	}

	headers := map[string][]string{}
	for k, v := range r.Header {
		headers[k] = v
	}

	params := map[string]string{}
	if paramNames != nil {
		for _, p := range *paramNames {
			params[p] = chi.URLParam(r, p)
		}
	}

	query := map[string][]string{}
	for k, v := range r.URL.Query() {
		query[k] = v
	}

	req := &presentation.HttpRequest{
		Headers: headers,
		Body:    reqBody,
		Params:  params,
		Query:   query,
		Ctx:     r.Context(),
	}

	return req, nil
}

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	entry := httplog.LogEntry(r.Context())
	cid := r.Context().Value(utils.CidContextKey)
	entry.Err(err).Msgf("An error happened at request %v: %v", cid, err)
	cErr := utils.GetCustomError(err)
	body := &contractDomain.HttpError{
		Message: cErr.Message,
	}
	w.WriteHeader(cErr.Status)
	json.NewEncoder(w).Encode(body)
}
