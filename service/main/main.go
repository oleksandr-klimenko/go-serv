package main

import (
    "go-serv/core/request"
    "go-serv/core/response"
    "log"
)

func SumEndpoint(req request.RequestReader, resp response.ResponseWriter, context request.RequestContextReader) {

    body := req.GetBody().(map[string]int)

    resp.Write(&response.Response{
        Body: map[string] int {
            "sum": Sum(body["a"], body["b"]),
        },
        Status: 200,
    })

}

func Sum(a int, b int) int {
    return a + b
}

type EchoWriter struct {}

func (w EchoWriter) Write(resp *response.Response) {
    log.Println(resp.Body)
}

func main () {

    req := request.Request{}
    req.Init(request.RequestParams{
        Body: map[string] int {
            "a": 2,
            "b": 4,
        },
        Resource: "/test",
        Method: request.GET,
    })

    context := request.RequestContext{}
    context.Init(request.RequestContextParams{
        ContentType: "application/json",
        AuthorizationString: "",
        RequesterId: "asd",
        RequestId: "asdd",
        Params: map[string] interface {} {
            "some": "other",
        },
    })

    writer := EchoWriter{}

    SumEndpoint(&req, &writer, &context)

}


