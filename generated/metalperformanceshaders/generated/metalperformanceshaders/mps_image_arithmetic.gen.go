// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageArithmetic */


/* debug [class_header]: Header for MPSImageArithmetic */
// The class instance for the [ImageArithmetic] class.
var (
	ImageArithmeticClass     _ImageArithmeticClass
	ImageArithmeticClassOnce sync.Once
)

func getImageArithmeticClass() _ImageArithmeticClass {
	ImageArithmeticClassOnce.Do(func() {
		ImageArithmeticClass = _ImageArithmeticClass{objc.GetClass("MPSImageArithmetic")}
	})
	return ImageArithmeticClass
}

type _ImageArithmeticClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageArithmetic */
// An interface definition for the [ImageArithmetic] class.
type IImageArithmetic interface {
	IBinaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageArithmetic */
	// properties:
	SecondaryScale() objectivec.IObject
	SetSecondaryScale(value objectivec.IObject)
	PrimaryScale() objectivec.IObject
	SetPrimaryScale(value objectivec.IObject)
	Bias() objectivec.IObject
	SetBias(value objectivec.IObject)
	PrimaryStrideInPixels() Size get set /* not a class type */
	SetPrimaryStrideInPixels(value Size get set /* not a class type */)
	SecondaryStrideInPixels() Size get set /* not a class type */
	SetSecondaryStrideInPixels(value Size get set /* not a class type */)
	MaximumValue() objectivec.IObject
	SetMaximumValue(value objectivec.IObject)
	MinimumValue() objectivec.IObject
	SetMinimumValue(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageArithmetic */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageArithmetic */
// Alloc allocates a new instance without initialization.
func (ic _ImageArithmeticClass) Alloc() ImageArithmetic {
	rv := objc.Send[ImageArithmetic](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageArithmeticClass) New() ImageArithmetic {
	rv := objc.Send[ImageArithmetic](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageArithmetic) Init() ImageArithmetic {
	rv := objc.Send[ImageArithmetic](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageArithmetic) Autorelease() ImageArithmetic {
	rv := objc.Send[ImageArithmetic](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageArithmetic creates a new ImageArithmetic instance.
func NewImageArithmetic() ImageArithmetic {
	return getImageArithmeticClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageArithmetic */
// Base class for basic arithmetic nodes


// Base class for basic arithmetic nodes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageArithmetic
type ImageArithmetic struct {
	BinaryImageKernel
}

// ImageArithmeticFrom constructs a [ImageArithmetic] from an unsafe.Pointer.
//
// Base class for basic arithmetic nodes
func ImageArithmeticFrom(ptr unsafe.Pointer) ImageArithmetic {
	return ImageArithmetic{
		BinaryImageKernel: BinaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageArithmetic *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageArithmetic */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageArithmetic */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageArithmetic */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageArithmetic */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866601-secondaryscale
func (i_ ImageArithmetic) SecondaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("secondaryScale"))
	return rv
}/* debug [instance_properties/getter]: secondaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866601-secondaryscale
func (i_ ImageArithmetic) SetSecondaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSecondaryScale:"), value)
}/* debug [instance_properties/setter]: secondaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866602-primaryscale
func (i_ ImageArithmetic) PrimaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("primaryScale"))
	return rv
}/* debug [instance_properties/getter]: primaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866602-primaryscale
func (i_ ImageArithmetic) SetPrimaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPrimaryScale:"), value)
}/* debug [instance_properties/setter]: primaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866609-bias
func (i_ ImageArithmetic) Bias() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("bias"))
	return rv
}/* debug [instance_properties/getter]: bias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866609-bias
func (i_ ImageArithmetic) SetBias(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBias:"), value)
}/* debug [instance_properties/setter]: bias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2889864-primarystrideinpixels
func (i_ ImageArithmetic) PrimaryStrideInPixels() Size get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("primaryStrideInPixels"))
	return rv
}/* debug [instance_properties/getter]: primaryStrideInPixels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2889864-primarystrideinpixels
func (i_ ImageArithmetic) SetPrimaryStrideInPixels(value Size get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPrimaryStrideInPixels:"), value)
}/* debug [instance_properties/setter]: primaryStrideInPixels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2889865-secondarystrideinpixels
func (i_ ImageArithmetic) SecondaryStrideInPixels() Size get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("secondaryStrideInPixels"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrideInPixels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2889865-secondarystrideinpixels
func (i_ ImageArithmetic) SetSecondaryStrideInPixels(value Size get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSecondaryStrideInPixels:"), value)
}/* debug [instance_properties/setter]: secondaryStrideInPixels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2942356-maximumvalue
func (i_ ImageArithmetic) MaximumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("maximumValue"))
	return rv
}/* debug [instance_properties/getter]: maximumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2942356-maximumvalue
func (i_ ImageArithmetic) SetMaximumValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaximumValue:"), value)
}/* debug [instance_properties/setter]: maximumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2942357-minimumvalue
func (i_ ImageArithmetic) MinimumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("minimumValue"))
	return rv
}/* debug [instance_properties/getter]: minimumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2942357-minimumvalue
func (i_ ImageArithmetic) SetMinimumValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMinimumValue:"), value)
}/* debug [instance_properties/setter]: minimumValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageArithmetic */



