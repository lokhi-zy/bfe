// Copyright (c) 2019 Baidu, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mod_static

import (
	"testing"
)

func TestConfLoad(t *testing.T) {
	config, err := ConfLoad("./testdata/mod_static/mod_static.conf", "")
	if err != nil {
		t.Errorf("ConfLoad() error: %v", err)
		return
	}

	if config.Basic.DataPath != "mod_static/static_rule.data" {
		t.Errorf("DataPath should be mod_static/static_rule.data, not %s", config.Basic.DataPath)
	}

	if config.Basic.MimeTypePath != "mod_static/mime_type.data" {
		t.Errorf("MimeTypePath should be mod_static/mime_type.data, not %s", config.Basic.MimeTypePath)
	}

	if !config.Basic.EnableCompress {
		t.Errorf("EnableCompress should be true")
	}
}
