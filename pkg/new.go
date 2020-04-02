package pkg

import (
	"errors"
	"os"
)

const TemplatePath = "/pkg/templates"

func NewRun(args []string) error {
	appName := args[0]
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	
	templatePath := pwd + TemplatePath
	newPath := pwd + string(os.PathSeparator) + appName
	if IsPath(newPath) {
		return errors.New(newPath + " exists.")
	}
	
	err = NewDir(templatePath, newPath, appName)
	if err != nil {
		return err
	}
	
	return nil
}
