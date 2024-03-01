package models_rest

type BIMessage struct {
	ID       int64
	ChatId   int64
	SenderId int64
	Content  string
}

type BIChats struct {
	Id           int64
	FirstUserId  int64
	SecondUserId int64
}
