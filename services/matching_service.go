package services

import (
	"context"
	"study-kafka-ddb/controllers/dtos"
	"study-kafka-ddb/domains/interfaces"
	"study-kafka-ddb/utils/deftype"
)

type MatchingService struct {
	userRepo interfaces.UserRepository
	bandRepo interfaces.BandRepository
	postRepo interfaces.PostRepository
}

func (r *MatchingService) OpenPosition(ctx context.Context, req *dtos.OpenPositionReq) (*dtos.OpenPositionResp, deftype.Error) {
	user, err := r.userRepo.Get(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	band, err := r.bandRepo.Get(ctx, req.BandID)
	if err != nil {
		return nil, err
	}

	wanted, err := band.OpenPosition(user, req.Position, req.Contents)
	if err != nil {
		return nil, err
	}

	if err = r.postRepo.Save(ctx, wanted); err != nil {
		return nil, err
	}

	return &dtos.OpenPositionResp{
		ID:        wanted.ID,
		BandID:    wanted.BandID,
		CreatedAt: wanted.CreatedAt,
		UpdatedAt: wanted.UpdatedAt,
		Position:  wanted.Position,
		Contents:  wanted.Contents,
		IsOpened:  wanted.IsOpened,
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

	resume, err := user.SeekPosition(req.Contents, req.Position, req.FavoriteGenre)
	if err != nil {
		return nil, err
	}

	if err = r.postRepo.Save(ctx, resume); err != nil {
		return nil, err
	}

	return &dtos.SeekPositionResp{
		ID:            resume.ID,
		FavoriteGenre: resume.FavoriteGenre,
		CreatedAt:     resume.CreatedAt,
		UpdatedAt:     resume.UpdatedAt,
		Position:      resume.Position,
		Contents:      resume.Contents,
		IsOpened:      resume.IsOpened,
	}, nil
}
