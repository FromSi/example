package cqrs

import (
	"errors"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/commands"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/fromsi/example/internal/pkg/cqrs"
)

type DefaultCommandCQRS struct {
	MutableRepository *repositories.MutableRepository
	QueryRepository   *repositories.QueryRepository
}

func NewCommandCQRS(mutableRepository *repositories.MutableRepository, queryRepository *repositories.QueryRepository) cqrs.CommandCQRS {
	return &DefaultCommandCQRS{
		MutableRepository: mutableRepository,
		QueryRepository:   queryRepository,
	}
}

func (cqrs *DefaultCommandCQRS) Dispatch(command cqrs.Command) error {
	commandHandler, err := getCommandHandler(command, cqrs)

	if err != nil {
		return err
	}

	return commandHandler.Handle(command)
}

func getCommandHandler(command cqrs.Command, cqrs *DefaultCommandCQRS) (cqrs.CommandHandler, error) {
	switch command.(type) {
	case commands.CreatePostCommand:
		return &commands.CreatePostCommandHandler{MutableRepository: cqrs.MutableRepository}, nil
	case commands.UpdateByIdPostCommand:
		return &commands.UpdatePostCommandHandler{MutableRepository: cqrs.MutableRepository}, nil
	case commands.DeletePostCommand:
		return &commands.DeletePostCommandHandler{MutableRepository: cqrs.MutableRepository}, nil
	case commands.RestorePostCommand:
		return &commands.RestorePostCommandHandler{MutableRepository: cqrs.MutableRepository}, nil
	}

	return nil, errors.New("command handler not found")
}
