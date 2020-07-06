package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/evseevbl/posts/internal/app/posts/graph/generated"
	"github.com/evseevbl/posts/internal/app/posts/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, title string, description string) (*model.Post, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.idCounter++
	post := &model.Post{
		Title:       &title,
		Description: &description,
		CreatedAt:   time.Now(),
	}
	cnt := int(r.idCounter)
	post.ID = &cnt
	r.postStorage[r.idCounter] = post
	return post, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{Resolver: r}
}

type mutationResolver struct {
	*Resolver
}
