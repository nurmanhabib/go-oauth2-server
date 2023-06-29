package passport

import "github.com/go-oauth2/oauth2/v4/server"

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) Build() *server.Server {
	d.builder.Reset()
	d.builder.SetManager()
	d.builder.SetServer()

	return d.builder.GetServer()
}
