package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
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
	r.postStorage = append(r.postStorage, post)
	return post, nil
}

func (r *queryResolver) GetLastPosts(ctx context.Context, cnt *int) ([]*model.Post, error) {
	if *cnt <= 0 {
		return nil, errors.New("bad arg cnt")
	}
	if len(r.postStorage) <= *cnt {
		return r.postStorage, nil
	}
	sz := len(r.postStorage)
	return r.postStorage[sz-*cnt : sz-1], nil
}

func (r *subscriptionResolver) PostCreated(ctx context.Context) (<-chan *model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	events := make(chan *model.Post, 1)
	r.postReaders = append(r.postReaders, events)
	go func() {
		<-ctx.Done()
	}()
	return events, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
