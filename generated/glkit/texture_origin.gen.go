// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class textureOrigin */


/* debug [class_header]: Header for textureOrigin */
// The class instance for the [textureOrigin] class.
var (
	TextureOriginClass     _textureOriginClass
	TextureOriginClassOnce sync.Once
)

func gettextureOriginClass() _textureOriginClass {
	TextureOriginClassOnce.Do(func() {
		TextureOriginClass = _textureOriginClass{objc.GetClass("textureOrigin")}
	})
	return TextureOriginClass
}

type _textureOriginClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for textureOrigin */
// An interface definition for the [textureOrigin] class.
type ItextureOrigin interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for textureOrigin */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for textureOrigin */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for textureOrigin */
// Alloc allocates a new instance without initialization.
func (tc _textureOriginClass) Alloc() textureOrigin {
	rv := objc.Send[textureOrigin](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _textureOriginClass) New() textureOrigin {
	rv := objc.Send[textureOrigin](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ textureOrigin) Init() textureOrigin {
	rv := objc.Send[textureOrigin](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ textureOrigin) Autorelease() textureOrigin {
	rv := objc.Send[textureOrigin](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtextureOrigin creates a new textureOrigin instance.
func NewtextureOrigin() textureOrigin {
	return gettextureOriginClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for textureOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/textureOrigin-c.ivar
type textureOrigin struct {
	objectivec.Object
}

// textureOriginFrom constructs a [textureOrigin] from an unsafe.Pointer.
func textureOriginFrom(ptr unsafe.Pointer) textureOrigin {
	return textureOrigin{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for textureOrigin *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for textureOrigin */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for textureOrigin */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for textureOrigin */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for textureOrigin */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class textureOrigin */



