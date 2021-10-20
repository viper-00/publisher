package authorization

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/log"
)

type authService struct{}

func NewService() Service {
	return &authService{}
}

// implement service interface;

func (a *authService) Login(_ context.Context, account, password string) (string, error) {
	return "", nil
}

func (a *authService) Logout(_ context.Context, account, token string) (int, error) {
	return http.StatusOK, nil
}

func (a *authService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking the Service health...")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
