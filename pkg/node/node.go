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
	"os"
	"os/signal"
	"syscall"

	subsys "github.com/timothysc/sysproc/pkg/subsystems"
)

// A node is the most fundamental unit that is used to
// wrap what could be 1:N subsystems.  The purpose of the
// node is simply to read configs start the appropriate
// things and be a signal handler for cleanup.  Once configs
// are loaded it simply hands off the work.
// In Kubernetes "speak" it would be a "HYPER"-whatevers.
func Run() error {
	var err error
	// 0 - Load the master config which is a heirarchical
	// config file.
	/*cfg*/ _, err = subsys.LoadConfig()
	if err != nil {

		// 1 - Load the master config which is a heirarchical
		// config file.
		// subsystems, err := subsys.FactoryInit(cfg)

		if err != nil {
			// The following waits on a signal
			c := make(chan os.Signal, 1)
			signal.Notify(c,
				syscall.SIGINT,  // Ctrl+C
				syscall.SIGTERM, // Termination Request
				syscall.SIGSEGV, // FullDerp
				syscall.SIGABRT, // Abnormal termination
				syscall.SIGILL,  // illegal instruction
				syscall.SIGFPE,  // floating point - this is why we can't have nice things
				syscall.SIGHUP)  // RECONFIG!

			// Here is the main idle loop, it simply waits on process signals
			// to be raised.  This actually allows for much more gracefull cleanup
			// when processes are deep down in the bowels of 500-goroutines.
			// Instead of panic(), they can call a raise() wrapper (TO BE WRITTEN)
			// raise wrapper would dump stack details before the call.
			// The reason it's useful to not just panic, is to allow for ordered
			// shutdown or handoff... "good cluster citizenship"

			fmt.Println("Entering idleloop...")
			for {
				sig := <-c
				switch sig {
				case syscall.SIGHUP:
					fmt.Printf("\nHere is where we loop on reconfig!\n")
					// right now this is essentially empty
					/*for _, subsystem := range subsystems {
						// TODO
						// You need to break down the cfgs..
						err = subsystem.Reconfig(cfg)
						if err != nil {
							break
						}
					}*/
				default:
					fmt.Printf("\nSignal (%v) Detected, Shutting Down\n", sig)
					// again there is no return so it's doing nothing.
					/*for _, subsystem := range subsystems {
						err = subsystem.Shutdown()
						// we may want this to be an error array
						if err != nil {
							break idleloop
						}

					}*/

				} // end switch
				if err != nil {
					break
				}
			} // end for
		}
	}
	return err
}
