package commands

import (
	"errors"
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
)

type CreatePostCommand struct {
	Text string
}

type CreatePostCommandHandler struct {
	MutableRepository *repositories.MutableRepository
}

func (handler CreatePostCommandHandler) Handle(command Command) error {
	commandImplementation, exists := command.(CreatePostCommand)

	if !exists {
		return errors.New("invalid command type")
	}

	post, err := entities.NewPost("", commandImplementation.Text, nil, nil, nil)

	if err != nil {
		return err
	}

	err = handler.MutableRepository.PostRepository.Create(post)

	if err != nil {
		return err
	}

	return nil
}

type UpdateByIdPostCommand struct {
	ID   string
	Text *string
}

type UpdatePostCommandHandler struct {
	MutableRepository *repositories.MutableRepository
}

func (handler UpdatePostCommandHandler) Handle(command Command) error {
	commandImplementation, exists := command.(UpdateByIdPostCommand)

	if !exists {
		return errors.New("invalid command type")
	}

	post, err := entities.NewPost(commandImplementation.ID, *commandImplementation.Text, nil, nil, nil)

	if err != nil {
		return err
	}

	err = handler.MutableRepository.PostRepository.UpdateById(commandImplementation.ID, post)

	if err != nil {
		return err
	}

	return nil
}

type DeletePostCommand struct {
	ID string
}

type DeletePostCommandHandler struct {
	MutableRepository *repositories.MutableRepository
}

func (handler DeletePostCommandHandler) Handle(command Command) error {
	commandImplementation, exists := command.(DeletePostCommand)

	if !exists {
		return errors.New("invalid command type")
	}

	err := handler.MutableRepository.PostRepository.DeleteById(commandImplementation.ID)

	if err != nil {
		return err
	}

	return nil
}

type RestorePostCommand struct {
	ID string
}

type RestorePostCommandHandler struct {
	MutableRepository *repositories.MutableRepository
}

func (handler RestorePostCommandHandler) Handle(command Command) error {
	commandImplementation, exists := command.(RestorePostCommand)

	if !exists {
		return errors.New("invalid command type")
	}

	err := handler.MutableRepository.PostRepository.RestoreById(commandImplementation.ID)

	if err != nil {
		return err
	}

	return nil
}
