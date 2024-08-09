package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
)

func ToGetMnemonicAuthQueryResponse(mnemonic string) (*responses.GetMnemonicAuthQueryResponse, error) {
	response := responses.GetMnemonicAuthQueryResponse{
		Mnemonic: mnemonic,
	}

	return &response, nil
}
