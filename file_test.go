package gobase

import (
	"path"
	"runtime"
	"strings"
	"testing"
)

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func TestFileFolderExist(t *testing.T) {
	checkFile := getCurrentPath() + "/testdata/not_exist"
	if FileOrFolderExist(checkFile) == false {
		t.Log("test success")
	} else {
		t.Fatal(checkFile, " is not exist, but return true, check failed")
	}

	checkFile = getCurrentPath() + "/testdata/exist_folder"
	if FileOrFolderExist(checkFile) == true {
		t.Log("test success")
	} else {
		t.Fatal(checkFile, " is exist, but return false check failed")
	}
}

func TestIsFolder(t *testing.T) {
	checkFile := getCurrentPath() + "/testdata/exist_folder"
	if IsFolder(checkFile) == true {
		t.Log("test success")
	} else {
		t.Fatal(checkFile, " is exist, but return false check failed")
	}

	checkFile = getCurrentPath() + "/testdata/exist_folder/.gitkeeper"
	if IsFolder(checkFile) == false {
		t.Log("test success")
	} else {
		t.Fatal(checkFile, " is a file, but return true check failed")
	}
}

func TestIsFile(t *testing.T) {
	checkFile := getCurrentPath() + "/testdata/exist_folder/.gitkeeper"
	if IsFile(checkFile) == true {
		t.Log("test success")
	} else {
		t.Fatal(checkFile, " is exist and a file, but return false check failed")
	}
}

func TestReadFileToString(t *testing.T) {
	checkFile := getCurrentPath() + "/testdata/myfile.log"
	text, err := ReadFileAsString(checkFile)
	if err == 0 && strings.Compare(*text, "This is a test file") == 0 {
		t.Log("test success")
	} else {
		t.Fatal(checkFile, " read error!")
	}

}
