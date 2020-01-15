package repository

type User struct {
	Id          string
	Name        string
	Age         int
	Email       string
	Picture     string
	Description string
}

type Dog struct {
	Id          string
	Name        string
	Age         int
	Type        string
	Picture     string
	Description string
}
type DogOwner struct {
	Id    int64
	Name  string
	Age   int
	Email string
}

type DogWalker struct {
	Id    int
	Name  string
	Age   int
	Email string
}
