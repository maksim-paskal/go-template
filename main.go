/*
Copyright paskal.maksim@gmail.com
Licensed under the Apache License, Version 2.0 (the "License")
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"text/template"

	"gopkg.in/alecthomas/kingpin.v2"
)

type appConfigType struct {
	Version string
	file    *string
}

var appConfig = appConfigType{
	Version: "1.0.1",
	file: kingpin.Flag(
		"file",
		"file to parse",
	).String(),
}

func goTemplateFunc() map[string]interface{} {
	return template.FuncMap{
		"getEnv": func(env string) string {
			return os.Getenv(env)
		},
		"quote": func(str string) string {
			return fmt.Sprintf("%q", str)
		},
	}
}
func parseFromPipe() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t, err := template.New(scanner.Text()).Funcs(goTemplateFunc()).Parse(scanner.Text())

		if err != nil {
			panic(err)
		}

		var tpl bytes.Buffer
		err = t.Execute(&tpl, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println(tpl.String())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func parseFromFile() {
	/*
			files, err := filepath.Glob(*appConfig.file)
			if err != nil {
				panic(err)
			}

			fmt.Println(files)

		name := path.Base(files[0])
		t, err := template.New(name).Funcs(goTemplateFunc()).ParseFiles(files...)

		if err != nil {
			panic(err)
		}

		var tpl bytes.Buffer
		err = t.Execute(&tpl, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println(tpl.String())

	*/

	var templates = template.Must(template.ParseGlob(*appConfig.file))
	templates.Funcs(goTemplateFunc())

	var tpl bytes.Buffer
	err := templates.Execute(&tpl, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(tpl.String())
}

func main() {
	kingpin.Version(appConfig.Version)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	if len(*appConfig.file) > 0 {
		parseFromFile()
	} else {
		fi, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}

		if fi.Mode()&os.ModeNamedPipe == 0 {
			fmt.Println("no pipe :(")
			fmt.Println("")
			fmt.Println("use: cat test | go-template ")
			return
		}

		parseFromPipe()
	}
}
