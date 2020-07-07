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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

var (
	buildTime   string
	buildGitTag string
)

type appConfigType struct {
	Version string
	file    *string
	values  *string
}

var appConfig = appConfigType{
	Version: fmt.Sprintf("%s-%s", buildGitTag, buildTime),
	file: kingpin.Flag(
		"file",
		"file to parse",
	).String(),
	values: kingpin.Flag(
		"values",
		"values file to parse",
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
		err = t.Execute(&tpl, templateData)
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
	files, err := filepath.Glob(*appConfig.file)
	if err != nil {
		panic(err)
	}

	templateName := filepath.Base(files[0])

	// template file without _ prefix
	for _, i := range files {
		name := filepath.Base(i)

		if !strings.HasPrefix(name, "_") && !strings.EqualFold(name, "values.yaml") {
			templateName = name
			break
		}
	}

	templates := template.Must(template.New("").Funcs(goTemplateFunc()).ParseGlob(*appConfig.file))

	var tpl bytes.Buffer
	err = templates.ExecuteTemplate(&tpl, templateName, templateData)
	if err != nil {
		panic(err)
	}
	fmt.Println(tpl.String())
}

type Inventory struct {
	Values map[interface{}]interface{}
}

var templateData = Inventory{
	Values: make(map[interface{}]interface{}),
}

func main() {
	kingpin.Version(appConfig.Version)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	if len(*appConfig.values) > 0 {

		data, err := ioutil.ReadFile(*appConfig.values)
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(data, &templateData.Values)
		if err != nil {
			panic(err)
		}
	}
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
