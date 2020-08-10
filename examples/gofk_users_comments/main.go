package main

//go:generate gofk

// A User is a person on the internet with an opinion.
type User struct {
	Id uint64 `fk:"opt:primary"`
	// "opt:primary" indicates this is the default field for referencing.
	Name string
}

// A Comment is a piece of text on the internet that is best avoided.
type Comment struct {
	Id     uint64
	UserId uint64 `fk:"ref:User;opt:cascade"`
	// "ref:User" tells gofk that a Comment must reference User.Id, because Id is User's primary field.
	// "opt:cascade" tells gofk that if the User referenced in UserId is deleted or modified, this Comment should be also.
	Parent uint64 `fk:"ref:Comment.Id;opt:null"`
	// "ref:Comment.Id" tells gofk that Parent can refer to another Comment.
	// "opt:null" tells gofk that this reference can be the zero value of the type, to indicate no reference.
}

func main() {
	panic("not yet!")
}
