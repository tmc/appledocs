// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHCloudIdentifier] class.
var (
	PHCloudIdentifierClass     _PHCloudIdentifierClass
	PHCloudIdentifierClassOnce sync.Once
)

func getPHCloudIdentifierClass() _PHCloudIdentifierClass {
	PHCloudIdentifierClassOnce.Do(func() {
		PHCloudIdentifierClass = _PHCloudIdentifierClass{objc.GetClass("PHCloudIdentifier")}
	})
	return PHCloudIdentifierClass
}

type _PHCloudIdentifierClass struct {
	class objc.Class
}

// An interface definition for the [PHCloudIdentifier] class.
type IPHCloudIdentifier interface {
	objectivec.IObject
}

// An object that identifies an asset or collection that syncs through iCloud Photos.
//
// A cloud identifier is a type of identifier that behaves like a local identifier. Use cloud identifiers to identify objects that sync across devices through iCloud Photos. You can store, sync, and use cloud identifiers with devices synced with an iCloud account. You’re also able to use secure coding to encode and decode cloud identifiers. A local identifier is valid for referring to objects only in the context of a local device. These objects include , , and . Because a cloud identifier is universal, you can use it on any iCloud-synced device. Convert the cloud identifier back to a local identifier and perform a fetch to find the equivalent object on that device. Perform batch lookups of identifiers using and . Retrieving identifier mappings can be an expensive operation, so perform lookups sparingly. If a lookup fails, inspect the error property on or for details. See for additional error details.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCloudIdentifier
type PHCloudIdentifier struct {
	objectivec.Object
}

// PHCloudIdentifierFrom constructs a [PHCloudIdentifier] from an unsafe.Pointer.
//
// An object that identifies an asset or collection that syncs through iCloud Photos.
func PHCloudIdentifierFrom(ptr unsafe.Pointer) PHCloudIdentifier {
	return PHCloudIdentifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHCloudIdentifierClass) Alloc() PHCloudIdentifier {
	rv := objc.Send[PHCloudIdentifier](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHCloudIdentifierClass) New() PHCloudIdentifier {
	rv := objc.Send[PHCloudIdentifier](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHCloudIdentifier) Init() PHCloudIdentifier {
	rv := objc.Send[PHCloudIdentifier](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHCloudIdentifier) Autorelease() PHCloudIdentifier {
	rv := objc.Send[PHCloudIdentifier](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHCloudIdentifier creates a new PHCloudIdentifier instance.
func NewPHCloudIdentifier() PHCloudIdentifier {
	return getPHCloudIdentifierClass().New()
}




