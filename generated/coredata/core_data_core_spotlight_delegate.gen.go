// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CoreDataCoreSpotlightDelegate] class.
var coreDataCoreSpotlightDelegateClass = _CoreDataCoreSpotlightDelegateClass{objc.GetClass("NSCoreDataCoreSpotlightDelegate")}

type _CoreDataCoreSpotlightDelegateClass struct {
	class objc.Class
}

// A set of methods that enable integration with Core Spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate

type CoreDataCoreSpotlightDelegate struct {
	objectivec.Object
}

// CoreDataCoreSpotlightDelegateFrom constructs a [CoreDataCoreSpotlightDelegate] from an unsafe.Pointer.
//
// A set of methods that enable integration with Core Spotlight.
func CoreDataCoreSpotlightDelegateFrom(ptr unsafe.Pointer) CoreDataCoreSpotlightDelegate {
	return CoreDataCoreSpotlightDelegate{objectivec.Object{objc.ID(ptr)}}
}



