// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTLTextureReferenceType */


/* debug [class_header]: Header for MTLTextureReferenceType */
// The class instance for the [TextureReferenceType] class.
var (
	TextureReferenceTypeClass     _TextureReferenceTypeClass
	TextureReferenceTypeClassOnce sync.Once
)

func getTextureReferenceTypeClass() _TextureReferenceTypeClass {
	TextureReferenceTypeClassOnce.Do(func() {
		TextureReferenceTypeClass = _TextureReferenceTypeClass{objc.GetClass("MTLTextureReferenceType")}
	})
	return TextureReferenceTypeClass
}

type _TextureReferenceTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextureReferenceType */
// An interface definition for the [TextureReferenceType] class.
type ITextureReferenceType interface {
	IType
	
/* debug [class_interface_properties]: Properties for TextureReferenceType */
	// properties:
	Access() BindingAccess
	IsDepthTexture() bool
	TextureDataType() DataType
	TextureType() TextureType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextureReferenceType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextureReferenceType */
// Alloc allocates a new instance without initialization.
func (tc _TextureReferenceTypeClass) Alloc() TextureReferenceType {
	rv := objc.Send[TextureReferenceType](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextureReferenceTypeClass) New() TextureReferenceType {
	rv := objc.Send[TextureReferenceType](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextureReferenceType) Init() TextureReferenceType {
	rv := objc.Send[TextureReferenceType](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextureReferenceType) Autorelease() TextureReferenceType {
	rv := objc.Send[TextureReferenceType](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextureReferenceType creates a new TextureReferenceType instance.
func NewTextureReferenceType() TextureReferenceType {
	return getTextureReferenceTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextureReferenceType */
// A description of a texture.


// A description of a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType
type TextureReferenceType struct {
	Type
}

// TextureReferenceTypeFrom constructs a [TextureReferenceType] from an unsafe.Pointer.
//
// A description of a texture.
func TextureReferenceTypeFrom(ptr unsafe.Pointer) TextureReferenceType {
	return TextureReferenceType{
		Type: TypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextureReferenceType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextureReferenceType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextureReferenceType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextureReferenceType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextureReferenceType */

// The texture’s read/write access to the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType/access
func (t_ TextureReferenceType) Access() BindingAccess {
	rv := objc.Send[BindingAccess](t_.ID, objc.Sel("access"))
	return rv
}/* debug [instance_properties/getter]: access */


// A Boolean value that indicates whether the texture is a depth texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType/isDepthTexture
func (t_ TextureReferenceType) IsDepthTexture() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isDepthTexture"))
	return rv
}/* debug [instance_properties/getter]: isDepthTexture */


// The data type of the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType/textureDataType
func (t_ TextureReferenceType) TextureDataType() DataType {
	rv := objc.Send[DataType](t_.ID, objc.Sel("textureDataType"))
	return rv
}/* debug [instance_properties/getter]: textureDataType */


// The texture type of the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType/textureType
func (t_ TextureReferenceType) TextureType() TextureType {
	rv := objc.Send[TextureType](t_.ID, objc.Sel("textureType"))
	return rv
}/* debug [instance_properties/getter]: textureType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLTextureReferenceType */



