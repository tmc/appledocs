// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SharingService] class.
var (
	sharingServiceClass     _SharingServiceClass
	sharingServiceClassOnce sync.Once
)

func getSharingServiceClass() _SharingServiceClass {
	sharingServiceClassOnce.Do(func() {
		sharingServiceClass = _SharingServiceClass{objc.GetClass("NSSharingService")}
	})
	return sharingServiceClass
}

type _SharingServiceClass struct {
	class objc.Class
}

// An interface definition for the [SharingService] class.
type ISharingService interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (sc _SharingServiceClass) Alloc() SharingService {
	rv := objc.Send[SharingService](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SharingServiceClass) New() SharingService {
	rv := objc.Send[SharingService](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharingService) Init() SharingService {
	rv := objc.Send[SharingService](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharingService) Autorelease() SharingService {
	rv := objc.Send[SharingService](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharingService creates a new SharingService instance.
func NewSharingService() SharingService {
	return getSharingServiceClass().New()
}


// Returns a list of sharing services which could share all the provided items together. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/sharingServices(forItems:)
func (sc _SharingServiceClass) SharingServicesForItems(items unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sharingServicesForItems:"), items)
	return rv
}


