// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CompileOptions] class.
var (
	CompileOptionsClass     _CompileOptionsClass
	CompileOptionsClassOnce sync.Once
)

func getCompileOptionsClass() _CompileOptionsClass {
	CompileOptionsClassOnce.Do(func() {
		CompileOptionsClass = _CompileOptionsClass{objc.GetClass("MTLCompileOptions")}
	})
	return CompileOptionsClass
}

type _CompileOptionsClass struct {
	class objc.Class
}

// An interface definition for the [CompileOptions] class.
type ICompileOptions interface {
	objectivec.IObject
}

// Compilation settings for a Metal shader library.
//
// You can configure the Metal compiler’s options by setting any or all of an instance’s properties, including the following: Target previous OS releases by assigning the property to an case. Set preprocessor macros for the Metal compiler by assigning a dictionary to the property. Choose what the Metal compiler’s optimizer prioritizes by setting the property to an case. Allow the compiler to optimize for floating-point arithmetic that may violate the IEEE 754 standard by setting to . You can compile a library with your compile options instance by calling an instance’s or method.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions
type CompileOptions struct {
	objectivec.Object
}

// CompileOptionsFrom constructs a [CompileOptions] from an unsafe.Pointer.
//
// Compilation settings for a Metal shader library.
func CompileOptionsFrom(ptr unsafe.Pointer) CompileOptions {
	return CompileOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CompileOptionsClass) Alloc() CompileOptions {
	rv := objc.Send[CompileOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CompileOptionsClass) New() CompileOptions {
	rv := objc.Send[CompileOptions](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CompileOptions) Init() CompileOptions {
	rv := objc.Send[CompileOptions](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CompileOptions) Autorelease() CompileOptions {
	rv := objc.Send[CompileOptions](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCompileOptions creates a new CompileOptions instance.
func NewCompileOptions() CompileOptions {
	return getCompileOptionsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/allowReferencingUndefinedSymbols
func (c_ CompileOptions) AllowReferencingUndefinedSymbols() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowReferencingUndefinedSymbols"))
	return rv
}


// SetAllowReferencingUndefinedSymbols sets the value of the allowReferencingUndefinedSymbols property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/allowReferencingUndefinedSymbols
func (c_ CompileOptions) SetAllowReferencingUndefinedSymbols(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowReferencingUndefinedSymbols:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/compileSymbolVisibility
func (c_ CompileOptions) CompileSymbolVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("compileSymbolVisibility"))
	return rv
}


// SetCompileSymbolVisibility sets the value of the compileSymbolVisibility property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/compileSymbolVisibility
func (c_ CompileOptions) SetCompileSymbolVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompileSymbolVisibility:"), value)
}

// A Boolean value that enables shader logging.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/enableLogging
func (c_ CompileOptions) EnableLogging() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enableLogging"))
	return rv
}


// SetEnableLogging sets the value of the enableLogging property.
// A Boolean value that enables shader logging.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/enableLogging
func (c_ CompileOptions) SetEnableLogging(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnableLogging:"), value)
}

// A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/fastMathEnabled
func (c_ CompileOptions) FastMathEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("fastMathEnabled"))
	return rv
}


// SetFastMathEnabled sets the value of the fastMathEnabled property.
// A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/fastMathEnabled
func (c_ CompileOptions) SetFastMathEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFastMathEnabled:"), value)
}

// For a dynamic library, the name to use when installing the library.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/installName
func (c_ CompileOptions) InstallName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("installName"))
	return rv
}


// SetInstallName sets the value of the installName property.
// For a dynamic library, the name to use when installing the library.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/installName
func (c_ CompileOptions) SetInstallName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstallName:"), objc.String(value))
}

// The language version for interpreting the library source code.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/languageVersion
func (c_ CompileOptions) LanguageVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("languageVersion"))
	return rv
}


// SetLanguageVersion sets the value of the languageVersion property.
// The language version for interpreting the library source code.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/languageVersion
func (c_ CompileOptions) SetLanguageVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguageVersion:"), value)
}

// An array of dynamic libraries the Metal compiler links against.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/libraries
func (c_ CompileOptions) Libraries() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("libraries"))
	return rv
}


// SetLibraries sets the value of the libraries property.
// An array of dynamic libraries the Metal compiler links against.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/libraries
func (c_ CompileOptions) SetLibraries(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setLibraries:"), nsArray)
}

// The kind of library to create.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/libraryType
func (c_ CompileOptions) LibraryType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("libraryType"))
	return rv
}


// SetLibraryType sets the value of the libraryType property.
// The kind of library to create.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/libraryType
func (c_ CompileOptions) SetLibraryType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLibraryType:"), value)
}

// The FP32 math functions Metal uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/mathFloatingPointFunctions
func (c_ CompileOptions) MathFloatingPointFunctions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("mathFloatingPointFunctions"))
	return rv
}


// SetMathFloatingPointFunctions sets the value of the mathFloatingPointFunctions property.
// The FP32 math functions Metal uses.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/mathFloatingPointFunctions
func (c_ CompileOptions) SetMathFloatingPointFunctions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMathFloatingPointFunctions:"), value)
}

// An indication of whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/mathMode
func (c_ CompileOptions) MathMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("mathMode"))
	return rv
}


// SetMathMode sets the value of the mathMode property.
// An indication of whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/mathMode
func (c_ CompileOptions) SetMathMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMathMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/maxTotalThreadsPerThreadgroup
func (c_ CompileOptions) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}


// SetMaxTotalThreadsPerThreadgroup sets the value of the maxTotalThreadsPerThreadgroup property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/maxTotalThreadsPerThreadgroup
func (c_ CompileOptions) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}

// An option that tells the compiler what to prioritize when it compiles Metal shader code.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/optimizationLevel
func (c_ CompileOptions) OptimizationLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("optimizationLevel"))
	return rv
}


// SetOptimizationLevel sets the value of the optimizationLevel property.
// An option that tells the compiler what to prioritize when it compiles Metal shader code.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/optimizationLevel
func (c_ CompileOptions) SetOptimizationLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOptimizationLevel:"), value)
}

// A list of preprocessor macros to apply when compiling the library source.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/preprocessorMacros
func (c_ CompileOptions) PreprocessorMacros() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("preprocessorMacros"))
	return rv
}


// SetPreprocessorMacros sets the value of the preprocessorMacros property.
// A list of preprocessor macros to apply when compiling the library source.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/preprocessorMacros
func (c_ CompileOptions) SetPreprocessorMacros(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreprocessorMacros:"), value)
}

// A Boolean value that indicates whether the compiler compiles vertex shaders conservatively to generate consistent position calculations.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/preserveInvariance
func (c_ CompileOptions) PreserveInvariance() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preserveInvariance"))
	return rv
}


// SetPreserveInvariance sets the value of the preserveInvariance property.
// A Boolean value that indicates whether the compiler compiles vertex shaders conservatively to generate consistent position calculations.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/preserveInvariance
func (c_ CompileOptions) SetPreserveInvariance(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreserveInvariance:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/requiredThreadsPerThreadgroup
func (c_ CompileOptions) RequiredThreadsPerThreadgroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}


// SetRequiredThreadsPerThreadgroup sets the value of the requiredThreadsPerThreadgroup property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/requiredThreadsPerThreadgroup
func (c_ CompileOptions) SetRequiredThreadsPerThreadgroup(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}



