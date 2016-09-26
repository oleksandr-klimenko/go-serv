package request

import "errors"

type Method int

const (
    GET Method = 1 + iota
    POST
    UPDATE
    DELETE
    OPTIONS
)

type RequestReader interface {
    GetMethod() Method
    GetResource() string
    GetBody() interface{}
}

type RequestParams struct {
    Method Method `json:"method"`
    Resource string `json:"resourse"`
    Body interface{} `json:"body"`
}

type Request struct {
    method Method
    resource string
    body interface{}
    isInitialized bool
}

func (request *Request) Init(params RequestParams) error {

    if request.isInitialized {
        return errors.New("Request is initialized")
    }

    request.body = params.Body
    request.method = params.Method
    request.resource = params.Resource

    return nil
}

func (request Request) GetMethod() Method {
    return request.method
}

func (request Request) GetResource() string {
    return request.resource
}

func (request Request) GetBody() interface {} {
    return request.body
}