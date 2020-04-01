/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/unionline/gin-gen/pkg"
	"os"
	
	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var currentDir = false
var logo = figure.NewFigure("gin-gen", "", true).String() + "\n"

var newCmd = &cobra.Command{
	Use:   "new appName",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(logo)
		fmt.Println("Creating application...")
		
		if len(args) == 0 {
			logrus.Fatal("appName empty")
			_ = cmd.Help()
			os.Exit(1)
		}
		
		if err := pkg.NewRun(args); err != nil {
			logrus.Fatal(err)
		}
		
		fmt.Println("New application successfully created!")
	},
}

func init() {
	newCmd.Flags().BoolVarP(&currentDir, "currentDir", "", false, "if create current dir, default is false")
	rootCmd.AddCommand(newCmd)
}
