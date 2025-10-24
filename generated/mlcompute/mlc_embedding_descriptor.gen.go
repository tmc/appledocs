// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCEmbeddingDescriptor */


/* debug [class_header]: Header for MLCEmbeddingDescriptor */
// The class instance for the [CEmbeddingDescriptor] class.
var (
	CEmbeddingDescriptorClass     _CEmbeddingDescriptorClass
	CEmbeddingDescriptorClassOnce sync.Once
)

func getCEmbeddingDescriptorClass() _CEmbeddingDescriptorClass {
	CEmbeddingDescriptorClassOnce.Do(func() {
		CEmbeddingDescriptorClass = _CEmbeddingDescriptorClass{objc.GetClass("MLCEmbeddingDescriptor")}
	})
	return CEmbeddingDescriptorClass
}

type _CEmbeddingDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CEmbeddingDescriptor */
// An interface definition for the [CEmbeddingDescriptor] class.
type ICEmbeddingDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CEmbeddingDescriptor */
	// properties:
	EmbeddingCount() objc.IObject /* cross-framework: NSNumber */
	EmbeddingDimension() objc.IObject /* cross-framework: NSNumber */
	MaximumNorm() objc.IObject /* cross-framework: NSNumber */
	PaddingIndex() objc.IObject /* cross-framework: NSNumber */
	PNorm() objc.IObject /* cross-framework: NSNumber */
	ScalesGradientByFrequency() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CEmbeddingDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CEmbeddingDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CEmbeddingDescriptorClass) Alloc() CEmbeddingDescriptor {
	rv := objc.Send[CEmbeddingDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CEmbeddingDescriptorClass) New() CEmbeddingDescriptor {
	rv := objc.Send[CEmbeddingDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CEmbeddingDescriptor) Init() CEmbeddingDescriptor {
	rv := objc.Send[CEmbeddingDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CEmbeddingDescriptor) Autorelease() CEmbeddingDescriptor {
	rv := objc.Send[CEmbeddingDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCEmbeddingDescriptor creates a new CEmbeddingDescriptor instance.
func NewCEmbeddingDescriptor() CEmbeddingDescriptor {
	return getCEmbeddingDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CEmbeddingDescriptor */
// A configuration object you use to create an embedding layer.


// A configuration object you use to create an embedding layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingDescriptor
type CEmbeddingDescriptor struct {
	objectivec.Object
}

// CEmbeddingDescriptorFrom constructs a [CEmbeddingDescriptor] from an unsafe.Pointer.
//
// A configuration object you use to create an embedding layer.
func CEmbeddingDescriptorFrom(ptr unsafe.Pointer) CEmbeddingDescriptor {
	return CEmbeddingDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CEmbeddingDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CEmbeddingDescriptor */

// Creates an embedding descriptor with the size of the dictionary and dimension of embedding vectors you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingDescriptor/descriptorWithEmbeddingCount:embeddingDimension:
func (cc _CEmbeddingDescriptorClass) DescriptorWithEmbeddingCountEmbeddingDimension(embeddingCount objc.IObject /* cross-framework: NSNumber */, embeddingDimension objc.IObject /* cross-framework: NSNumber */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithEmbeddingCount:embeddingDimension:"), embeddingCount, embeddingDimension)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithEmbeddingCountEmbeddingDimension) */


// Creates a new embedding descriptor with the size and dimension of embedding vectors, padding index, and norm and scaling options that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingDescriptor/descriptorWithEmbeddingCount:embeddingDimension:paddingIndex:maximumNorm:pNorm:scalesGradientByFrequency:
func (cc _CEmbeddingDescriptorClass) DescriptorWithEmbeddingCountEmbeddingDimensionPaddingIndexMaximumNormPNormScalesGradientByFrequency(embeddingCount objc.IObject /* cross-framework: NSNumber */, embeddingDimension objc.IObject /* cross-framework: NSNumber */, paddingIndex objc.IObject /* cross-framework: NSNumber */, maximumNorm objc.IObject /* cross-framework: NSNumber */, pNorm objc.IObject /* cross-framework: NSNumber */, scalesGradientByFrequency bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithEmbeddingCount:embeddingDimension:paddingIndex:maximumNorm:pNorm:scalesGradientByFrequency:"), embeddingCount, embeddingDimension, paddingIndex, maximumNorm, pNorm, scalesGradientByFrequency)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithEmbeddingCountEmbeddingDimensionPaddingIndexMaximumNormPNormScalesGradientByFrequency) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CEmbeddingDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CEmbeddingDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CEmbeddingDescriptor */

// The size of the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingDescriptor/embeddingCount-5vs5t
func (c_ CEmbeddingDescriptor) EmbeddingCount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("embeddingCount"))
	return rv
}/* debug [instance_properties/getter]: embeddingCount */


// The dimension of embedding vectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingDescriptor/embeddingDimension-3u8w7
func (c_ CEmbeddingDescriptor) EmbeddingDimension() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("embeddingDimension"))
	return rv
}/* debug [instance_properties/getter]: embeddingDimension */


// A float value that, if set, causes the layer to renormalize the selected embedding vectors to have an Lp norm less than this value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingDescriptor/maximumNorm-4mrjj
func (c_ CEmbeddingDescriptor) MaximumNorm() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("maximumNorm"))
	return rv
}/* debug [instance_properties/getter]: maximumNorm */


// An unsigned integer value that, if set, causes the layer to initialize the embedding vector at that index to zero.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingDescriptor/paddingIndex-50o5o
func (c_ CEmbeddingDescriptor) PaddingIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("paddingIndex"))
	return rv
}/* debug [instance_properties/getter]: paddingIndex */


// The p of the Lp norm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingDescriptor/pNorm-8mto8
func (c_ CEmbeddingDescriptor) PNorm() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pNorm"))
	return rv
}/* debug [instance_properties/getter]: pNorm */


// A Boolean that indicates whether the layer scales gradients by the inverse of the frequency of words in batch before the weight update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingDescriptor/scalesGradientByFrequency
func (c_ CEmbeddingDescriptor) ScalesGradientByFrequency() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("scalesGradientByFrequency"))
	return rv
}/* debug [instance_properties/getter]: scalesGradientByFrequency */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCEmbeddingDescriptor */



