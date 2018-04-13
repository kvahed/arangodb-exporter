// Copyright (c) 2016 Pulcy.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package settings

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	InitialVersion = "0.0.0+git"
	VersionFile    = "VERSION"
)

// Try to read VERSION
func ReadVersion(projectDir string) (string, error) {
	if data, err := ioutil.ReadFile(filepath.Join(projectDir, VersionFile)); err != nil {
		if os.IsNotExist(err) {
			return InitialVersion, nil
		} else {
			return "", maskAny(err)
		}
	} else {
		return strings.TrimSpace(string(data)), nil
	}
}
