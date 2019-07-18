package util

var id ObjectID = 0

func GenNewID() ObjectID {
	id += 1
	return id
}
