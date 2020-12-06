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
	"flag"
	"fmt"
)

type appConfigType struct {
	Version     string
	showVersion *bool
	file        *string
	values      *string
}

//nolint:gochecknoglobals
var appConfig = appConfigType{
	Version:     fmt.Sprintf("%s-%s", gitVersion, buildTime),
	showVersion: flag.Bool("version", false, "show version"),
	file:        flag.String("file", "", "file to parse"),
	values:      flag.String("values", "", "values file to parse"),
}
