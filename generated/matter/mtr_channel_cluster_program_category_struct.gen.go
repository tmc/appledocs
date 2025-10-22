// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Category() string
	SetCategory(value string)
	SubCategory() string
	SetSubCategory(value string)
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct/category
func (m_ MTRChannelClusterProgramCategoryStruct) Category() string {
	rv := objc.Send[string](m_.ID, objc.Sel("category"))
	return rv
}


// SetCategory sets the value of the category property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct/category
func (m_ MTRChannelClusterProgramCategoryStruct) SetCategory(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCategory:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct/subCategory
func (m_ MTRChannelClusterProgramCategoryStruct) SubCategory() string {
	rv := objc.Send[string](m_.ID, objc.Sel("subCategory"))
	return rv
}


// SetSubCategory sets the value of the subCategory property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct/subCategory
func (m_ MTRChannelClusterProgramCategoryStruct) SetSubCategory(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubCategory:"), objc.String(value))
}



