package constants

import (
	"log"
	"os/user"
)

var currUser string = getCurrUser()
var TodoDirectory string = "/Users/" + currUser + "/calcli/todo.txt"
var CompleteDirectory string = "/Users/" + currUser + "/calcli/complete.txt"
var TodoFolderDirectory string = "/Users/" + currUser + "/calcli/"
var TodoRenameDirectory string = "/Users/" + currUser + "/calcli/foo.txt"
var CompleteRenameDirectory string = "/Users/" + currUser + "/calcli/temp.txt"

func getCurrUser() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username
	return username
}
