package interfaces

import (
	"context"
	"study-kafka-ddb/domains"
	"study-kafka-ddb/domains/enums"
	"study-kafka-ddb/utils/deftype"
)

type UserRepository interface {
	Get(ctx context.Context, userID enums.UserID) (*domains.User, deftype.Error)
	GetByEmail(ctx context.Context, email string) (*domains.User, deftype.Error)
	Save(ctx context.Context, user *domains.User) deftype.Error
}

type BandRepository interface {
	Get(ctx context.Context, bandID enums.BandID) (*domains.Band, deftype.Error)
	GetByPositionAndGenre(ctx context.Context, position enums.BandPosition, genre enums.Genre) (*domains.Band, deftype.Error)
}

type PostRepository interface {
	Save(ctx context.Context, post *domains.Post) deftype.Error
	Get(ctx context.Context, postID int64) (*domains.Post, deftype.Error)
	GetUserPostByPositionAndGenre(ctx context.Context, userID enums.UserID, position enums.BandPosition, genre enums.Genre) (*domains.Post, deftype.Error)
}

type EventRepository interface {
	Produce(ctx context.Context, topic string, msg []byte) deftype.Error
	Subscribe(ctx context.Context, topic string, handler func(context.Context, []byte) error) deftype.Error
}

type EmailRepository interface {
	Send(ctx context.Context, to, subejct, body string) deftype.Error
}
