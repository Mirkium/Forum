package models

import (
	"time"
)

// ================================= Structure DB =================================

type User struct {
	Id         int
	Username   string
	Password   string
	Role       string
	Banned     bool
	IsConnect  bool
	ThreadLike []Thread
	TopicLike  []Topic
	Friends    []User
	Followers  int
	Subscribe  []User
	Post       []Thread
	IsFollow   bool
}

type Thread struct {
	Id          int
	Title       string
	NameCreator string
	Description string
	Content     string
	NbLike      int
	NbDisLike   int
	Comment     []Comment
	IsLike      bool
	TimeCreate  string
}

type Comment struct {
	UserComment string
	Content     string
	NbLike      int
	NbDisLike   int
	IsLike      bool
	TimeCreate  time.Time
}

type Topic struct {
	Id          int
	Name        string
	Category    string
	Followers   int
	ListThread  []Thread
	IsSubscribe bool
}

type Tag struct {
	Id   int
	Name string
}

// ================================= Structure Web =================================

type Error struct {
	ValueError bool
	IsLogin    bool
}
