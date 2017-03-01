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

	"github.com/spf13/viper"
)

type foo struct {
	err error
}

func (f foo) TODO() error {
	fmt.Println("Config-TODO")
	return nil
}

// Load will load the configation data for the process.
// I'm not going to pretend and create some massive
// arg-tree to doom.  Why, b/c it always ends in some
// cat-ass-trophy of maintainability.
//
// Config files allow me to simply say load this file and go..
// don't have one, we have a generate option for you.
func LoadConfig( /*args meh*/ ) (Config, error) {

	fmt.Println("Loading the Configs!!!")
	// 0. load in configurations data before doing anything
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/sysproc")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("No configuration file loaded - using defaults")
	}

	// If no config is found, use the default(s)
	viper.SetDefault("msg", "Hello World (default)")
	theMessage := viper.GetString("msg")
	fmt.Println(theMessage)

	// BIG OLD FAT TODO HERE!

	return &foo{}, err
}
