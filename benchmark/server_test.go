// Copyright (c) 2014 Dataence, LLC. All rights reserved.
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

package benchmark

import (
	"testing"

	"github.com/surge/surgemq/service"
)

// Usage: go test -run=Publish0Topic1 -vv=2 -logtostderr [ -clients 2 ] [ -messages 100000 ] [ -msgsize 1024 ]
func TestServer(t *testing.T) {
	service.ListenAndServe(newTempContext(), "tcp://127.0.0.1:1883")
}
