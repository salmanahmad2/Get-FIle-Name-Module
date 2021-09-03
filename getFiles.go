package getFiles

import (
	"fmt"
	"io/fs"
)

func Files(directory []fs.FileInfo) {

	//file.Name() will get all the name of the files in the folder
	for _, file := range directory {
		fmt.Println(file.Name())
	}
}
