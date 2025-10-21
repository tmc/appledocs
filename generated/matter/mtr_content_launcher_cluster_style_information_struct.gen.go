// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentLauncherClusterStyleInformationStruct] class.
var (
	MTRContentLauncherClusterStyleInformationStructClass     _MTRContentLauncherClusterStyleInformationStructClass
	MTRContentLauncherClusterStyleInformationStructClassOnce sync.Once
)

func getMTRContentLauncherClusterStyleInformationStructClass() _MTRContentLauncherClusterStyleInformationStructClass {
	MTRContentLauncherClusterStyleInformationStructClassOnce.Do(func() {
		MTRContentLauncherClusterStyleInformationStructClass = _MTRContentLauncherClusterStyleInformationStructClass{objc.GetClass("MTRContentLauncherClusterStyleInformationStruct")}
	})
	return MTRContentLauncherClusterStyleInformationStructClass
}

type _MTRContentLauncherClusterStyleInformationStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterStyleInformationStruct] class.
type IMTRContentLauncherClusterStyleInformationStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterStyleInformationStruct
type MTRContentLauncherClusterStyleInformationStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterStyleInformationStructFrom constructs a [MTRContentLauncherClusterStyleInformationStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterStyleInformationStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterStyleInformationStruct {
	return MTRContentLauncherClusterStyleInformationStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterStyleInformationStructClass) Alloc() MTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterStyleInformationStructClass) New() MTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterStyleInformationStruct) Init() MTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterStyleInformationStruct) Autorelease() MTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterStyleInformationStruct creates a new MTRContentLauncherClusterStyleInformationStruct instance.
func NewMTRContentLauncherClusterStyleInformationStruct() MTRContentLauncherClusterStyleInformationStruct {
	return getMTRContentLauncherClusterStyleInformationStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformationstruct/color
func (m_ MTRContentLauncherClusterStyleInformationStruct) Color() string {
	rv := objc.Send[string](m_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformationstruct/color
func (m_ MTRContentLauncherClusterStyleInformationStruct) SetColor(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColor:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformationstruct/imageurl-8k1s1
func (m_ MTRContentLauncherClusterStyleInformationStruct) ImageURL() string {
	rv := objc.Send[string](m_.ID, objc.Sel("imageURL"))
	return rv
}


// SetImageURL sets the value of the imageURL property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformationstruct/imageurl-8k1s1
func (m_ MTRContentLauncherClusterStyleInformationStruct) SetImageURL(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageURL:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformationstruct/imageurl-8jrvl
func (m_ MTRContentLauncherClusterStyleInformationStruct) ImageUrl() string {
	rv := objc.Send[string](m_.ID, objc.Sel("imageUrl"))
	return rv
}


// SetImageUrl sets the value of the imageUrl property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformationstruct/imageurl-8jrvl
func (m_ MTRContentLauncherClusterStyleInformationStruct) SetImageUrl(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageUrl:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformationstruct/size
func (m_ MTRContentLauncherClusterStyleInformationStruct) Size() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("size"))
	return rv
}


// SetSize sets the value of the size property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformationstruct/size
func (m_ MTRContentLauncherClusterStyleInformationStruct) SetSize(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSize:"), value)
}



