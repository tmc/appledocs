
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SharingService] class.
var SharingServiceClass _SharingServiceClass

func init() {
	SharingServiceClass = _SharingServiceClass{objc.GetClass("NSSharingService")}
}

type _SharingServiceClass struct {
	objc.Class
}

// An interface definition for the [SharingService] class.
type ISharingService interface {
	ID() objc.ID
}

type SharingService struct {
	id objc.ID
}

func SharingServiceFrom(ptr unsafe.Pointer) SharingService {
	return SharingService{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SharingService) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SharingServiceClass) Alloc() SharingService {
	rv := objc.Send[SharingService](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SharingServiceClass) New() SharingService {
	rv := objc.Send[SharingService](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSharingService creates and returns a new initialized instance.
func NewSharingService() SharingService {
	return SharingServiceClass.New()
}

// Init initializes the instance.
func (s_ SharingService) Init() SharingService {
	rv := objc.Send[SharingService](s_.ID(), selInit)
	return rv
}
// Returns a list of sharing services which could share all the provided items together. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSharingService/sharingServices(forItems:)
func (sc _SharingServiceClass) SharingServicesForItems(items unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.Class), objc.RegisterName("sharingServicesForItems:"), items)
	return rv
}

// SharingService_SharingServicesForItems creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSharingService/sharingServices(forItems:)
func SharingService_SharingServicesForItems(items unsafe.Pointer) unsafe.Pointer {
	return SharingServiceClass.SharingServicesForItems(items)
}
