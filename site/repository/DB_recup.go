package repository

import (
	"database/sql"
	"fmt"
	models "server/models"
	"time"
)

//------------------ vérifier un utilisateur ------------------

func (utilisateur *UserRepository) VerifUser(connect string, password string) (*models.User, error) {
	var user models.User

	query := "SELECT id, username, password_hash, role, banned, is_connect FROM users WHERE username = ? AND password_hash = ?"
	rows, err := DbContext.Query(query, connect, password)
	if err != nil {
		return nil, fmt.Errorf("Erreur requête SQL : %v", err)
	}
	defer rows.Close()

	if rows.Next() {
		errScan := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Banned,
			&user.IsConnect,
		)
		if errScan != nil {
			return nil, fmt.Errorf("Erreur lors du scan des données : %v", errScan)
		}
		return &user, nil
	}

	return nil, fmt.Errorf("Utilisateur non trouvé ou mot de passe incorrect")
}

//------------------ récupération des topics ------------------

func RecupTopics(userID int) []models.Topic {
	query := `
		SELECT 
			t.id, t.name, c.name as categories, t.followers,
			COALESCE(ts.is_subscribed, 0) AS is_subscribed
		FROM topics t
		LEFT JOIN categories c ON t.category_id = c.id
		LEFT JOIN topic_subscriptions ts ON t.id = ts.topic_id AND ts.user_id = ?
	`

	rows, err := DbContext.Query(query, userID)
	if err != nil {
		fmt.Println(Red, "Erreur lors de la récupération des topics : ", err.Error(), Reset)
		return nil
	}
	defer rows.Close()

	var topics []models.Topic

	for rows.Next() {
		var topic models.Topic
		var isSubscribed int
		var id int

		err := rows.Scan(&id, &topic.Name, &topic.Category, &topic.Followers, &isSubscribed)
		if err != nil {
			fmt.Println(Red, "Erreur lors du scan : ", err.Error(), Reset)
			continue
		}

		topic.IsSubscribe = isSubscribed == 1
		topic.ListThread = []models.Thread{}
		topic.Id = id
		topics = append(topics, topic)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(Red, "Erreur itération : ", err.Error(), Reset)
	}

	return topics
}

func RecupTopicByID(topicId int) (models.Topic, error) {
	query := `
		SELECT 
			t.id, t.name, c.name AS category_name, t.followers
		FROM topics t
		JOIN categories c ON t.category_id = c.id
		WHERE t.id = ?
	`

	var topic models.Topic
	err := DbContext.QueryRow(query, topicId).Scan(&topic.Id, &topic.Name, &topic.Category, &topic.Followers)
	if err != nil {
		return models.Topic{}, fmt.Errorf("erreur récupération topic : %v", err)
	}

	return topic, nil
}

func RecupTopicByThreadID(threadID int) (models.Topic, error) {
	query := `
		SELECT 
			t.id, t.name, c.name AS category_name, t.followers
		FROM threads th
		JOIN topics t ON th.topic_id = t.id
		JOIN categories c ON t.category_id = c.id
		WHERE th.id = ?
	`

	var topic models.Topic
	err := DbContext.QueryRow(query, threadID).Scan(
		&topic.Id,
		&topic.Name,
		&topic.Category,
		&topic.Followers,
	)
	if err != nil {
		return models.Topic{}, fmt.Errorf("erreur récupération topic via threadID : %v", err)
	}

	return topic, nil
}

func RecupTopicByName(name string) (models.Topic, error) {
	query := `
		SELECT t.id, t.name, c.name AS category_name, t.followers
		FROM topics t
		JOIN categories c ON t.category_id = c.id
		WHERE t.name = ?
	`

	var topic models.Topic
	err := DbContext.QueryRow(query, name).Scan(&topic.Id, &topic.Name, &topic.Category, &topic.Followers)
	if err != nil {
		return models.Topic{}, fmt.Errorf("erreur récupération topic par nom : %v", err)
	}

	return topic, nil
}

