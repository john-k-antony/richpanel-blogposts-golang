package responder

import (
	"net/http"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/john-k-antony/richpanel-blogposts-golang/models/apimodels"
)

// Responder implements middleware.Responder and allows us to
// build responses easily.
type Responder struct {
	req    *http.Request
	status int
	body   interface{}
	header map[string]string
}

// New creates a new responder.
func New(req *http.Request) (resp *Responder) {
	resp = &Responder{req: req}
	return resp
}

type customResponder func(http.ResponseWriter, runtime.Producer)

func (c customResponder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	c(w, p)
}

// NewCustomResponder creates a new custome HTTP responder.
func NewCustomResponder(r *http.Request, h http.Handler) middleware.Responder {
	return customResponder(func(w http.ResponseWriter, _ runtime.Producer) {
		h.ServeHTTP(w, r)
	})
}

// WriteResponse writes HTTP response.
func (resp *Responder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	if resp.status == 0 {
		resp.status = 200
	}

	// Put other logic before writing.

	resp.write(rw, producer)
}

// write writes response:
//  1. Set header fields.
//  2. Write header with status.
//  3. Write body.
func (resp *Responder) write(rw http.ResponseWriter, producer runtime.Producer) {
	// If response body is nil, remove Content-Type header field only.
	if resp.body == nil {
		rw.Header().Del(runtime.HeaderContentType)
	}

	if resp.header != nil {
		for key, value := range resp.header {
			rw.Header().Set(key, value)
		}
	}
	rw.WriteHeader(resp.status)

	if resp.body != nil {
		resp.fixNull()

		var err error
		if b, ok := resp.body.([]byte); ok {
			_, err = rw.Write(b)
		} else {
			err = producer.Produce(rw, resp.body)
		}

		if err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// fixNull makes sure that we don't marshal to "null".
func (resp *Responder) fixNull() {
	v := reflect.ValueOf(resp.body)
	if v.Kind() == reflect.Slice && v.Len() == 0 {
		resp.body = []int{} // Any array which is marshalled to []
	}
}

// Status sets HTTP status code.
func (resp *Responder) Status(status int) *Responder {
	resp.status = status
	return resp
}

// Body sets response body to respond with.
func (resp *Responder) Body(body interface{}) *Responder {
	resp.body = body
	return resp
}

// Error prepares response body.
func (resp *Responder) Error(httpStatus int, code int64, err error) *Responder {
	errMsg := err.Error()
	body := apimodels.ResponseError{Message: &errMsg}
	resp.status = httpStatus
	resp.body = body
	return resp
}

// Header adds a new header field e.g. "Content-Type: application/json"
func (resp *Responder) Header(key string, value string) *Responder {
	if resp.header == nil {
		resp.header = make(map[string]string)
	}
	resp.header[key] = value
	return resp
}
