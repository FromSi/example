package queries

import (
	"errors"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/mappers"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/fromsi/example/internal/pkg/tools"
)

type GetMnemonicAuthQuery struct{}

type GetMnemonicAuthQueryHandler struct {
	QueryRepository *repositories.QueryRepository
}

func (handler GetMnemonicAuthQueryHandler) Handle(query Query) (any, error) {
	_, exists := query.(GetMnemonicAuthQuery)

	if !exists {
		return nil, errors.New("invalid query type")
	}

	addressBTC := tools.NewAddressBTC()

	response, err := mappers.ToGetMnemonicAuthQueryResponse(addressBTC.GetMnemonic())

	if err != nil {
		return nil, err
	}

	return response, nil
}
