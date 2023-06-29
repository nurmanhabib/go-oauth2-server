package passport

import "github.com/go-oauth2/oauth2/v4/server"

type Builder interface {
	Reset()
	SetManager()
	SetServer()
	GetServer() *server.Server
}