// ------------------ récupération des threads ------------------
func RecupThreadsByTopicID(topicId int, userId int) []models.Thread {
	query := `
		SELECT 
			t.id, t.name, t.description, t.content, t.nb_like, t.nb_dislike, t.created_at,
			u.username, tp.name AS topic_name,
			COALESCE(tl.is_like, 0) AS is_liked
		FROM threads t
		JOIN users u ON u.id = t.author_id
		JOIN topics tp ON tp.id = t.topic_id
		LEFT JOIN thread_likes tl ON t.id = tl.thread_id AND tl.user_id = ?
		WHERE t.topic_id = ?
	`

	rows, err := DbContext.Query(query, userId, topicId)
	if err != nil {
		fmt.Println("Erreur récupération des threads :", err.Error())
		return nil
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var th models.Thread
		var isLike int
		var createdAtStr string

		err := rows.Scan(
			&th.Id, &th.Title, &th.Description, &th.Content,
			&th.NbLike, &th.NbDisLike, &createdAtStr,
			&th.NameCreator, &th.NameTopic, &isLike,
		)
		if err != nil {
			fmt.Println("Erreur scan thread :", err.Error())
			continue
		}

		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			fmt.Println("Erreur parse created_at :", err.Error())
		} else {
			duration := time.Since(parsedTime)
			th.TimeCreate = fmt.Sprintf("%v ago", duration.Round(time.Minute))
		}

		th.IsLike = isLike == 1
		threads = append(threads, th)
	}

	return threads
}

func RecupAllThreadsByDateDesc() ([]models.Thread, error) {
	query := `
		SELECT 
			t.id, t.name, t.description, t.content, t.nb_like, t.nb_dislike, t.created_at,
			u.username, tp.name AS topic_name
		FROM threads t
		JOIN users u ON u.id = t.author_id
		JOIN topics tp ON tp.id = t.topic_id
		ORDER BY t.created_at DESC
	`

	rows, err := DbContext.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la récupération des threads : %v", err)
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var th models.Thread
		var createdAtStr string

		err := rows.Scan(
			&th.Id, &th.Title, &th.Description, &th.Content,
			&th.NbLike, &th.NbDisLike, &createdAtStr,
			&th.NameCreator, &th.NameTopic,
		)
		if err != nil {
			fmt.Println("Erreur scan thread :", err.Error())
			continue
		}

		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			fmt.Println("Erreur parse created_at :", err.Error())
		} else {
			duration := time.Since(parsedTime)
			th.TimeCreate = fmt.Sprintf("%v ago", duration.Round(time.Minute))
		}

		threads = append(threads, th)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Erreur d'itération sur les threads : %v", err)
	}

	return threads, nil
}

func RecupThreadByID(threadID int) (models.Thread, error) {
	query := `
		SELECT 
			t.id, t.name, t.description, t.content, t.nb_like, t.nb_dislike, t.created_at,
			u.username, tp.name AS topic_name
		FROM threads t
		JOIN users u ON u.id = t.author_id
		JOIN topics tp ON tp.id = t.topic_id
		WHERE t.id = ?
	`

	var th models.Thread
	var createdAtStr string

	err := DbContext.QueryRow(query, threadID).Scan(
		&th.Id, &th.Title, &th.Description, &th.Content,
		&th.NbLike, &th.NbDisLike, &createdAtStr,
		&th.NameCreator, &th.NameTopic,
	)
	if err != nil {
		return models.Thread{}, fmt.Errorf("Erreur récupération thread (id=%d) : %v", threadID, err)
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		fmt.Println("Erreur parse created_at :", err.Error())
	} else {
		duration := time.Since(parsedTime)
		th.TimeCreate = fmt.Sprintf("%v ago", duration.Round(time.Minute))
	}

	return th, nil
}

func RecupThreadsByUserID(userID int) ([]models.Thread, error) {
	query := `
		SELECT 
			t.id, t.name, t.description, t.content, t.nb_like, t.nb_dislike, t.created_at,
			u.username, tp.name AS topic_name
		FROM threads t
		JOIN users u ON u.id = t.author_id
		JOIN topics tp ON tp.id = t.topic_id
		WHERE t.author_id = ?
		ORDER BY t.created_at DESC
	`

	rows, err := DbContext.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la récupération des threads par user : %v", err)
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var th models.Thread
		var createdAt string

		err := rows.Scan(&th.Id, &th.Title, &th.Description, &th.Content, &th.NbLike, &th.NbDisLike, &createdAt, &th.NameCreator, &th.NameTopic)
		if err != nil {
			fmt.Println("Erreur scan thread :", err)
			continue
		}

		parsed, err := time.Parse("2006-01-02 15:04:05", createdAt)
		if err == nil {
			th.TimeCreate = fmt.Sprintf("%v ago", time.Since(parsed).Round(time.Minute))
		}

		threads = append(threads, th)
	}

	return threads, nil
}

