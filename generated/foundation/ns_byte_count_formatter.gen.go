// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ByteCountFormatter] class.
var (
	ByteCountFormatterClass     _ByteCountFormatterClass
	ByteCountFormatterClassOnce sync.Once
)

func getByteCountFormatterClass() _ByteCountFormatterClass {
	ByteCountFormatterClassOnce.Do(func() {
		ByteCountFormatterClass = _ByteCountFormatterClass{objc.GetClass("NSByteCountFormatter")}
	})
	return ByteCountFormatterClass
}

type _ByteCountFormatterClass struct {
	class objc.Class
}

// An interface definition for the [ByteCountFormatter] class.
type IByteCountFormatter interface {
	IFormatter
}

// A formatter that converts a byte count value into a localized description that is formatted with the appropriate byte modifier (KB, MB, GB and so on).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter
type ByteCountFormatter struct {
	Formatter
}

// ByteCountFormatterFrom constructs a [ByteCountFormatter] from an unsafe.Pointer.
//
// A formatter that converts a byte count value into a localized description that is formatted with the appropriate byte modifier (KB, MB, GB and so on).
func ByteCountFormatterFrom(ptr unsafe.Pointer) ByteCountFormatter {
	return ByteCountFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _ByteCountFormatterClass) Alloc() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _ByteCountFormatterClass) New() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ ByteCountFormatter) Init() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ ByteCountFormatter) Autorelease() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewByteCountFormatter creates a new ByteCountFormatter instance.
func NewByteCountFormatter() ByteCountFormatter {
	return getByteCountFormatterClass().New()
}


// Specify the units that can be used in the output.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/allowedunits
func (b_ ByteCountFormatter) AllowedUnits() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("allowedUnits"))
	return rv
}


// SetAllowedUnits sets the value of the allowedUnits property.
// Specify the units that can be used in the output.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/allowedunits
func (b_ ByteCountFormatter) SetAllowedUnits(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowedUnits:"), value)
}

// Determines whether to allow more natural display of some values.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/allowsnonnumericformatting
func (b_ ByteCountFormatter) AllowsNonnumericFormatting() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsNonnumericFormatting"))
	return rv
}


// SetAllowsNonnumericFormatting sets the value of the allowsNonnumericFormatting property.
// Determines whether to allow more natural display of some values.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/allowsnonnumericformatting
func (b_ ByteCountFormatter) SetAllowsNonnumericFormatting(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsNonnumericFormatting:"), value)
}

// Specify the number of bytes to be used for kilobytes.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/countstyle-swift.property
func (b_ ByteCountFormatter) CountStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("countStyle"))
	return rv
}


// SetCountStyle sets the value of the countStyle property.
// Specify the number of bytes to be used for kilobytes.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/countstyle-swift.property
func (b_ ByteCountFormatter) SetCountStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCountStyle:"), value)
}

// Specify the formatting context for the formatted string.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/formattingcontext
func (b_ ByteCountFormatter) FormattingContext() int {
	rv := objc.Send[int](b_.ID, objc.Sel("formattingContext"))
	return rv
}


// SetFormattingContext sets the value of the formattingContext property.
// Specify the formatting context for the formatted string.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/formattingcontext
func (b_ ByteCountFormatter) SetFormattingContext(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFormattingContext:"), value)
}

// Determines whether to include the number of bytes after the formatted string.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/includesactualbytecount
func (b_ ByteCountFormatter) IncludesActualByteCount() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("includesActualByteCount"))
	return rv
}


// SetIncludesActualByteCount sets the value of the includesActualByteCount property.
// Determines whether to include the number of bytes after the formatted string.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/includesactualbytecount
func (b_ ByteCountFormatter) SetIncludesActualByteCount(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIncludesActualByteCount:"), value)
}

// Determines whether to include the count in the resulting formatted string.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/includescount
func (b_ ByteCountFormatter) IncludesCount() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("includesCount"))
	return rv
}


// SetIncludesCount sets the value of the includesCount property.
// Determines whether to include the count in the resulting formatted string.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/includescount
func (b_ ByteCountFormatter) SetIncludesCount(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIncludesCount:"), value)
}

// Determines whether to include the units in the resulting formatted string.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/includesunit
func (b_ ByteCountFormatter) IncludesUnit() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("includesUnit"))
	return rv
}


// SetIncludesUnit sets the value of the includesUnit property.
// Determines whether to include the units in the resulting formatted string.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/includesunit
func (b_ ByteCountFormatter) SetIncludesUnit(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIncludesUnit:"), value)
}

// Determines the display style of the size representation.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/isadaptive
func (b_ ByteCountFormatter) IsAdaptive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isAdaptive"))
	return rv
}


// SetIsAdaptive sets the value of the isAdaptive property.
// Determines the display style of the size representation.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/isadaptive
func (b_ ByteCountFormatter) SetIsAdaptive(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsAdaptive:"), value)
}

// Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/zeropadsfractiondigits
func (b_ ByteCountFormatter) ZeroPadsFractionDigits() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("zeroPadsFractionDigits"))
	return rv
}


// SetZeroPadsFractionDigits sets the value of the zeroPadsFractionDigits property.
// Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/zeropadsfractiondigits
func (b_ ByteCountFormatter) SetZeroPadsFractionDigits(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setZeroPadsFractionDigits:"), value)
}



