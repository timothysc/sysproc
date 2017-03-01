/*
Copyright 2017 Heptio Inc.

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

package subsystems

import (
	"fmt"
	//"github.com/timothysc/sysproc/pkg/subsystems/operators"
	//"github.com/timothysc/sysproc/pkg/subsystems/subsystems"
	//"github.com/timothysc/sysproc/pkg/subsystems/webhook"
)

// The factory will take in a heirarchical config
// and return an array of subsystems
//
// TODO: Config file formant etc.
func FactoryInit(cfg Config) ([]Interface, error) {

	fmt.Println("Here is where we would create and init subsystems")
	/*
		for _,subsystem := range subsystems {
			if err = subsystem.Init(); err != nil {
				break
			}
		}
	*/
	return nil, nil
}
