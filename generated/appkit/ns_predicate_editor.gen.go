// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSPredicateEditor */


/* debug [class_header]: Header for NSPredicateEditor */
// The class instance for the [PredicateEditor] class.
var (
	PredicateEditorClass     _PredicateEditorClass
	PredicateEditorClassOnce sync.Once
)

func getPredicateEditorClass() _PredicateEditorClass {
	PredicateEditorClassOnce.Do(func() {
		PredicateEditorClass = _PredicateEditorClass{objc.GetClass("NSPredicateEditor")}
	})
	return PredicateEditorClass
}

type _PredicateEditorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PredicateEditor */
// An interface definition for the [PredicateEditor] class.
type IPredicateEditor interface {
	IRuleEditor
	
/* debug [class_interface_properties]: Properties for PredicateEditor */
	// properties:
	RowTemplates() []PredicateEditorRowTemplate
	SetRowTemplates(value []PredicateEditorRowTemplate)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PredicateEditor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PredicateEditor */
// Alloc allocates a new instance without initialization.
func (pc _PredicateEditorClass) Alloc() PredicateEditor {
	rv := objc.Send[PredicateEditor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PredicateEditorClass) New() PredicateEditor {
	rv := objc.Send[PredicateEditor](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PredicateEditor) Init() PredicateEditor {
	rv := objc.Send[PredicateEditor](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PredicateEditor) Autorelease() PredicateEditor {
	rv := objc.Send[PredicateEditor](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPredicateEditor creates a new PredicateEditor instance.
func NewPredicateEditor() PredicateEditor {
	return getPredicateEditorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PredicateEditor */
// A defined set of rules that allows the editing of predicate objects.
//
// provides an property— (inherited from )—that you can get and set directly, and that you can bind using Cocoa bindings (you typically configure a predicate editor in Interface Builder). depends on another class, , that describes the available predicates and how to display them. Unlike , does not depend on its delegate to populate its rows (and ). Instead, its rows are populated from its property (an instance of ). relies on instances , which are responsible for mapping back and forth between the displayed view values and various predicates. exposes one property, , which is an array of objects.


// A defined set of rules that allows the editing of predicate objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditor
type PredicateEditor struct {
	RuleEditor
}

// PredicateEditorFrom constructs a [PredicateEditor] from an unsafe.Pointer.
//
// A defined set of rules that allows the editing of predicate objects.
func PredicateEditorFrom(ptr unsafe.Pointer) PredicateEditor {
	return PredicateEditor{
		RuleEditor: RuleEditorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PredicateEditor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PredicateEditor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PredicateEditor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PredicateEditor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PredicateEditor */

// The row templates for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditor/rowTemplates
func (p_ PredicateEditor) RowTemplates() []PredicateEditorRowTemplate {
	rv := objc.Send[[]PredicateEditorRowTemplate](p_.ID, objc.Sel("rowTemplates"))
	return rv
}/* debug [instance_properties/getter]: rowTemplates */


// The row templates for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditor/rowTemplates
func (p_ PredicateEditor) SetRowTemplates(value []PredicateEditorRowTemplate) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setRowTemplates:"), nsArray)
}/* debug [instance_properties/setter]: rowTemplates */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPredicateEditor */



