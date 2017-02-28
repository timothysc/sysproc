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

package node

import (
	"fmt"

	"github.com/timothysc/sysproc/pkg/utils/config"
)

func Run() error {

	// 0 - Load the config file
	/*cfg,*/ err := config.Load()
	if err != nil {

		//
		// err := factory.Init(cfg)

		// The following waits on a signal
		// TODO: Add all signal handling into place.
		// for {
		//    select {}
		//}
	} else {
		fmt.Printf("Error Loading Node Config: %v\n", err)
	}

	return err
}
