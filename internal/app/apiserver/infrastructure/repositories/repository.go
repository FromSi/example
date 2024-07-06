package repositories

import "github.com/fromsi/example/internal/app/apiserver/domain/repositories"

type MutableRepository struct {
	PostRepository repositories.MutablePostRepository
}

func NewMutableRepository(postRepository repositories.MutablePostRepository) *MutableRepository {
	return &MutableRepository{
		PostRepository: postRepository,
	}
}

type QueryRepository struct {
	PostRepository repositories.QueryPostRepository
}

func NewQueryRepository(postRepository repositories.QueryPostRepository) *QueryRepository {
	return &QueryRepository{
		PostRepository: postRepository,
	}
}
