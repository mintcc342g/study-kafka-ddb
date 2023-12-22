package interfaces

import (
	"context"
	"study-kafka-ddb/domains"
	"study-kafka-ddb/utils/deftype"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*domains.User, deftype.Error)
	Save(ctx context.Context, user *domains.User) deftype.Error
}
