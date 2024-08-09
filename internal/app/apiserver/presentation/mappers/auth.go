package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	presentationresponses "github.com/fromsi/example/internal/app/apiserver/presentation/responses"
)

func ToGinShowLoginAuthResponse(auth *responses.GetMnemonicAuthQueryResponse) (*presentationresponses.Response, error) {
	if auth == nil {
		return nil, nil
	}

	return &presentationresponses.Response{
		Data: presentationresponses.ShowLoginAuthResponse{
			Mnemonic: auth.Mnemonic,
		},
	}, nil
}
