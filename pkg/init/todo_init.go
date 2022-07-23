package init

import (
	"fmt"
	"os"
	"os/user"
	"path"
)

func main() {

	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	dirPath := path.Join("/Users", user.Username, "calcli")
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		fmt.Println(err)
	}

	filepath := path.Join(dirPath, "todo.txt")
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	fmt.Printf("File created successfully at %s\n", filepath)
}

func init() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	dirPath := path.Join("/Users", user.Username, "calcli")
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		fmt.Println(err)
	}

	filepath := path.Join(dirPath, "todo.txt")
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	fmt.Printf("File created successfully at %s\n", filepath)
}
