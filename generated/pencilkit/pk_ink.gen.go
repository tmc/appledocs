// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Ink] class.
var (
	InkClass     _InkClass
	InkClassOnce sync.Once
)

func getInkClass() _InkClass {
	InkClassOnce.Do(func() {
		InkClass = _InkClass{objc.GetClass("PKInk")}
	})
	return InkClass
}

type _InkClass struct {
	class objc.Class
}

// An interface definition for the [Ink] class.
type IInk interface {
	objectivec.IObject
}

// Provides a description of the creation and rendering of marks on a canvas.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkReference
type Ink struct {
	objectivec.Object
}

// InkFrom constructs a [Ink] from an unsafe.Pointer.
//
// Provides a description of the creation and rendering of marks on a canvas.
func InkFrom(ptr unsafe.Pointer) Ink {
	return Ink{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _InkClass) Alloc() Ink {
	rv := objc.Send[Ink](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InkClass) New() Ink {
	rv := objc.Send[Ink](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ Ink) Init() Ink {
	rv := objc.Send[Ink](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ Ink) Autorelease() Ink {
	rv := objc.Send[Ink](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInk creates a new Ink instance.
func NewInk() Ink {
	return getInkClass().New()
}




// Create a new ink, specifying its type, color.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkReference/init(inkType:color:)
func NewInkWithInkTypeColor(type_ unsafe.Pointer, color appkit.IColor) Ink {
	instance := getInkClass().Alloc()
	rv := objc.Send[Ink](instance.ID, objc.Sel("initWithInkType:color:"), type_, color)
	rv.Autorelease()
	return rv
}


// The base color for this ink.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkReference/color
func (i_ Ink) Color() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("color"))
	return rv
}

// The type of ink, such as pen or pencil, as defined in the enumeration.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkReference/inkType
func (i_ Ink) InkType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("inkType"))
	return rv
}

// The version of PencilKit necessary to use the ink.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkReference/requiredContentVersion
func (i_ Ink) RequiredContentVersion() ContentVersion {
	rv := objc.Send[ContentVersion](i_.ID, objc.Sel("requiredContentVersion"))
	return rv
}


