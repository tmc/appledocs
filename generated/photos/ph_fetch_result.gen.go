// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHFetchResult] class.
var (
	PHFetchResultClass     _PHFetchResultClass
	PHFetchResultClassOnce sync.Once
)

func getPHFetchResultClass() _PHFetchResultClass {
	PHFetchResultClassOnce.Do(func() {
		PHFetchResultClass = _PHFetchResultClass{objc.GetClass("PHFetchResult")}
	})
	return PHFetchResultClass
}

type _PHFetchResultClass struct {
	class objc.Class
}

// An interface definition for the [PHFetchResult] class.
type IPHFetchResult interface {
	objectivec.IObject
	// properties:
	Count() int
	SetCount(value int)
	FirstObject() unsafe.Pointer
	SetFirstObject(value unsafe.Pointer)
	LastObject() unsafe.Pointer
	SetLastObject(value unsafe.Pointer)
	LocalIdentifier() objc.IObject /* cross-framework: NSString */
	SetLocalIdentifier(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An ordered list of assets or collections returned from a Photos fetch method.
//
// When you use class methods on the , , , and classes to retrieve objects, Photos provides the resulting objects in a fetch result. You access the contents of a fetch result with the same methods and conventions used by the class. Unlike an object, however, a object dynamically loads its contents from the Photos library as needed, providing optimal performance even when handling a large number of results. A fetch result provides thread-safe access to its contents. After a fetch, the fetch result’s value is constant, and all objects in the fetch result keep the same value. (To get updated content for a fetch, register a change observer with the shared object.) A fetch result caches its contents, keeping a batch of objects around the most recently accessed index. Because objects outside of the batch are no longer cached, accessing these objects results in refetching those objects. This process can result in changes to values previously read from those objects.

// An ordered list of assets or collections returned from a Photos fetch method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchResult
type PHFetchResult struct {
	objectivec.Object
}

// PHFetchResultFrom constructs a [PHFetchResult] from an unsafe.Pointer.
//
// An ordered list of assets or collections returned from a Photos fetch method.
func PHFetchResultFrom(ptr unsafe.Pointer) PHFetchResult {
	return PHFetchResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHFetchResultClass) Alloc() PHFetchResult {
	rv := objc.Send[PHFetchResult](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHFetchResultClass) New() PHFetchResult {
	rv := objc.Send[PHFetchResult](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHFetchResult) Init() PHFetchResult {
	rv := objc.Send[PHFetchResult](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHFetchResult) Autorelease() PHFetchResult {
	rv := objc.Send[PHFetchResult](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHFetchResult creates a new PHFetchResult instance.
func NewPHFetchResult() PHFetchResult {
	return getPHFetchResultClass().New()
}

// The number of objects in the fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresult/count
func (p_ PHFetchResult) Count() int {
	rv := objc.Send[int](p_.ID, objc.Sel("count"))
	return rv
}

// The number of objects in the fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresult/count
func (p_ PHFetchResult) SetCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCount:"), value)
}

// The first object in the fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresult/firstobject
func (p_ PHFetchResult) FirstObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("firstObject"))
	return rv
}

// The first object in the fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresult/firstobject
func (p_ PHFetchResult) SetFirstObject(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFirstObject:"), value)
}

// The last object in the fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresult/lastobject
func (p_ PHFetchResult) LastObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("lastObject"))
	return rv
}

// The last object in the fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresult/lastobject
func (p_ PHFetchResult) SetLastObject(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLastObject:"), value)
}

// A unique string that persistently identifies the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phobject/localidentifier
func (p_ PHFetchResult) LocalIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localIdentifier"))
	return rv
}

// A unique string that persistently identifies the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phobject/localidentifier
func (p_ PHFetchResult) SetLocalIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalIdentifier:"), value)
}
