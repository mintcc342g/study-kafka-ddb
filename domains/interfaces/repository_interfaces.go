package interfaces

import (
	"context"
	"study-kafka-ddb/domains"
	"study-kafka-ddb/utils/deftype"
)

type UserRepository interface {
	Save(ctx context.Context, user *domains.User) deftype.Error
}
