package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"ztravel/pkg/clock"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type Signer interface {
	Sign(ctx context.Context, content []byte) ([]byte, error)
	GetKeyId() string
}

type HttpResponseSigner struct {
	Signer Signer
	Clock  clock.Clockwork
}

func (s *HttpResponseSigner) WithSignedResponses(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := NewHttpResponseInterceptor(w)
		h.ServeHTTP(ww, r)
		body, _ := io.ReadAll(ww.Body)
		s.Sign(w, r, body)
	})
}

func (s *HttpResponseSigner) Sign(w http.ResponseWriter, r *http.Request, body []byte) {
	now := s.Clock().UTC().Format(http.TimeFormat)

	requestId := middleware.GetReqID(r.Context())
	if requestId == "" {
		log.Warn().Msg("Failed to generate signature; missing request ID")
		requestId = "request-id-not-set"
	}

	var parts []string
	parts = append(parts,
		requestId,
		r.Method,
		r.URL.Path,
		now,
		string(body),
	)

	signed, err := s.Signer.Sign(r.Context(), []byte(strings.Join(parts, ":")))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate signature")
	}

	encoded := base64.StdEncoding.EncodeToString(signed)
	signature := fmt.Sprintf("keyId=%s,signature=%s", s.Signer.GetKeyId(), encoded)
	w.Header().Add("Signature", signature)
	w.Header().Add("Signature-Date", now)
	w.Write(body)
}

func NewHttpResponseSigner(signer Signer, clock clock.Clockwork) *HttpResponseSigner {
	return &HttpResponseSigner{
		Signer: signer,
		Clock:  clock,
	}
}

type HttpResponseInterceptor struct {
	http.ResponseWriter
	Body *bytes.Buffer
}

func (i *HttpResponseInterceptor) Write(buf []byte) (int, error) {
	if i.Body != nil {
		i.Body.Write(buf)
	}
	return len(buf), nil
}

func NewHttpResponseInterceptor(w http.ResponseWriter) *HttpResponseInterceptor {
	return &HttpResponseInterceptor{
		ResponseWriter: w,
		Body:           new(bytes.Buffer),
	}
}