func RecupThreadsLikedByUser(userID int) ([]models.Thread, error) {
	query := `
		SELECT 
			t.id, t.name, t.description, t.content, t.nb_like, t.nb_dislike, t.created_at,
			u.username, tp.name AS topic_name
		FROM thread_likes tl
		JOIN threads t ON tl.thread_id = t.id
		JOIN users u ON t.author_id = u.id
		JOIN topics tp ON tp.id = t.topic_id
		WHERE tl.user_id = ? AND tl.is_like = 1
		ORDER BY t.created_at DESC
	`

	rows, err := DbContext.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("Erreur récupération threads likés : %v", err)
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var th models.Thread
		var createdAt string

		err := rows.Scan(&th.Id, &th.Title, &th.Description, &th.Content, &th.NbLike, &th.NbDisLike, &createdAt, &th.NameCreator, &th.NameTopic)
		if err != nil {
			fmt.Println("Erreur scan thread liked :", err)
			continue
		}

		parsed, err := time.Parse("2006-01-02 15:04:05", createdAt)
		if err == nil {
			th.TimeCreate = fmt.Sprintf("%v ago", time.Since(parsed).Round(time.Minute))
		}

		th.IsLike = true
		threads = append(threads, th)
	}

	return threads, nil
}

// ------------------ récupération des tags ------------------

func RecupTags() []models.Tag {
	query := `SELECT id, name FROM categories ORDER BY name ASC;`

	rows, err := DbContext.Query(query)
	if err != nil {
		fmt.Println("Erreur récupération des catégories :", err.Error())
		return nil
	}
	defer rows.Close()

	var tags []models.Tag

	for rows.Next() {
		var tag models.Tag
		err := rows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			fmt.Println(Red, "Erreur scan catégorie :", err.Error(), Reset)
			continue
		}
		tags = append(tags, tag)
	}

	return tags
}

func RecupCommentairesByThreadID(threadID int, userID int) []models.Comment {
	query := `
		SELECT u.username, c.content, c.nb_like, c.nb_dislike, c.created_at
		FROM comments c
		JOIN users u ON u.id = c.user_id
		WHERE c.thread_id = ?
		ORDER BY c.created_at DESC
	`

	rows, err := DbContext.Query(query, threadID)
	if err != nil {
		fmt.Println("Erreur récupération commentaires:", err.Error())
		return nil
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var c models.Comment
		err := rows.Scan(&c.UserComment, &c.Content, &c.NbLike, &c.NbDisLike, &c.TimeCreate)
		if err != nil {
			fmt.Println("Erreur scan commentaire:", err.Error())
			continue
		}
		comments = append(comments, c)
	}
	return comments
}

func RecupUserProfil(userId int) models.User {
	query := `
		SELECT id, username, role, banned, is_connect
		FROM users
		WHERE id = ?
	`

	var user models.User
	err := DbContext.QueryRow(query, userId).Scan(
		&user.Id, &user.Username, &user.Role, &user.Banned, &user.IsConnect,
	)

	if err != nil {
		fmt.Println("Erreur récupération profil :", err)
		return models.User{}
	}

	// On peut ajouter d’autres champs (comme TopicLike) plus tard
	return user
}

func CheckLike(userId, threadId int) (bool, bool, error) {
	query := `SELECT is_like FROM thread_likes WHERE user_id = ? AND thread_id = ? LIMIT 1`

	var isLike bool
	err := DbContext.QueryRow(query, userId, threadId).Scan(&isLike)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, false, nil
		}
		return false, false, err
	}
	return true, isLike, nil
}

func AddLike(userId, threadId int) error {
	tx, err := DbContext.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO thread_likes (user_id, thread_id, is_like) VALUES(?, ?, 1)", userId, threadId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE threads SET nb_like = nb_like + 1 WHERE id = ?", threadId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()

}

