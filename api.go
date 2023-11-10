package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"price-fetcher/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONAPIService struct {
	listerAddr string
	svc        PriceFetcher
}

func NewJSONService(listenAddr string, svc PriceFetcher) *JSONAPIService {
	return &JSONAPIService{
		listerAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIService) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))

	http.ListenAndServe(s.listerAddr, nil)
}

func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONAPIService) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)

	if err != nil {
		return err
	}

	priceResp := types.PriceResp{
		Price:  price,
		Ticker: ticker,
	}

	return WriteJSON(w, http.StatusOK, &priceResp)

}

func WriteJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
