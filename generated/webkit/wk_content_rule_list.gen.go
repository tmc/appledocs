// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKContentRuleList */

/* debug [class_header]: Header for WKContentRuleList */
// The class instance for the [ContentRuleList] class.
var (
	ContentRuleListClass     _ContentRuleListClass
	ContentRuleListClassOnce sync.Once
)

func getContentRuleListClass() _ContentRuleListClass {
	ContentRuleListClassOnce.Do(func() {
		ContentRuleListClass = _ContentRuleListClass{objc.GetClass("WKContentRuleList")}
	})
	return ContentRuleListClass
}

type _ContentRuleListClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for ContentRuleList */
// An interface definition for the [ContentRuleList] class.
type IContentRuleList interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for ContentRuleList */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for ContentRuleList */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for ContentRuleList */
// Alloc allocates a new instance without initialization.
func (cc _ContentRuleListClass) Alloc() ContentRuleList {
	rv := objc.Send[ContentRuleList](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContentRuleListClass) New() ContentRuleList {
	rv := objc.Send[ContentRuleList](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentRuleList) Init() ContentRuleList {
	rv := objc.Send[ContentRuleList](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentRuleList) Autorelease() ContentRuleList {
	rv := objc.Send[ContentRuleList](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentRuleList creates a new ContentRuleList instance.
func NewContentRuleList() ContentRuleList {
	return getContentRuleListClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for ContentRuleList */
// A compiled list of rules to apply to web content.
//
// A object represents a compiled set of rules for modifying how a webpage loads content. You don’t create a directly. Instead, you specify your rules in JSON format and compile them using the method of . That method compiles your rules into an efficient byte format and returns them in an instance of this class. Content rule lists use the same syntax as content blocker extensions in Safari. For more information on how to specify the JSON for your rule lists, see .

// A compiled list of rules to apply to web content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleList
type ContentRuleList struct {
	objectivec.Object
}

// ContentRuleListFrom constructs a [ContentRuleList] from an unsafe.Pointer.
//
// A compiled list of rules to apply to web content.
func ContentRuleListFrom(ptr unsafe.Pointer) ContentRuleList {
	return ContentRuleList{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for ContentRuleList */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for ContentRuleList */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for ContentRuleList */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for ContentRuleList */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for ContentRuleList */

// The identifier for the rule list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentRuleList/identifier
func (c_ ContentRuleList) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
} /* debug [instance_properties/getter]: identifier */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKContentRuleList */
