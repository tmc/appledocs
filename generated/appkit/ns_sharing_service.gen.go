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
	SharingServiceClass     _SharingServiceClass
	SharingServiceClassOnce sync.Once
)

func getSharingServiceClass() _SharingServiceClass {
	SharingServiceClassOnce.Do(func() {
		SharingServiceClass = _SharingServiceClass{objc.GetClass("NSSharingService")}
	})
	return SharingServiceClass
}

type _SharingServiceClass struct {
	class objc.Class
}

// An interface definition for the [SharingService] class.
type ISharingService interface {
	objectivec.IObject
}

// An object that facilitates the sharing of content with social media services, or with apps like Mail or Safari.
//
// An object provides a consistent user experience for sharing items— objects, objects, objects, video (through file URLs), of any object that implements the protocol—in macOS. For any item or group of items, the displays a sheet with the content to share. A sharing service can create a post on a social network like Twitter or Facebook, send a message by email or iMessage, upload videos to viewing services, or send a file using AirDrop. You can use objects directly in your app. The following example shows how to create a button that shares content directly to a social media service.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns a sharing service instance representing the specified service name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(named:)
func NewSharingServiceNamed(serviceName unsafe.Pointer) SharingService {
	rv := objc.Send[SharingService](objc.ID(getSharingServiceClass().class), objc.Sel("sharingServiceNamed:"), serviceName)
	return rv
}

// Creates a custom sharing service object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(title:image:alternateImage:handler:)
func NewSharingServiceWithTitleImageAlternateImageHandler(title string, image unsafe.Pointer, alternateImage unsafe.Pointer, block unsafe.Pointer) SharingService {
	instance := getSharingServiceClass().Alloc()
	rv := objc.Send[SharingService](instance.ID, objc.Sel("initWithTitle:image:alternateImage:handler:"), objc.String(title), image, alternateImage, block)
	rv.Autorelease()
	return rv
}


// Returns a sharing service instance representing the specified service name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/init(named:)
func (sc _SharingServiceClass) SharingServiceNamed(serviceName unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sharingServiceNamed:"), serviceName)
	return rv
}

// Returns a list of sharing services which could share all the provided items together.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/sharingServices(forItems:)
func (sc _SharingServiceClass) SharingServicesForItems(items objc.ID) []SharingService {
	rv := objc.Send[[]SharingService](objc.ID(sc.class), objc.Sel("sharingServicesForItems:"), items)
	return rv
}


