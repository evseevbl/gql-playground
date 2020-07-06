package graph

import (
	"sync"

	"github.com/evseevbl/posts/internal/app/posts/graph/model"
)

func NewResolver() *Resolver {
	return &Resolver{
		postStorage: make([]*model.Post, 0),
		postChan:    make(chan *model.Post, 1),
		postReaders: make(map[int64]chan *model.Post),
	}
}

type Resolver struct {
	mu sync.Mutex

	idCounter         int64
	subscriberCounter int64
	postStorage       []*model.Post
	postChan          chan *model.Post
	postReaders       map[int64]chan *model.Post
}
