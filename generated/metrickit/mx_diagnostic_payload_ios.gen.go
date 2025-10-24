//go:build darwin && ios

// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MXDiagnosticPayload


// iOS-only properties

// The diagnostic reports for the app launch time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/appLaunchDiagnostics
func (m_ MXDiagnosticPayload) AppLaunchDiagnostics() []MXAppLaunchDiagnostic {
	rv := objc.Send[[]MXAppLaunchDiagnostic](m_.ID, objc.Sel("appLaunchDiagnostics"))
	return rv
}





