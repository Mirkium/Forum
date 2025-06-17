package models

type Home struct {
	Profil    User
	Post      []Thread
	ListTopic []Topic
	ListTag   []Tag
}

type TopicPage struct {
	Topic      Topic
	Profil     User
	ListTopic  []Topic
	ListThread []Thread
}

type AddTopic_tag struct {
	ListTag []Tag
	Profil  User
}

type AddTopic_Thread struct {
	TopicId int
	Profil  User
}

type Profil_Page struct {
	Profil    User
	ListTopic []Topic
}
