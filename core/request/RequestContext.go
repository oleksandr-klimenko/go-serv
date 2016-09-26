package request

import (
    "errors"
)

type RequestContextReader interface {
    GetContentType() string
    GetAuthorizationString() string
    GetRequesterId() string
    GetRequestId() string
    GetParams() map [string] interface{}
}

type RequestContextParams struct  {
    ContentType string `json:"contentType"`
    AuthorizationString string `json:"authorization"`
    RequesterId string `json:"requsterId"`
    RequestId string `json:"requestId"`
    Params map [string] interface{}
}

type RequestContext struct  {
    contentType string
    authorizationString string
    requesterId string
    requestId string
    isInitialized bool
    params map [string] interface{}
}

func (context *RequestContext) Init(params RequestContextParams) error {

    if !context.isInitialized {
        return errors.New("Request context is initialized")
    }

    context.contentType = params.ContentType
    context.authorizationString = params.AuthorizationString
    context.requesterId = params.RequesterId
    context.requestId = params.RequestId
    context.params = params.Params

    context.isInitialized = true

    return nil
}

func (context RequestContext) GetContentType() string {
    return context.contentType
}

func (context RequestContext) GetAuthorizationString() string {
    return context.authorizationString
}

func (context RequestContext) GetRequesterId() string {
    return context.requesterId
}

func (context RequestContext) GetRequestId() string {
    return context.requestId
}

func (context RequestContext) GetParams() map [string] interface{} {
    return context.params
}