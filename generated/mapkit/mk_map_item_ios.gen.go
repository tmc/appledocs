//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/gameplaykit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKMapItem


// Opens the Maps app from a particular scene using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openInMaps(launchOptions:from:completionHandler:)
func (m_ MKMapItem) OpenInMapsWithLaunchOptionsFromSceneCompletionHandler(launchOptions foundation.IDictionary, scene gameplaykit.Scene, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("openInMapsWithLaunchOptions:fromScene:completionHandler:"), launchOptions, scene, completion)
}

// iOS-only properties




