package dir

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	All = iota
	Files
	Folders
)

type Dir struct {
	path string
}

func New(path string) (Dir, error) {
	if path == "" {
		fullpath, err := os.Getwd()
		if err != nil {
			return Dir{}, err
		}
		return Dir{path: fullpath}, nil
	}
	return Dir{path: path}, nil
}

func (d Dir) Run(cmd string, shell bool) (string, error) {
	if shell {
		cmd2 := exec.Command("sh", "-c", cmd)
		var outb, errb bytes.Buffer
		cmd2.Stdout = &outb
		cmd2.Stderr = &errb
		err := cmd2.Run()
		if err != nil {
			log.Fatal(err)
		}
		return string(outb.String()), err
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			log.Fatal(err)
		}
		return string(out), err
	}
}

func (d Dir) ListAll() []string {
	return d.list(All, true, "*")
}

func (d Dir) ListFiles() []string {
	return d.list(Files, false, "*")
}

func (d Dir) ListFilesAll() []string {
	return d.list(Files, true, "*")
}

func (d Dir) ListFolders() []string {
	return d.list(Folders, false, "*")
}

func (d Dir) ListFoldersAll() []string {
	return d.list(Folders, true, "*")
}

func (d Dir) ListFilesType(ext string) []string {
	return d.list(Files, true, ext)
}

func (d Dir) list(s int, hidden bool, extension string) []string {
	var output []string
	files, _ := ioutil.ReadDir(d.path)
	for _, f := range files {
		if d.shouldInclude(s, hidden, f, extension) {
			output = append(output, f.Name())
		}
	}
	return output
}

func (d Dir) getExtension(f string) string {
	splits := strings.Split(f, ".")
	if len(splits) > 0 {
		return splits[len(splits)-1]
	}
	return ""
}

func (d Dir) shouldInclude(s int, h bool, f os.FileInfo, e string) bool {
	if s == 0 {
		return true
	} else if s == 1 {
		if f.IsDir() {
			return false
		}
		// Definately a file
		if len(f.Name()) > 0 {
			// We have a name for the file
			if f.Name()[0:1] == "." {
				if h {
					if e == "*" {
						return true
					} else {
						if d.getExtension(f.Name()) == e {
							return true
						} else {
							return false
						}
					}
				}
				return false
			}
			if e == "*" {
				return true
			} else {
				if d.getExtension(f.Name()) == e {
					return true
				} else {
					return false
				}
			}
		}
		return true
	} else if s == 2 {
		if f.IsDir() {
			// We have a name for the file
			if f.Name()[0:1] == "." {
				if h {
					return true
				}
				return false
			}
			return true
		}
	}
	return false
}
