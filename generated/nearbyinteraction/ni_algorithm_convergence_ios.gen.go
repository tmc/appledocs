//go:build darwin && ios

// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NIAlgorithmConvergence


// iOS-only properties

// An array of reasons that contribute to the convergence status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIAlgorithmConvergence/reasons
func (n_ NIAlgorithmConvergence) Reasons() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("reasons"))
	return rv
}

// The current state of the framework’s Camera Assistance feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIAlgorithmConvergence/status-j61c
func (n_ NIAlgorithmConvergence) Status() NIAlgorithmConvergenceStatus {
	rv := objc.Send[NIAlgorithmConvergenceStatus](n_.ID, objc.Sel("status"))
	return rv
}





