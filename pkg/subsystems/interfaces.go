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

// This is a big old fat TODO.
type Config interface {
	TODO() error
}

// The subsystems.Interface is the
type Interface interface {
	// Init will initialize a subsystem from a configration block
	Init(cfg Config) error

	// Reconfig will attempt to gracefully adjust a running
	// subsystems config on the fly.  This is very useful
	// when trying to tweek params without forcing a restart.
	// It is triggered by a SIGUP.  In kubernetes this would be
	// update a ConfigMap + SIGUP to trigger a reload.
	Reconfig(cfg Config) error

	// DumpConfig will take the defaults of a subsystem
	// and dump the config to a string which can be appended
	// to an array of subsystems
	DumpConfig() (string, error)

	// Shutdown will try to gracefully shutdown a pr
	Shutdown() error
}
