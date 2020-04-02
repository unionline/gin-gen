package pkg

import (
	"github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"os"
)

const (
	Base = "/Users/ganodermaking/code/gin-gen/pkg/"
)

func NewDir(templatePath, newPath, appName string) error {
	pathSep := string(os.PathSeparator)
	
	err := os.MkdirAll(newPath, 0755)
	if err != nil {
		return err
	}
	
	dirs, _ := ioutil.ReadDir(templatePath)
	for _, fi := range dirs {
		templatePath := templatePath + pathSep + fi.Name()
		newPath := newPath + pathSep + fi.Name()
		
		if fi.IsDir() {
			err = NewDir(templatePath, newPath, appName)
			if err != nil {
				return err
			}
			continue
		}
		
		newFile, err := os.Create(newPath)
		if err != nil {
			return err
		}
		
		defer newFile.Close()
		
		data := map[string]string{
			"Appname": appName,
		}
		err = template.Must(template.ParseFiles(templatePath)).Execute(newFile, data)
		if err != nil {
			return err
		}
		
		logrus.Info(newPath)
	}
	
	return nil
}

func IsPath(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	
	return s.IsDir()
}
