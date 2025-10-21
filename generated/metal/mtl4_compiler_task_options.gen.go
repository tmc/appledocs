// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTL4CompilerTaskOptions] class.
var (
	MTL4CompilerTaskOptionsClass     _MTL4CompilerTaskOptionsClass
	MTL4CompilerTaskOptionsClassOnce sync.Once
)

func getMTL4CompilerTaskOptionsClass() _MTL4CompilerTaskOptionsClass {
	MTL4CompilerTaskOptionsClassOnce.Do(func() {
		MTL4CompilerTaskOptionsClass = _MTL4CompilerTaskOptionsClass{objc.GetClass("MTL4CompilerTaskOptions")}
	})
	return MTL4CompilerTaskOptionsClass
}

type _MTL4CompilerTaskOptionsClass struct {
	class objc.Class
}

// An interface definition for the [MTL4CompilerTaskOptions] class.
type IMTL4CompilerTaskOptions interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskOptions
type MTL4CompilerTaskOptions struct {
	objectivec.Object
}

// MTL4CompilerTaskOptionsFrom constructs a [MTL4CompilerTaskOptions] from an unsafe.Pointer.
func MTL4CompilerTaskOptionsFrom(ptr unsafe.Pointer) MTL4CompilerTaskOptions {
	return MTL4CompilerTaskOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4CompilerTaskOptionsClass) Alloc() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4CompilerTaskOptionsClass) New() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CompilerTaskOptions) Init() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CompilerTaskOptions) Autorelease() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CompilerTaskOptions creates a new MTL4CompilerTaskOptions instance.
func NewMTL4CompilerTaskOptions() MTL4CompilerTaskOptions {
	return getMTL4CompilerTaskOptionsClass().New()
}


// Specifies a set of archive instances this compilation process uses for accelerating the build process.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskOptions/lookupArchives
func (m_ MTL4CompilerTaskOptions) LookupArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("lookupArchives"))
	return rv
}


// SetLookupArchives sets the value of the lookupArchives property.
// Specifies a set of archive instances this compilation process uses for accelerating the build process.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskOptions/lookupArchives
func (m_ MTL4CompilerTaskOptions) SetLookupArchives(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setLookupArchives:"), nsArray)
}


