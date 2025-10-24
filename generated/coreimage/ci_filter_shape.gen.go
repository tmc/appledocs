// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIFilterShape */


/* debug [class_header]: Header for CIFilterShape */
// The class instance for the [FilterShape] class.
var (
	FilterShapeClass     _FilterShapeClass
	FilterShapeClassOnce sync.Once
)

func getFilterShapeClass() _FilterShapeClass {
	FilterShapeClassOnce.Do(func() {
		FilterShapeClass = _FilterShapeClass{objc.GetClass("CIFilterShape")}
	})
	return FilterShapeClass
}

type _FilterShapeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FilterShape */
// An interface definition for the [FilterShape] class.
type IFilterShape interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FilterShape */
	// properties:
	Extent() corefoundation.CGRect
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FilterShape */
	// methods:
	InsetByXY(dx int, dy int) IFilterShape
	IntersectWithRect(r corefoundation.CGRect) IFilterShape
	IntersectWith(s2 ICIFilterShape) IFilterShape
	TransformByInterior(m corefoundation.CGAffineTransform, flag bool) IFilterShape
	UnionWith(s2 ICIFilterShape) IFilterShape
	UnionWithRect(r corefoundation.CGRect) IFilterShape
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FilterShape */
// Alloc allocates a new instance without initialization.
func (fc _FilterShapeClass) Alloc() FilterShape {
	rv := objc.Send[FilterShape](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FilterShape */
// A description of the bounding shape of a filter and the domain of definition for a filter operation.
//
// You use objects in conjunction with Core Image classes, such as , , and , to create custom filters.


// A description of the bounding shape of a filter and the domain of definition for a filter operation.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FilterShape */

// Initializes a filter shape object with a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/init(rect:)
func NewFilterShapeWithRect(r corefoundation.CGRect) FilterShape {
	instance := getFilterShapeClass().Alloc()
	rv := objc.Send[FilterShape](instance.ID, objc.Sel("initWithRect:"), r)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFilterShapeWithRect */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FilterShape */

// Creates a filter shape object and initializes it with a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/shapeWithRect:
func (fc _FilterShapeClass) ShapeWithRect(r corefoundation.CGRect) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("shapeWithRect:"), r)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ShapeWithRect) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FilterShape */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FilterShape */

// Modifies a filter shape object so that it is inset by the specified x and y values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/insetBy(x:y:)
func (f_ FilterShape) InsetByXY(dx int, dy int) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("insetByX:Y:"), dx, dy)
	return rv
}/* debug [instance_methods/method]: InsetByXY */


// Creates a filter shape that represents the intersection of the current filter shape and a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/intersect(with:)-2o2n8
func (f_ FilterShape) IntersectWithRect(r corefoundation.CGRect) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("intersectWithRect:"), r)
	return rv
}/* debug [instance_methods/method]: IntersectWithRect */


// Creates a filter shape object that represents the intersection of the current filter shape and the specified filter shape object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/intersect(with:)-8iw
func (f_ FilterShape) IntersectWith(s2 ICIFilterShape) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("intersectWith:"), s2)
	return rv
}/* debug [instance_methods/method]: IntersectWith */


// Creates a filter shape that results from applying a transform to the current filter shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/transform(by:interior:)
func (f_ FilterShape) TransformByInterior(m corefoundation.CGAffineTransform, flag bool) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("transformBy:interior:"), m, flag)
	return rv
}/* debug [instance_methods/method]: TransformByInterior */


// Creates a filter shape that results from the union of the current filter shape and another filter shape object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/union(with:)-52mnd
func (f_ FilterShape) UnionWith(s2 ICIFilterShape) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("unionWith:"), s2)
	return rv
}/* debug [instance_methods/method]: UnionWith */


// Creates a filter shape that results from the union of the current filter shape and a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/union(with:)-75ebo
func (f_ FilterShape) UnionWithRect(r corefoundation.CGRect) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("unionWithRect:"), r)
	return rv
}/* debug [instance_methods/method]: UnionWithRect */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FilterShape */

// The extent of the filter shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/extent
func (f_ FilterShape) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](f_.ID, objc.Sel("extent"))
	return rv
}/* debug [instance_properties/getter]: extent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIFilterShape */


