// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DDMatchLink */


/* debug [class_header]: Header for DDMatchLink */
// The class instance for the [DDMatchLink] class.
var (
	DDMatchLinkClass     _DDMatchLinkClass
	DDMatchLinkClassOnce sync.Once
)

func getDDMatchLinkClass() _DDMatchLinkClass {
	DDMatchLinkClassOnce.Do(func() {
		DDMatchLinkClass = _DDMatchLinkClass{objc.GetClass("DDMatchLink")}
	})
	return DDMatchLinkClass
}

type _DDMatchLinkClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DDMatchLink */
// An interface definition for the [DDMatchLink] class.
type IDDMatchLink interface {
	IDDMatch
	
/* debug [class_interface_properties]: Properties for DDMatchLink */
	// properties:
	URL() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DDMatchLink */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DDMatchLink */
// Alloc allocates a new instance without initialization.
func (dc _DDMatchLinkClass) Alloc() DDMatchLink {
	rv := objc.Send[DDMatchLink](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DDMatchLinkClass) New() DDMatchLink {
	rv := objc.Send[DDMatchLink](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchLink) Init() DDMatchLink {
	rv := objc.Send[DDMatchLink](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchLink) Autorelease() DDMatchLink {
	rv := objc.Send[DDMatchLink](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchLink creates a new DDMatchLink instance.
func NewDDMatchLink() DDMatchLink {
	return getDDMatchLinkClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DDMatchLink */
// An object that contains a web link that the data detection system matches.
//
// The DataDetection framework returns a link match in a object, which contains a .


// An object that contains a web link that the data detection system matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchLink
type DDMatchLink struct {
	DDMatch
}

// DDMatchLinkFrom constructs a [DDMatchLink] from an unsafe.Pointer.
//
// An object that contains a web link that the data detection system matches.
func DDMatchLinkFrom(ptr unsafe.Pointer) DDMatchLink {
	return DDMatchLink{
		DDMatch: DDMatchFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DDMatchLink *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DDMatchLink */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DDMatchLink */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DDMatchLink */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DDMatchLink */

// An address for a web resource, such as a webpage or image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchLink/url
func (d_ DDMatchLink) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DDMatchLink */



