package server

import (
	api "../../AIO_API"
)

var RoutesMap = map[string] interface{} {
	"/Api/CreateUser" : api.CreateUser,
	"/Api/GetUser" : api.GetUser,

}