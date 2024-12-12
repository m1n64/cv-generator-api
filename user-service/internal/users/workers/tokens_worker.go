package workers

import (
	"fmt"
	"time"
	"user-service/internal/users/repositories"
	"user-service/pkg/utils"
)

func StartRemoveExpiredTokensWorker(repo repositories.TokenRepository, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	logger := utils.GetLogger()

	for {
		<-ticker.C
		logger.Info(fmt.Sprintf("Removing expired tokens from last %s...", interval.String()))

		err := repo.DeleteExpiredTokens()
		if err != nil {
			logger.Error(fmt.Sprintf("Error removing expired tokens: %v", err))
		} else {
			logger.Info("Expired tokens removed successfully.")
		}
	}
}
