package server

import (
	api "../AIO_API"
)

/*
Карта доступных марщрутов
 */
var RoutesMap = map[string] interface{} {
	"/Api/CreateUser" : api.CreateUser,
	"/Api/GetUser" : api.GetUser,

}