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
	"os"
	"testing"
)

func TestParseFromSingleFile(t *testing.T) {
	os.Setenv("SOMEVAR", "some-env-value-file")

	inventory, err := parseValues(testValuesFile)
	if err != nil {
		t.Fatal(err)
	}

	result, err := parseFromFile(testFile, inventory)
	if err != nil {
		t.Fatal(err)
	}

	if result != "START\n\ntext-\"some-env-value-file\"\nEND" {
		t.Fatalf("file %s is incorrect, return=%s", testFile, result)
	}
}

func TestParseFromDir(t *testing.T) {
	os.Setenv("SOMEVAR", "some-env-value-dir")

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
"some-env-value-dir"valuetest3: value
test4: some-env-value-dir
          test1:
            test2:
              test3: "value"
              test4: 'some-env-value-dir'
END`

	if result != ok {
		t.Fatalf("file %s is incorrect\n=return=%s\nok=%s", testDir, result, ok)
	}
}
