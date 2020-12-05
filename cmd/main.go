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
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
)

//nolint:gochecknoglobals
var (
	buildTime   string
	buildGitTag string
)

func main() {
	kingpin.Version(appConfig.Version)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	var err error

	templateData := Inventory{}

	fileValuesLen := len(*appConfig.file)

	if fileValuesLen > 0 {
		templateData, err = parseValues(*appConfig.values)
		if err != nil {
			panic(errors.Wrap(err, fmt.Sprintf("error in parseValues file %s", *appConfig.values)))
		}
	}

	fileLen := len(*appConfig.file)

	if fileLen > 0 {
		result, err := parseFromFile(*appConfig.file, templateData)
		if err != nil {
			panic(errors.Wrap(err, fmt.Sprintf("error in parseFromFile file %s", *appConfig.file)))
		}

		fmt.Println(result)

		return
	}

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(errors.Wrap(err, "errors in os.Stdin.Stat"))
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("no pipe :(")
		fmt.Println("")
		fmt.Println("use: cat test | go-template ")

		return
	}

	err = parseFromPipe(templateData)
	if err != nil {
		panic(errors.Wrap(err, "errors in parseFromPipe"))
	}
}
