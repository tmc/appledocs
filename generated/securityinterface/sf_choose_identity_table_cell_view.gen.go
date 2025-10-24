// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class SFChooseIdentityTableCellView */

/* debug [class_header]: Header for SFChooseIdentityTableCellView */
// The class instance for the [SFChooseIdentityTableCellView] class.
var (
	SFChooseIdentityTableCellViewClass     _SFChooseIdentityTableCellViewClass
	SFChooseIdentityTableCellViewClassOnce sync.Once
)

func getSFChooseIdentityTableCellViewClass() _SFChooseIdentityTableCellViewClass {
	SFChooseIdentityTableCellViewClassOnce.Do(func() {
		SFChooseIdentityTableCellViewClass = _SFChooseIdentityTableCellViewClass{objc.GetClass("SFChooseIdentityTableCellView")}
	})
	return SFChooseIdentityTableCellViewClass
}

type _SFChooseIdentityTableCellViewClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFChooseIdentityTableCellView */
// An interface definition for the [SFChooseIdentityTableCellView] class.
type ISFChooseIdentityTableCellView interface {
	appkit.ITableCellView

	/* debug [class_interface_properties]: Properties for SFChooseIdentityTableCellView */
	// properties:
	IssuerTextField() appkit.TextField
	SetIssuerTextField(value appkit.TextField)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFChooseIdentityTableCellView */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFChooseIdentityTableCellView */
// Alloc allocates a new instance without initialization.
func (sc _SFChooseIdentityTableCellViewClass) Alloc() SFChooseIdentityTableCellView {
	rv := objc.Send[SFChooseIdentityTableCellView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFChooseIdentityTableCellViewClass) New() SFChooseIdentityTableCellView {
	rv := objc.Send[SFChooseIdentityTableCellView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFChooseIdentityTableCellView) Init() SFChooseIdentityTableCellView {
	rv := objc.Send[SFChooseIdentityTableCellView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFChooseIdentityTableCellView) Autorelease() SFChooseIdentityTableCellView {
	rv := objc.Send[SFChooseIdentityTableCellView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFChooseIdentityTableCellView creates a new SFChooseIdentityTableCellView instance.
func NewSFChooseIdentityTableCellView() SFChooseIdentityTableCellView {
	return getSFChooseIdentityTableCellViewClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFChooseIdentityTableCellView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityTableCellView
type SFChooseIdentityTableCellView struct {
	appkit.TableCellView
}

// SFChooseIdentityTableCellViewFrom constructs a [SFChooseIdentityTableCellView] from an unsafe.Pointer.
func SFChooseIdentityTableCellViewFrom(ptr unsafe.Pointer) SFChooseIdentityTableCellView {
	return SFChooseIdentityTableCellView{
		TableCellView: appkit.TableCellViewFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFChooseIdentityTableCellView */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFChooseIdentityTableCellView */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFChooseIdentityTableCellView */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFChooseIdentityTableCellView */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFChooseIdentityTableCellView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityTableCellView/issuerTextField-swift.property
func (s_ SFChooseIdentityTableCellView) IssuerTextField() appkit.TextField {
	rv := objc.Send[appkit.TextField](s_.ID, objc.Sel("issuerTextField"))
	return rv
} /* debug [instance_properties/getter]: issuerTextField */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityTableCellView/issuerTextField-swift.property
func (s_ SFChooseIdentityTableCellView) SetIssuerTextField(value appkit.TextField) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIssuerTextField:"), value)
} /* debug [instance_properties/setter]: issuerTextField */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFChooseIdentityTableCellView */
