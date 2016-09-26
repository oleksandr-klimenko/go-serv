package service

import (
    "go-serv/core/request"
    "go-serv/core/response"
)

type Endpoint func(request request.RequestReader, response response.ResponseWriter, context request.RequestContextReader)