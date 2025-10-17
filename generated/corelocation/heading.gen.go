// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Heading] class.
var headingClass = _HeadingClass{objc.GetClass("CLHeading")}

type _HeadingClass struct {
	class objc.Class
}

// The orientation of the user’s device, relative to true or magnetic north. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLHeading

type Heading struct {
	objectivec.Object
}

// HeadingFrom constructs a [Heading] from an unsafe.Pointer.
//
// The orientation of the user’s device, relative to true or magnetic north.
func HeadingFrom(ptr unsafe.Pointer) Heading {
	return Heading{objectivec.Object{objc.ID(ptr)}}
}



