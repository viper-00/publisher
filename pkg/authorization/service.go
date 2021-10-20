package authorization

import "context"

type Service interface {
	Login(ctx context.Context, account, password string) (string, error)
	Logout(ctx context.Context, account, token string) (int, error)
	ServiceStatus(ctx context.Context) (int, error)
}
