// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class templateName */


/* debug [class_header]: Header for templateName */
// The class instance for the [templateName] class.
var (
	TemplateNameClass     _templateNameClass
	TemplateNameClassOnce sync.Once
)

func gettemplateNameClass() _templateNameClass {
	TemplateNameClassOnce.Do(func() {
		TemplateNameClass = _templateNameClass{objc.GetClass("templateName")}
	})
	return TemplateNameClass
}

type _templateNameClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for templateName */
// An interface definition for the [templateName] class.
type ItemplateName interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for templateName */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for templateName */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for templateName */
// Alloc allocates a new instance without initialization.
func (tc _templateNameClass) Alloc() templateName {
	rv := objc.Send[templateName](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _templateNameClass) New() templateName {
	rv := objc.Send[templateName](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ templateName) Init() templateName {
	rv := objc.Send[templateName](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ templateName) Autorelease() templateName {
	rv := objc.Send[templateName](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtemplateName creates a new templateName instance.
func NewtemplateName() templateName {
	return gettemplateNameClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for templateName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/templateName-c.ivar
type templateName struct {
	objectivec.Object
}

// templateNameFrom constructs a [templateName] from an unsafe.Pointer.
func templateNameFrom(ptr unsafe.Pointer) templateName {
	return templateName{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for templateName *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for templateName */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for templateName */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for templateName */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for templateName */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class templateName */



