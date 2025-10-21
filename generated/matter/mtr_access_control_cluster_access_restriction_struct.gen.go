// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterAccessRestrictionStruct] class.
var (
	MTRAccessControlClusterAccessRestrictionStructClass     _MTRAccessControlClusterAccessRestrictionStructClass
	MTRAccessControlClusterAccessRestrictionStructClassOnce sync.Once
)

func getMTRAccessControlClusterAccessRestrictionStructClass() _MTRAccessControlClusterAccessRestrictionStructClass {
	MTRAccessControlClusterAccessRestrictionStructClassOnce.Do(func() {
		MTRAccessControlClusterAccessRestrictionStructClass = _MTRAccessControlClusterAccessRestrictionStructClass{objc.GetClass("MTRAccessControlClusterAccessRestrictionStruct")}
	})
	return MTRAccessControlClusterAccessRestrictionStructClass
}

type _MTRAccessControlClusterAccessRestrictionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessRestrictionStruct] class.
type IMTRAccessControlClusterAccessRestrictionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionStruct
type MTRAccessControlClusterAccessRestrictionStruct struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessRestrictionStructFrom constructs a [MTRAccessControlClusterAccessRestrictionStruct] from an unsafe.Pointer.
func MTRAccessControlClusterAccessRestrictionStructFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessRestrictionStruct {
	return MTRAccessControlClusterAccessRestrictionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessRestrictionStructClass) Alloc() MTRAccessControlClusterAccessRestrictionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessRestrictionStructClass) New() MTRAccessControlClusterAccessRestrictionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessRestrictionStruct) Init() MTRAccessControlClusterAccessRestrictionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessRestrictionStruct) Autorelease() MTRAccessControlClusterAccessRestrictionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessRestrictionStruct creates a new MTRAccessControlClusterAccessRestrictionStruct instance.
func NewMTRAccessControlClusterAccessRestrictionStruct() MTRAccessControlClusterAccessRestrictionStruct {
	return getMTRAccessControlClusterAccessRestrictionStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionStruct/id
func (m_ MTRAccessControlClusterAccessRestrictionStruct) Id() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("id"))
	return rv
}


// SetId sets the value of the id property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionStruct/id
func (m_ MTRAccessControlClusterAccessRestrictionStruct) SetId(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionStruct/type
func (m_ MTRAccessControlClusterAccessRestrictionStruct) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionStruct/type
func (m_ MTRAccessControlClusterAccessRestrictionStruct) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}



