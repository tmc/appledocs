// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ServiceSessionDiagnostic] class.
var serviceSessionDiagnosticClass = _ServiceSessionDiagnosticClass{objc.GetClass("CLServiceSessionDiagnostic")}

type _ServiceSessionDiagnosticClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic

type ServiceSessionDiagnostic struct {
	objectivec.Object
}

// ServiceSessionDiagnosticFrom constructs a [ServiceSessionDiagnostic] from an unsafe.Pointer.
func ServiceSessionDiagnosticFrom(ptr unsafe.Pointer) ServiceSessionDiagnostic {
	return ServiceSessionDiagnostic{objectivec.Object{objc.ID(ptr)}}
}