func RemoveLike(userId, threadId int) error {
	tx, err := DbContext.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM thread_likes WHERE user_id = ? AND thread_id = ?", userId, threadId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE threads SET nb_like = nb_like - 1 WHERE id = ?", threadId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// ------------------ Récupération utilisateur ------------------

func GetUserByID(id int) (models.User, error) {
	query := `
		SELECT id, username, password_hash, role, banned, is_connect,followers, is_follow
		FROM users
		WHERE id = ?
	`

	var user models.User
	var banned, isConnect, isFollow int

	err := DbContext.QueryRow(query, id).Scan(
		&user.Id, &user.Username, &user.Password,
		&user.Role, &banned, &isConnect, &user.Followers, &isFollow,
	)

	if err != nil {
		return user, fmt.Errorf("Erreur lors de la récupération de l'utilisateur : %v", err)
	}

	user.Banned = banned == 1
	user.IsConnect = isConnect == 1
	user.IsFollow = isFollow == 1

	// Récupération des topics follow
	topicQuery := `
		SELECT t.id, t.name, c.name AS category, t.followers
		FROM topic_subscriptions ts
		JOIN topics t ON ts.topic_id = t.id
		JOIN categories c ON t.category_id = c.id
		WHERE ts.is_subscribed = 1 AND ts.user_id = ?
	`

	rows, err := DbContext.Query(topicQuery, id)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des topics :", err)
		// On garde simplement une liste vide
		user.TopicLike = []models.Topic{}
	} else {
		defer rows.Close()
		for rows.Next() {
			var topic models.Topic
			if err := rows.Scan(&topic.Id, &topic.Name, &topic.Category, &topic.Followers); err != nil {
				fmt.Println("Erreur lors du scan d'un topic :", err)
				continue
			}
			topic.IsSubscribe = true
			user.TopicLike = append(user.TopicLike, topic)
		}
	}

	// ThreadLike, Friends, Subscribe, Post seront vides par défaut
	// A lancer d’autres méthodes si besoin

	return user, nil
}
func GetUserByUsername(username string) (models.User, error) {
	query := `
		SELECT id, username, password_hash, role, banned, is_connect,followers, is_follow
		FROM users
		WHERE username = ?
	`

	var user models.User
	var banned, isConnect, isFollow int

	err := DbContext.QueryRow(query, username).Scan(
		&user.Id, &user.Username, &user.Password,
		&user.Role, &banned, &isConnect, &user.Followers, &isFollow,
	)

	if err != nil {
		return user, fmt.Errorf("Erreur lors de la récupération de l'utilisateur : %v", err)
	}

	user.Banned = banned == 1
	user.IsConnect = isConnect == 1
	user.IsFollow = isFollow == 1

	// Récupération des topics follow
	topicQuery := `
		SELECT t.id, t.name, c.name AS category, t.followers
		FROM topic_subscriptions ts
		JOIN topics t ON ts.topic_id = t.id
		JOIN categories c ON t.category_id = c.id
		WHERE ts.is_subscribed = 1 AND ts.user_id = ?
	`

	rows, err := DbContext.Query(topicQuery, user.Id)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des topics :", err)
		user.TopicLike = []models.Topic{}
	} else {
		defer rows.Close()
		for rows.Next() {
			var topic models.Topic
			if err := rows.Scan(&topic.Id, &topic.Name, &topic.Category, &topic.Followers); err != nil {
				fmt.Println("Erreur lors du scan d'un topic :", err)
				continue
			}
			topic.IsSubscribe = true
			user.TopicLike = append(user.TopicLike, topic)
		}
	}

	// ThreadLike, Friends, Subscribe, Post seront vides par défaut
	// A lancer d’autres méthodes si besoin

	return user, nil
}

func GetTopicByName(name string) (models.Topic, error) {
	var topic models.Topic

	query := `SELECT id, name, followers FROM topics WHERE name = ? LIMIT 1`

	row := DbContext.QueryRow(query, name)

	if err := row.Scan(&topic.Id, &topic.Name, &topic.Followers); err != nil {
		return topic, fmt.Errorf("Erreur lors de la récupération du topic '%s': %v", name, err)
	}

	// tu peux lancer d’autres recherches, par ex. récupérer les threads liés
	// ou le suivi de l’utilisateur.

	return topic, nil
}
