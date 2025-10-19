// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FilterShape] class.
var (
	filterShapeClass     _FilterShapeClass
	filterShapeClassOnce sync.Once
)

func getFilterShapeClass() _FilterShapeClass {
	filterShapeClassOnce.Do(func() {
		filterShapeClass = _FilterShapeClass{objc.GetClass("CIFilterShape")}
	})
	return filterShapeClass
}

type _FilterShapeClass struct {
	class objc.Class
}

// An interface definition for the [FilterShape] class.
type IFilterShape interface {
	objectivec.IObject
	InsetByXY(dx int, dy int) unsafe.Pointer
	IntersectWithRect(r unsafe.Pointer) unsafe.Pointer
	IntersectWith(s2 unsafe.Pointer) unsafe.Pointer
	TransformByInterior(m unsafe.Pointer, flag bool) unsafe.Pointer
	UnionWith(s2 unsafe.Pointer) unsafe.Pointer
	UnionWithRect(r unsafe.Pointer) unsafe.Pointer
}

// A description of the bounding shape of a filter and the domain of definition for a filter operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape
type FilterShape struct {
	objectivec.Object
}

// FilterShapeFrom constructs a [FilterShape] from an unsafe.Pointer.
//
// A description of the bounding shape of a filter and the domain of definition for a filter operation.
func FilterShapeFrom(ptr unsafe.Pointer) FilterShape {
	return FilterShape{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FilterShapeClass) Alloc() FilterShape {
	rv := objc.Send[FilterShape](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FilterShapeClass) New() FilterShape {
	rv := objc.Send[FilterShape](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilterShape) Init() FilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilterShape) Autorelease() FilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilterShape creates a new FilterShape instance.
func NewFilterShape() FilterShape {
	return getFilterShapeClass().New()
}


// Initializes a filter shape object with a rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/init(rect:)
func NewFilterShapeWithRect(r unsafe.Pointer) FilterShape {
	instance := getFilterShapeClass().Alloc()
	rv := objc.Send[FilterShape](instance.ID, objc.Sel("initWithRect:"), r)
	rv.Autorelease()
	return rv
}


// Creates a filter shape object and initializes it with a rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/shapeWithRect:
func (fc _FilterShapeClass) ShapeWithRect(r unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("shapeWithRect:"), r)
	return rv
}
// Modifies a filter shape object so that it is inset by the specified x and y values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/insetBy(x:y:)
func (f_ FilterShape) InsetByXY(dx int, dy int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("insetByX:Y:"), dx, dy)
	return rv
}
// Creates a filter shape that represents the intersection of the current filter shape and a rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/intersect(with:)-2o2n8
func (f_ FilterShape) IntersectWithRect(r unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("intersectWithRect:"), r)
	return rv
}
// Creates a filter shape object that represents the intersection of the current filter shape and the specified filter shape object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/intersect(with:)-8iw
func (f_ FilterShape) IntersectWith(s2 unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("intersectWith:"), s2)
	return rv
}
// Creates a filter shape that results from applying a transform to the current filter shape.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/transform(by:interior:)
func (f_ FilterShape) TransformByInterior(m unsafe.Pointer, flag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("transformBy:interior:"), m, flag)
	return rv
}
// Creates a filter shape that results from the union of the current filter shape and another filter shape object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/union(with:)-52mnd
func (f_ FilterShape) UnionWith(s2 unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("unionWith:"), s2)
	return rv
}
// Creates a filter shape that results from the union of the current filter shape and a rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/union(with:)-75ebo
func (f_ FilterShape) UnionWithRect(r unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("unionWithRect:"), r)
	return rv
}

