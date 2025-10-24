// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMProcessingInstruction */


/* debug [class_header]: Header for DOMProcessingInstruction */
// The class instance for the [DOMProcessingInstruction] class.
var (
	DOMProcessingInstructionClass     _DOMProcessingInstructionClass
	DOMProcessingInstructionClassOnce sync.Once
)

func getDOMProcessingInstructionClass() _DOMProcessingInstructionClass {
	DOMProcessingInstructionClassOnce.Do(func() {
		DOMProcessingInstructionClass = _DOMProcessingInstructionClass{objc.GetClass("DOMProcessingInstruction")}
	})
	return DOMProcessingInstructionClass
}

type _DOMProcessingInstructionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMProcessingInstruction */
// An interface definition for the [DOMProcessingInstruction] class.
type IDOMProcessingInstruction interface {
	IDOMCharacterData
	
/* debug [class_interface_properties]: Properties for DOMProcessingInstruction */
	// properties:
	Sheet() IDOMStyleSheet
	Target() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMProcessingInstruction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMProcessingInstruction */
// Alloc allocates a new instance without initialization.
func (dc _DOMProcessingInstructionClass) Alloc() DOMProcessingInstruction {
	rv := objc.Send[DOMProcessingInstruction](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMProcessingInstructionClass) New() DOMProcessingInstruction {
	rv := objc.Send[DOMProcessingInstruction](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMProcessingInstruction) Init() DOMProcessingInstruction {
	rv := objc.Send[DOMProcessingInstruction](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMProcessingInstruction) Autorelease() DOMProcessingInstruction {
	rv := objc.Send[DOMProcessingInstruction](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMProcessingInstruction creates a new DOMProcessingInstruction instance.
func NewDOMProcessingInstruction() DOMProcessingInstruction {
	return getDOMProcessingInstructionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMProcessingInstruction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMProcessingInstruction
type DOMProcessingInstruction struct {
	DOMCharacterData
}

// DOMProcessingInstructionFrom constructs a [DOMProcessingInstruction] from an unsafe.Pointer.
func DOMProcessingInstructionFrom(ptr unsafe.Pointer) DOMProcessingInstruction {
	return DOMProcessingInstruction{
		DOMCharacterData: DOMCharacterDataFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMProcessingInstruction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMProcessingInstruction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMProcessingInstruction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMProcessingInstruction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMProcessingInstruction */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMProcessingInstruction/sheet
func (d_ DOMProcessingInstruction) Sheet() IDOMStyleSheet {
	rv := objc.Send[DOMStyleSheet](d_.ID, objc.Sel("sheet"))
	return rv
}/* debug [instance_properties/getter]: sheet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMProcessingInstruction/target
func (d_ DOMProcessingInstruction) Target() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMProcessingInstruction */



