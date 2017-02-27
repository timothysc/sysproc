/*
<INSERT BLOCK>

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

package main

import (
	"error"
	"os"

	"github.com/timothysc/sysproc/node"
)

// The sole purpose of main in this example is a bootstrap
// point for the node libraries to start.  A node is the
// most fundamental unit that is used to wrap what could be
// 1:N subsystems.  The purpose of the node is simply
// to read configs start the appropriate things and be a
// signal handler for cleanup.  Once configs are loaded it
// simply hands off the work.
func main() {
	err := node.Run()
	os.Exit(err.)
}
