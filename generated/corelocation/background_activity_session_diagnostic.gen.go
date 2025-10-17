// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BackgroundActivitySessionDiagnostic] class.
var backgroundActivitySessionDiagnosticClass = _BackgroundActivitySessionDiagnosticClass{objc.GetClass("CLBackgroundActivitySessionDiagnostic")}

type _BackgroundActivitySessionDiagnosticClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic

type BackgroundActivitySessionDiagnostic struct {
	objectivec.Object
}

// BackgroundActivitySessionDiagnosticFrom constructs a [BackgroundActivitySessionDiagnostic] from an unsafe.Pointer.
func BackgroundActivitySessionDiagnosticFrom(ptr unsafe.Pointer) BackgroundActivitySessionDiagnostic {
	return BackgroundActivitySessionDiagnostic{objectivec.Object{objc.ID(ptr)}}
}



