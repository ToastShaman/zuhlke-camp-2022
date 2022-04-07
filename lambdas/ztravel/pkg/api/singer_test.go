package api

import (
	"context"
	"crypto/sha1"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"ztravel/pkg/aws"
	"ztravel/pkg/clock"

	"github.com/stretchr/testify/assert"
)

type HashSigner struct{}

func (s *HashSigner) Sign(ctx context.Context, content []byte) ([]byte, error) {
	h := sha1.New()
	h.Write(content)
	bs := h.Sum(nil)
	return bs, nil
}

func (s *HashSigner) GetKeyId() string {
	return "my-key-id"
}

func TestHttpResponseSigner(t *testing.T) {
	signer := NewHttpResponseSigner(&HashSigner{}, clock.EPOCH)

	r := NewDefaultRouter()
	r.Get("/foobar", signer.WithSignedResponses(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}))

	lc := &aws.LogContext{
		ApiRequestId:    "RT48qPKBd2",
		LambdaRequestId: "RT48qPKBd2",
	}

	req := httptest.NewRequest("GET", "/foobar", nil)
	req = req.WithContext(lc.NewContext(req.Context()))

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	res := w.Result()

	payload, _ := ioutil.ReadAll(res.Body)
	signature := res.Header.Get("Signature")
	date := res.Header.Get("Signature-Date")

	assert.Equal(t, "Hello", string(payload), "they should be equal")
	assert.Equal(t, "keyId=my-key-id,signature=DcnNs2iWwE4Uj/T+bgMRpr+X+GA=", signature, "they should be equal")
	assert.Equal(t, "Thu, 01 Jan 1970 00:00:00 GMT", date, "they should be equal")
}
