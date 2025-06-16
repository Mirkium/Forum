package repository

import (
	"database/sql"
	"fmt"
)

const (
	Red   = "\033[91m"
	Reset = "\033[0m"
)

type UserRepository struct{}

func IsFollowing(userId, topicId int) (bool, error) {
	query := `SELECT is_subscribed FROM topic_subscriptions WHERE user_id = ? AND topic_id = ? LIMIT 1`

	var isSubscribed int
	err := DbContext.QueryRow(query, userId, topicId).Scan(&isSubscribed)

	if err != nil {
		if err == sql.ErrNoRows {
			// Pas de ligne -> pas abonné
			return false, nil
		}
		return false, fmt.Errorf("Erreur lors de la vérification d'abonnement : %v", err)
	}

	return isSubscribed == 1, nil
}
