// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SharingService] class.
var SharingServiceClass objc.Class

func init() {
	SharingServiceClass = objc.GetClass("NSSharingService")
}

type SharingService struct {
	objc.ID
}

func SharingServiceFrom(ptr unsafe.Pointer) SharingService {
	return SharingService{
		ID: objc.ID(ptr),
	}
}


// Returns a list of sharing services which could share all the provided items together. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSharingService/sharingServices(forItems:)
func (sc SharingService) SharingServicesForItems(items unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("sharingServicesForItems:")
	ret := objc.ID(SharingServiceClass).Send(sel, items)
	return unsafe.Pointer(ret)
}


