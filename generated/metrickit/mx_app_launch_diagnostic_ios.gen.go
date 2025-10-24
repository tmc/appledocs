//go:build darwin && ios

// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for MXAppLaunchDiagnostic


// iOS-only properties

// The call stack tree associated with the app launch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchDiagnostic/callStackTree
func (m_ MXAppLaunchDiagnostic) CallStackTree() IMXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("callStackTree"))
	return rv
}

// The total app launch duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchDiagnostic/launchDuration
func (m_ MXAppLaunchDiagnostic) LaunchDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("launchDuration"))
	return rv
}





