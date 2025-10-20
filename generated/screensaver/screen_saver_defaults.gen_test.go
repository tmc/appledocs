// Code generated from Apple documentation for ScreenSaver. DO NOT EDIT.

package screensaver_test

import (
	"github.com/tmc/appledocs/generated/screensaver"
)

// Suppress unused import errors
var _ = screensaver.NewScreenSaverDefaults


// ExampleNewScreenSaverDefaultsForModuleWithName demonstrates how to create a ScreenSaverDefaults instance using NewScreenSaverDefaultsForModuleWithName.
// Returns a screen saver defaults instance that reads and writes defaults for the specified module.
func ExampleNewScreenSaverDefaultsForModuleWithName() {
	_ = screensaver.NewScreenSaverDefaultsForModuleWithName(
		"inModuleName", // inModuleName string
	)
	// Output:
}


