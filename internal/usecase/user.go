package usecase

import (
	"context"

	"github.com/Karzoug/meower-common-go/ucerr"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"

	"github.com/Karzoug/meower-web-service/internal/entity"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc/user"
)

type UsersUseCase struct {
	userServiceClient user.Client
}

func NewUsersUseCase(userClient user.Client) UsersUseCase {
	return UsersUseCase{}
}

func (uc UsersUseCase) GetUser(ctx context.Context, userID xid.ID) (entity.UserProjection, error) {
	return entity.UserProjection{}, ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (uc UsersUseCase) GetUserWithUsername(ctx context.Context, username string) (entity.UserProjection, error) {
	return entity.UserProjection{}, ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

// Returns information about an authorized user.
func (uc UsersUseCase) GetMe(ctx context.Context, authUserID xid.ID) (entity.User, error) {
	return entity.User{}, ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (uc UsersUseCase) UpdateMe(ctx context.Context, authUserID xid.ID, u entity.User) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (uc UsersUseCase) FollowAnotherUser(ctx context.Context, authUserID, targetUserID xid.ID) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (uc UsersUseCase) UnfollowAnotherUser(ctx context.Context, authUserID, targetUserID xid.ID) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (uc UsersUseCase) MuteAnotherUser(ctx context.Context, authUserID, targetUserID xid.ID) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (uc UsersUseCase) UnmuteAnotherUser(ctx context.Context, authUserID, targetUserID xid.ID) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}
