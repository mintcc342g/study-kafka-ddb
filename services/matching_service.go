package services

import (
	"context"
	"study-kafka-ddb/controllers/dtos"
	"study-kafka-ddb/domains"
	"study-kafka-ddb/domains/interfaces"
	"study-kafka-ddb/utils/deftype"
)

type MatchingService struct {
	userRepo  interfaces.UserRepository
	bandRepo  interfaces.BandRepository
	postRepo  interfaces.PostRepository
	eventRepo interfaces.EventRepository
}

const (
	openPositionTopic = "study-app.matching.open-position.event.v1"
	seekPositionTopic = "study-app.matching.seek-position.event.v1"
)

func (r *MatchingService) OpenPosition(ctx context.Context, req *dtos.OpenPositionReq) (*dtos.OpenPositionResp, deftype.Error) {
	user, err := r.userRepo.Get(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	band, err := r.bandRepo.Get(ctx, req.BandID)
	if err != nil {
		return nil, err
	}

	post, err := band.OpenPosition(user, req.Position, req.Contents)
	if err != nil {
		return nil, err
	}

	if err = r.postRepo.Save(ctx, post); err != nil {
		return nil, err
	}

	go r.produce(ctx, openPositionTopic, post)

	return &dtos.OpenPositionResp{
		ID:        post.ID,
		BandID:    post.BandID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Position:  post.Position,
		Contents:  post.Contents,
		IsOpened:  post.IsOpened,
	}, nil
}

func (r *MatchingService) SeekPosition(ctx context.Context, req *dtos.SeekPositionReq) (*dtos.SeekPositionResp, deftype.Error) {
	user, err := r.userRepo.Get(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	_, err = r.postRepo.GetUserPostByPositionAndGenre(ctx, user.ID, req.Position, req.FavoriteGenre)
	if err == nil {
		return nil, deftype.ErrDuplicatedRequest
	} else if !err.Equal(deftype.ErrNotFound) {
		return nil, err
	}

	post, err := user.SeekPosition(req.Contents, req.Position, req.FavoriteGenre)
	if err != nil {
		return nil, err
	}

	if err = r.postRepo.Save(ctx, post); err != nil {
		return nil, err
	}

	go r.produce(ctx, seekPositionTopic, post)

	return &dtos.SeekPositionResp{
		ID:            post.ID,
		FavoriteGenre: post.FavoriteGenre,
		CreatedAt:     post.CreatedAt,
		UpdatedAt:     post.UpdatedAt,
		Position:      post.Position,
		Contents:      post.Contents,
		IsOpened:      post.IsOpened,
	}, nil
}

func (r *MatchingService) produce(ctx context.Context, topic string, post *domains.Post) {
	m, err := post.MakeMessage()
	if err != nil {
		return
	}

	err = r.eventRepo.Produce(ctx, topic, m)
	if err != nil {
		return
	}
}
