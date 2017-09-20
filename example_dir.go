package main

import (
	"fmt"

	"./dir"
)

func main() {
	d, err := dir.New("/Users/roporter/Documents")
	if err == nil {
		fmt.Println(d)
	}
	fmt.Println("ALL")
	fmt.Println(d.ListAll())
	fmt.Println("FILES")
	fmt.Println(d.ListFiles())
	fmt.Println("FOLDERS")
	fmt.Println(d.ListFolders())
	fmt.Println("ALL FILES")
	fmt.Println(d.ListFilesAll())
	fmt.Println("ALL FOLDERS")
	fmt.Println(d.ListFoldersAll())
	fmt.Println("ALL SPECIFIC FILES")
	fmt.Println(d.ListFilesType("go"))
	fmt.Println("ALL SPECIFIC GO FILES")
	fmt.Println(d.ListFilesType("go"))
	fmt.Println("ALL SPECIFIC TXT FILES")
	fmt.Println(d.ListFilesType("txt"))

	d2, err := dir.New("")
	if err == nil {
		fmt.Println(d2)
	}
	fmt.Println("ALL")
	fmt.Println(d2.ListAll())
	fmt.Println("FILES")
	fmt.Println(d2.ListFiles())
	fmt.Println("FOLDERS")
	fmt.Println(d2.ListFolders())
	fmt.Println("ALL FILES")
	fmt.Println(d2.ListFilesAll())
	fmt.Println("ALL FOLDERS")
	fmt.Println(d2.ListFoldersAll())
	fmt.Println("ALL SPECIFIC GO FILES")
	fmt.Println(d2.ListFilesType("go"))
	fmt.Println("ALL SPECIFIC TXT FILES")
	fmt.Println(d2.ListFilesType("txt"))

	output, err := d2.Run("ls -la", true)
	if err == nil {
		fmt.Println(output)
	}
}
