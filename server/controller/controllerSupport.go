package controller

const (
	Red   = "\033[91m"
	Reset = "\033[0m"
)

type Profile struct {
	Username string
	Bio      string
	Avatar   string
}

type FriendRequest struct {
	From string
	To   string
	Type string // "request" or "accept"
}

type Thread struct {
	ID    string
	Title string
}

type Message struct {
	ID       string
	ThreadID string
	Content  string
}

func GetUserProfile(username string) Profile {
	return Profile{
		Username: username,
		Bio:      "Just a Go developer.",
		Avatar:   "/assets/img/default.png",
	}
}

func UpdateUserProfile(username string, p Profile) {

}

func SendFriendRequest(from, to string) {

}

func AcceptFriendRequest(from, to string) {

}

func GetPrivateThreadsForUser(username string) []Thread {
	return []Thread{
		{ID: "p1", Title: "Private chat with best friend"},
	}
}

func GetAllThreads() []Thread {
	return []Thread{
		{ID: "1", Title: "First Thread"},
		{ID: "2", Title: "Second Thread"},
	}
}

func GetThreadByID(id string) Thread {
	return Thread{ID: id, Title: "Thread " + id}
}

func SaveMessage(m Message) {
	// save logic here
}

func GetMessageByID(id string) Message {
	return Message{ID: id, ThreadID: "1", Content: "Sample message"}
}

func UpdateLikeDislike(id, action string) {
	// update logic here
}

func Search(query string) []Thread {
	return []Thread{
		{ID: "1", Title: "Result for " + query},
	}
}

func GetThreadsByTag(tag string) []Thread {
	return []Thread{
		{ID: "1", Title: "Tagged: " + tag},
	}
}

func DeleteThread(id string) {
	// delete logic here
}

func DeleteMessage(id string) {
	// delete logic here
}

func BanUser(username string) {
	// ban logic here
}
