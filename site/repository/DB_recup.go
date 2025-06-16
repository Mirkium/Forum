package repository

import (
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

// ------------------ récupération des threads ------------------
func RecupThreadsByTopicID(topicId int, userId int) []models.Thread {
	query := `
		SELECT 
			t.id, t.name, t.description, t.content, t.nb_like, t.nb_dislike, t.created_at,
			u.username,
			COALESCE(tl.is_like, 0) AS is_liked
		FROM threads t
		JOIN users u ON u.id = t.author_id
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
			&th.NameCreator, &isLike,
		)
		if err != nil {
			fmt.Println("Erreur scan thread :", err.Error())
			continue
		}

		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			fmt.Println("Erreur parse created_at :", err.Error())
		} else {
			// On garde le décalage en heures, minutes...
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
			u.username
		FROM threads t
		JOIN users u ON u.id = t.author_id
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
			&th.NameCreator,
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


//------------------ récupération des tags ------------------

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

func CheckLike(userId, threadId int) bool {
	query := `SELECT 1 FROM thread_likes WHERE user_id = ? AND thread_id = ? LIMIT 1`
	var exist int
	err := DbContext.QueryRow(query, userId, threadId).Scan(&exist)
	return err == nil
}

func AddLike(userId, threadId int) error {
	_, err := DbContext.Exec("INSERT INTO thread_likes (user_id, thread_id, is_like) VALUES(?, ?, 1)", userId, threadId)
	return err
}

func RemoveLike(userId, threadId int) error {
	_, err := DbContext.Exec("DELETE FROM thread_likes WHERE user_id = ? AND thread_id = ?", userId, threadId)
	return err
}
