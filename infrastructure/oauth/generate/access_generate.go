package generate

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
)

type AccessGenerate struct{}

func (a *AccessGenerate) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error) {
	// TODO implement me
	panic("implement me")
}
