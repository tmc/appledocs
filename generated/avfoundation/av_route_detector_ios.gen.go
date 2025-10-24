//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for RouteDetector


// iOS-only properties

// A Boolean value that indicates whether route detection includes custom routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector/detectsCustomRoutes
func (r_ RouteDetector) DetectsCustomRoutes() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("detectsCustomRoutes"))
	return rv
}
func (r_ RouteDetector) SetDetectsCustomRoutes(value bool) {
	r_.ID.Send(objc.RegisterName("setDetectsCustomRoutes:"), value)
}





