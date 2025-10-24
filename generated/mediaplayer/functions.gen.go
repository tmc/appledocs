// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

/* debug [functions.gen.go]: Generating 3 functions for MediaPlayer */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// MediaPlayer Functions (3 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MPVolumeSettingsAlertHide func()
	_MPVolumeSettingsAlertIsVisible func() bool
	_MPVolumeSettingsAlertShow func()
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MPVolumeSettingsAlertHide, lib, "MPVolumeSettingsAlertHide")
	tryRegister(&_MPVolumeSettingsAlertIsVisible, lib, "MPVolumeSettingsAlertIsVisible")
	tryRegister(&_MPVolumeSettingsAlertShow, lib, "MPVolumeSettingsAlertShow")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Hides the alert panel that controls the system volume.

// Hides the alert panel that controls the system volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeSettingsAlertHide()
func MPVolumeSettingsAlertHide() {
	_MPVolumeSettingsAlertHide()
}/* debug [functions.gen.go/function]: MPVolumeSettingsAlertHide */

// Returns a Boolean value indicating whether the volume alert panel is currently visible.

// Returns a Boolean value indicating whether the volume alert panel is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeSettingsAlertIsVisible()
func MPVolumeSettingsAlertIsVisible() bool {
	return _MPVolumeSettingsAlertIsVisible()
}/* debug [functions.gen.go/function]: MPVolumeSettingsAlertIsVisible */

// Displays an alert panel for controlling the system volume.

// Displays an alert panel for controlling the system volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeSettingsAlertShow()
func MPVolumeSettingsAlertShow() {
	_MPVolumeSettingsAlertShow()
}/* debug [functions.gen.go/function]: MPVolumeSettingsAlertShow */




