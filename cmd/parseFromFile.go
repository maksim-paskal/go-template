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
	"bytes"
	"path/filepath"
	"strings"
	"text/template"
)

func parseFromFile(fileName string, templateData Inventory) (string, error) {
	files, err := filepath.Glob(fileName)
	if err != nil {
		return "", err
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

	t := template.New("")
	templates := template.Must(t.Funcs(goTemplateFunc(t)).ParseGlob(fileName))

	var tpl bytes.Buffer

	err = templates.ExecuteTemplate(&tpl, templateName, templateData)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
