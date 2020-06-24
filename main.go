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
)

func main() {
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

	funcs := template.FuncMap{
		"getEnv": func(env string) string {
			return os.Getenv(env)
		},
		"quote": func(str string) string {
			return fmt.Sprintf("%q", str)
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t, err := template.New(scanner.Text()).Funcs(funcs).Parse(scanner.Text())

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
