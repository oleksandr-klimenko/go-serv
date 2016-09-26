package response

type Response struct {
    Body interface{} `json:"body"`
    Status int `json:"status"`
}

type ResponseContext map[string] interface{}

type ResponseWriter interface {
    Write(response *Response)
}

type ResponseContextWriter interface {
    Write(response *Response, context *ResponseContext)
}

