// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FSContainerIdentifier] class.
var (
	FSContainerIdentifierClass     _FSContainerIdentifierClass
	FSContainerIdentifierClassOnce sync.Once
)

func getFSContainerIdentifierClass() _FSContainerIdentifierClass {
	FSContainerIdentifierClassOnce.Do(func() {
		FSContainerIdentifierClass = _FSContainerIdentifierClass{objc.GetClass("FSContainerIdentifier")}
	})
	return FSContainerIdentifierClass
}

type _FSContainerIdentifierClass struct {
	class objc.Class
}

// An interface definition for the [FSContainerIdentifier] class.
type IFSContainerIdentifier interface {
	IFSEntityIdentifier
	// properties:
	VolumeIdentifier() objc.IObject /* cross-framework: FSVolumeIdentifier */
	// methods:
}

// A type that identifies a container.
//
// The identifier is either a UUID or a UUID with additional differentiating bytes. Some network protocols evaluate access based on a user ID when connecting. In this situation, when a file server receives multiple client connections with different user IDs, the server provides different file hierarchies to each. For such systems, represent the container identifier as the UUID associated with the server, followed by four or eight bytes to differentiate connections.


// A type that identifies a container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerIdentifier
type FSContainerIdentifier struct {
	FSEntityIdentifier
}

// FSContainerIdentifierFrom constructs a [FSContainerIdentifier] from an unsafe.Pointer.
//
// A type that identifies a container.
func FSContainerIdentifierFrom(ptr unsafe.Pointer) FSContainerIdentifier {
	return FSContainerIdentifier{
		FSEntityIdentifier: FSEntityIdentifierFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FSContainerIdentifierClass) Alloc() FSContainerIdentifier {
	rv := objc.Send[FSContainerIdentifier](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSContainerIdentifierClass) New() FSContainerIdentifier {
	rv := objc.Send[FSContainerIdentifier](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSContainerIdentifier) Init() FSContainerIdentifier {
	rv := objc.Send[FSContainerIdentifier](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSContainerIdentifier) Autorelease() FSContainerIdentifier {
	rv := objc.Send[FSContainerIdentifier](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSContainerIdentifier creates a new FSContainerIdentifier instance.
func NewFSContainerIdentifier() FSContainerIdentifier {
	return getFSContainerIdentifierClass().New()
}



// The volume identifier associated with the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerIdentifier/volumeIdentifier
func (f_ FSContainerIdentifier) VolumeIdentifier() objc.IObject /* cross-framework: FSVolumeIdentifier */ {
	rv := objc.Send[FSVolumeIdentifier](f_.ID, objc.Sel("volumeIdentifier"))
	return rv
}



