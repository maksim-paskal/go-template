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
	"testing"
)

func TestParseFromSingleFile(t *testing.T) {
	t.Parallel()

	inventory, err := parseValues(testValuesFile)
	if err != nil {
		t.Fatal(err)
	}

	result, err := parseFromFile(testFile, inventory)
	if err != nil {
		t.Fatal(err)
	}

	ok := "START\n\ntext-\"some-env-value\"\nEND"

	if result != ok {
		t.Fatalf("file %s is incorrect\n=return=%s\nok=%s", testDir, result, ok)
	}
}

func TestParseFromDir(t *testing.T) {
	t.Parallel()

	inventory, err := parseValues(testDirValues)
	if err != nil {
		t.Fatal(err)
	}

	result, err := parseFromFile(testDir, inventory)
	if err != nil {
		t.Fatal(err)
	}

	ok := `START
T1 

T2
"some-env-value"valuetest3: value
test4: some-env-value
          test1:
            test2:
              test3: "value"
              test4: 'some-env-value'
END`

	if result != ok {
		t.Fatalf("file %s is incorrect\n=return=%s\nok=%s", testDir, result, ok)
	}
}
