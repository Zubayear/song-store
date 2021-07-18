package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Zubayear/song-store/entity"
	"github.com/Zubayear/song-store/graph/generated"
	"github.com/Zubayear/song-store/graph/model"
	"github.com/Zubayear/song-store/repo"
)

func (r *mutationResolver) CreateSong(ctx context.Context, input model.SongInput) (*model.Song, error) {
	song := entity.Song{
		SongName:     input.SongName,
		SongDuration: input.SongDuration,
		SongHits:     input.SongHits,
	}
	s, _ := repo.NewSongRepository().CreateSong(&song)
	return &model.Song{
		SongName:     s.SongName,
		SongDuration: s.SongDuration,
		SongHits:     s.SongHits,
	}, nil
}

func (r *mutationResolver) UpdateSong(ctx context.Context, input model.SongInput) (*model.Song, error) {
	// song := entity.Song{
	// 	SongName:     input.SongName,
	// 	SongDuration: input.SongDuration,
	// 	SongHits:     input.SongHits,
	// }
	// s, _ := repo.NewSongRepository().UpdateSong(&song, inp)
	panic("Hello")
}

func (r *mutationResolver) DeleteSong(ctx context.Context, id int) (*string, error) {
	s, _ := repo.NewSongRepository().DeleteSongPermanently(uint(id))
	return &s, nil
}

func (r *queryResolver) GetSong(ctx context.Context, id int) (*model.Song, error) {
	v, _ := repo.NewSongRepository().GetSong(uint(id))
	return &model.Song{
		ID:           int(v.ID),
		SongName:     v.SongName,
		SongDuration: v.SongDuration,
		SongHits:     v.SongHits,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
