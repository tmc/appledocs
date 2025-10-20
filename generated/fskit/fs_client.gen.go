// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSClient] class.
var (
	FSClientClass     _FSClientClass
	FSClientClassOnce sync.Once
)

func getFSClientClass() _FSClientClass {
	FSClientClassOnce.Do(func() {
		FSClientClass = _FSClientClass{objc.GetClass("FSClient")}
	})
	return FSClientClass
}

type _FSClientClass struct {
	class objc.Class
}

// An interface definition for the [FSClient] class.
type IFSClient interface {
	objectivec.IObject
	FetchInstalledExtensionsWithCompletionHandler(completionHandler unsafe.Pointer)
}

// An interface for apps and daemons to interact with FSKit.
//
// FSClient is the primary management interface for FSKit. Use this class to discover FSKit extensions installed on the system, including your own.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSClient
type FSClient struct {
	objectivec.Object
}

// FSClientFrom constructs a [FSClient] from an unsafe.Pointer.
//
// An interface for apps and daemons to interact with FSKit.
func FSClientFrom(ptr unsafe.Pointer) FSClient {
	return FSClient{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSClientClass) Alloc() FSClient {
	rv := objc.Send[FSClient](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSClientClass) New() FSClient {
	rv := objc.Send[FSClient](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSClient) Init() FSClient {
	rv := objc.Send[FSClient](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSClient) Autorelease() FSClient {
	rv := objc.Send[FSClient](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSClient creates a new FSClient instance.
func NewFSClient() FSClient {
	return getFSClientClass().New()
}


// Asynchronously retrieves an list of installed file system modules.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSClient/fetchInstalledExtensions(completionHandler:)
func (f_ FSClient) FetchInstalledExtensionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("fetchInstalledExtensionsWithCompletionHandler:"), completionHandler)
}



