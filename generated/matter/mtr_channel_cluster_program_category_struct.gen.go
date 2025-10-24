// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterProgramCategoryStruct] class.
var (
	MTRChannelClusterProgramCategoryStructClass     _MTRChannelClusterProgramCategoryStructClass
	MTRChannelClusterProgramCategoryStructClassOnce sync.Once
)

func getMTRChannelClusterProgramCategoryStructClass() _MTRChannelClusterProgramCategoryStructClass {
	MTRChannelClusterProgramCategoryStructClassOnce.Do(func() {
		MTRChannelClusterProgramCategoryStructClass = _MTRChannelClusterProgramCategoryStructClass{objc.GetClass("MTRChannelClusterProgramCategoryStruct")}
	})
	return MTRChannelClusterProgramCategoryStructClass
}

type _MTRChannelClusterProgramCategoryStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterProgramCategoryStruct] class.
type IMTRChannelClusterProgramCategoryStruct interface {
	objectivec.IObject
	// properties:
	Category() objc.IObject /* cross-framework: NSString */
	SetCategory(value objc.IObject /* cross-framework: NSString */)
	SubCategory() objc.IObject /* cross-framework: NSString */
	SetSubCategory(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct
type MTRChannelClusterProgramCategoryStruct struct {
	objectivec.Object
}

// MTRChannelClusterProgramCategoryStructFrom constructs a [MTRChannelClusterProgramCategoryStruct] from an unsafe.Pointer.
func MTRChannelClusterProgramCategoryStructFrom(ptr unsafe.Pointer) MTRChannelClusterProgramCategoryStruct {
	return MTRChannelClusterProgramCategoryStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterProgramCategoryStructClass) Alloc() MTRChannelClusterProgramCategoryStruct {
	rv := objc.Send[MTRChannelClusterProgramCategoryStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterProgramCategoryStructClass) New() MTRChannelClusterProgramCategoryStruct {
	rv := objc.Send[MTRChannelClusterProgramCategoryStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterProgramCategoryStruct) Init() MTRChannelClusterProgramCategoryStruct {
	rv := objc.Send[MTRChannelClusterProgramCategoryStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterProgramCategoryStruct) Autorelease() MTRChannelClusterProgramCategoryStruct {
	rv := objc.Send[MTRChannelClusterProgramCategoryStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterProgramCategoryStruct creates a new MTRChannelClusterProgramCategoryStruct instance.
func NewMTRChannelClusterProgramCategoryStruct() MTRChannelClusterProgramCategoryStruct {
	return getMTRChannelClusterProgramCategoryStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct/category
func (m_ MTRChannelClusterProgramCategoryStruct) Category() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("category"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct/category
func (m_ MTRChannelClusterProgramCategoryStruct) SetCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCategory:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct/subCategory
func (m_ MTRChannelClusterProgramCategoryStruct) SubCategory() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subCategory"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct/subCategory
func (m_ MTRChannelClusterProgramCategoryStruct) SetSubCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubCategory:"), value)
}



