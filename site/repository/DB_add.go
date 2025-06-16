package repository

import "fmt"

//------------------ ajouter un utilisateur ------------------

func AddUser(username string, password string) (int, error) {
	query := `INSERT INTO users (username, password_hash, role, banned, is_connect, followers, is_follow)
	          VALUES (?, ?, 'member', false, false, 0, false);`
	sqlResult, err := DbContext.Exec(
		query,
		username,
		password,
	)
	if err != nil {
		return -1, fmt.Errorf(Red, "Error when try to add user, error : ", err.Error(), Reset)
	}

	id, err := sqlResult.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf(Red, "Error when try to recup id user, error : ", err.Error(), Reset)
	}

	return int(id), nil
}

//------------------ ajouter un Topic ------------------

func AddTopic(name string, categoryID int) (int, error) {

	query := `INSERT INTO topics (name, category_id, followers)
	          VALUES (?, ?, 0);`
	sqlResult, err := DbContext.Exec(
		query,
		name,
		categoryID,
	)
	if err != nil {
		return -1, fmt.Errorf(Red, "Erreur ajout topic : ", err.Error(), Reset)
	}

	id, err := sqlResult.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf(Red, "Erreur récup ID topic : ", err.Error(), Reset)
	}

	return int(id), nil
}

//------------------ abonnement à un topic ------------------

func SubscribeToTopic(userID int, topicID int) error {
	tx, err := DbContext.Begin()
	if err != nil {
		return fmt.Errorf("Erreur début transaction : %v", err)
	}

	querySub := `
		INSERT INTO topic_subscriptions (user_id, topic_id, is_subscribed)
		VALUES (?, ?, 1)
		ON DUPLICATE KEY UPDATE is_subscribed = 1;
	`
	_, err = tx.Exec(querySub, userID, topicID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Erreur insert abonnement : %v", err)
	}

	queryUpdate := `
		UPDATE topics 
		SET followers = followers + 1
		WHERE id = ? AND NOT EXISTS (
			SELECT 1 FROM topic_subscriptions 
			WHERE user_id = ? AND topic_id = ? AND is_subscribed = 1
		);
	`
	_, err = tx.Exec(queryUpdate, topicID, userID, topicID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Erreur update followers : %v", err)
	}

	return tx.Commit()
}

//------------------ désabonnement à un topic ------------------

func UnsubscribeFromTopic(userID int, topicID int) error {
	tx, err := DbContext.Begin()
	if err != nil {
		return fmt.Errorf("Erreur début transaction : %v", err)
	}

	queryUnsub := `
		UPDATE topic_subscriptions 
		SET is_subscribed = 0 
		WHERE user_id = ? AND topic_id = ?;
	`
	result, err := tx.Exec(queryUnsub, userID, topicID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Erreur désabonnement : %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		queryUpdate := `
			UPDATE topics SET followers = followers - 1
			WHERE id = ? AND followers > 0;
		`
		_, err = tx.Exec(queryUpdate, topicID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("Erreur update followers : %v", err)
		}
	}

	return tx.Commit()
}

//------------------ ajouter un commentaire ------------------

func AjouterCommentaire(threadID int, userID int, content string) error {
	query := `
		INSERT INTO comments (thread_id, user_id, content, created_at)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
	`
	_, err := DbContext.Exec(query, threadID, userID, content)
	return err
}

//------------------ ajouter un thread ------------------

func AddThread(idTopic int, name string, content string, description string, authorId int) (int64, error) {
	query := ` INSERT INTO threads (name, content, description, topic_id, author_id, created_at) VALUES(?, ?, ?, ?, ?, NOW()) `
	result, err := DbContext.Exec(query, name, content, description, idTopic, authorId)
	if err != nil {
		return 0, fmt.Errorf("Erreur lors de l'insertion du thread : %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Impossible de récupérer l'ID généré : %v", err)
	}

	return id, nil
}
