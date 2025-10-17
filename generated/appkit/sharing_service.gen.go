// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SharingService] class.
var sharingServiceClass = _SharingServiceClass{objc.GetClass("NSSharingService")}

type _SharingServiceClass struct {
	class objc.Class
}

// An object that facilitates the sharing of content with social media services, or with apps like Mail or Safari. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService

type SharingService struct {
	objectivec.Object
}

// SharingServiceFrom constructs a [SharingService] from an unsafe.Pointer.
//
// An object that facilitates the sharing of content with social media services, or with apps like Mail or Safari.
func SharingServiceFrom(ptr unsafe.Pointer) SharingService {
	return SharingService{objectivec.Object{objc.ID(ptr)}}
}

// Returns a list of sharing services which could share all the provided items together. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/sharingServices(forItems:)
func (sc _SharingServiceClass) SharingServicesForItems(items unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sharingServicesForItems:"), items)
	return rv
}


