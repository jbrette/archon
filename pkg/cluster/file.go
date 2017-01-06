/*
Copyright 2016 The Archon Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

type FileSpec struct {
	Name               string `json:"name,omitempty" yaml:"name,omitempty"`
	Encoding           string `json:"encoding,omitempty" yaml:"encoding,omitempty" valid:"^(base64|b64|gz|gzip|gz\\+base64|gzip\\+base64|gz\\+b64|gzip\\+b64)$"`
	Content            string `json:"content,omitempty" yaml:"content,omitempty"`
	Template           string `json:"template,omitempty" yaml:"template,omitempty"`
	Owner              string `json:"owner,omitempty" yaml:"owner,omitempty"`
	Path               string `json:"path,omitempty" yaml:"path,omitempty"`
	RawFilePermissions string `json:"permissions,omitempty" yaml:"permissions,omitempty" valid:"^0?[0-7]{3,4}$"`
}
