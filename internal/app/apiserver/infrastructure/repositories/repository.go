package repositories

import "github.com/fromsi/example/internal/app/apiserver/domain/repositories"

type MutableRepository struct {
	PostRepository repositories.MutablePostRepository
	UserRepository repositories.MutableUserRepository
}

func NewMutableRepository(postRepository repositories.MutablePostRepository, userRepository repositories.MutableUserRepository) *MutableRepository {
	return &MutableRepository{
		PostRepository: postRepository,
		UserRepository: userRepository,
	}
}

type QueryRepository struct {
	PostRepository repositories.QueryPostRepository
	UserRepository repositories.QueryUserRepository
}

func NewQueryRepository(postRepository repositories.QueryPostRepository, userRepository repositories.QueryUserRepository) *QueryRepository {
	return &QueryRepository{
		PostRepository: postRepository,
		UserRepository: userRepository,
	}
}
