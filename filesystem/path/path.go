package path

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	SLASH              = string(os.PathSeparator)
	DEFAULT_DIR_ACCESS = 0755
)

func SplitPath(path string) []string {
	return strings.Split(path, SLASH)
}

func ParentPath(path string) string {
	list := SplitPath(path)
	var isAbs bool = false
	if strings.HasPrefix(path, SLASH) {
		list = list[1:]
		isAbs = true
	}

	if strings.HasSuffix(path, SLASH) {
		list = list[:len(list)-1]
	}

	if len(list) <= 0 {
		return SLASH
	} else {
		list = list[:len(list)-1]
		if len(list) <= 0 {
			if isAbs {
				return SLASH
			}
			return ""
		} else {
			parent := strings.Join(list, SLASH)
			if isAbs {
				parent = SLASH + parent
			}
			if parent != "" {
				parent += SLASH
			}
			return parent
		}
	}
}

func GetFileMd5(file string) string {
	f, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}
	h := md5.New()
	const BUFSIZE = 10240
	buf := make([]byte, BUFSIZE)
	for {
		rlen, err := f.Read(buf)
		if err != nil {
			break
		}
		h.Write(buf[0:rlen])
	}
	return hex.EncodeToString(h.Sum(nil))
}

func GetPath(file string) string {
	file = FormatPath(file)
	pos := strings.LastIndex(file, "/")
	return file[0:pos]
}

func BaseName(path string) string {
	list := SplitPath(path)
	if strings.HasSuffix(path, SLASH) {
		list = list[:len(list)-1]
	}
	if list != nil && len(list) > 0 {
		return list[len(list)-1]
	}
	return ""
}

func MkDirSpecificMode(path string, mode os.FileMode) error {
	exist, err := IsExist(path)
	if err == nil {
		if !exist {
			return os.MkdirAll(path, mode)
		} else if exist && err == nil {
			return nil
		}
	}
	return err
}

func MkDir(path string) error {
	exist, err := IsExist(path)
	if err == nil {
		if !exist {
			return os.MkdirAll(path, DEFAULT_DIR_ACCESS)
		} else if exist && err == nil {
			return nil
		}
	}
	return err
}

func IsDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if fi.IsDir() {
		return true, nil
	}
	return false, nil
}

func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func ListFilesRecursive(prefix, path string, b bool) []string {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}

	list := make([]string, 0, 10)
	var dir_name string
	if !b {
		dir_name = ""
	} else {
		dir_name = BaseName(path) + SLASH
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			tmp_list := ListFilesRecursive(prefix+dir_name, path+info.Name()+SLASH, true)
			list = append(list, tmp_list...)
		} else if info.Mode().IsRegular() {
			list = append(list, prefix+dir_name+info.Name())
		}
	}
	return list
}

func FileMode(path string) (os.FileMode, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fi.Mode(), nil
}

func FileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func RelativePath(file string, dir string) string {
	rfile := file[len(dir):] //相对路径
	rfile = strings.TrimLeft(rfile, "/")
	return rfile
}

func FormatPath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	path = strings.TrimRight(path, "/")
	return path
}

// SearchFile Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullpath string, err error) {
	for _, path := range paths {
		if fullpath = filepath.Join(path, filename); FileExists(fullpath) {
			return
		}
	}
	err = errors.New(fullpath + " not found in paths")
	return
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func RemoveFileIfExist(filename string) error {
	if FileExists(filename) {
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateDirIfNotExist(dir string) error {
	if !FileExists(dir) {
		err := os.Mkdir(dir, 0744)
		if err != nil {
			return fmt.Errorf("Failed to create dir %v %v", dir, err)
		}
	}
	return nil
}

func GetFilenameNoExtension(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n >= 0 {
		return s[:n]
	}
	return s
}

func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func FileExtension(str string) string {
	for i := len(str) - 1; i > -1; i-- {
		if str[i] == '.' {
			return str[i+1 : len(str)]
		}
	}
	return ""
}
