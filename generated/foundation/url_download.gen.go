// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLDownload] class.
var (
	uRLDownloadClass     _URLDownloadClass
	uRLDownloadClassOnce sync.Once
)

func getURLDownloadClass() _URLDownloadClass {
	uRLDownloadClassOnce.Do(func() {
		uRLDownloadClass = _URLDownloadClass{objc.GetClass("NSURLDownload")}
	})
	return uRLDownloadClass
}

type _URLDownloadClass struct {
	class objc.Class
}

// An interface definition for the [URLDownload] class.
type IURLDownload interface {
	objectivec.IObject
}

// An object that downloads a resource asynchronously and saves the data to a file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload
type URLDownload struct {
	objectivec.Object
}

// URLDownloadFrom constructs a [URLDownload] from an unsafe.Pointer.
//
// An object that downloads a resource asynchronously and saves the data to a file.
func URLDownloadFrom(ptr unsafe.Pointer) URLDownload {
	return URLDownload{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLDownloadClass) Alloc() URLDownload {
	rv := objc.Send[URLDownload](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLDownloadClass) New() URLDownload {
	rv := objc.Send[URLDownload](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLDownload) Init() URLDownload {
	rv := objc.Send[URLDownload](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLDownload) Autorelease() URLDownload {
	rv := objc.Send[URLDownload](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLDownload creates a new URLDownload instance.
func NewURLDownload() URLDownload {
	return getURLDownloadClass().New()
}




