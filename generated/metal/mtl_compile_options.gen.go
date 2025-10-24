// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLCompileOptions */


/* debug [class_header]: Header for MTLCompileOptions */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CompileOptions */
// An interface definition for the [CompileOptions] class.
type ICompileOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CompileOptions */
	// properties:
	AllowReferencingUndefinedSymbols() bool
	SetAllowReferencingUndefinedSymbols(value bool)
	CompileSymbolVisibility() CompileSymbolVisibility
	SetCompileSymbolVisibility(value CompileSymbolVisibility)
	EnableLogging() bool
	SetEnableLogging(value bool)
	FastMathEnabled() bool
	SetFastMathEnabled(value bool)
	InstallName() objc.IObject /* cross-framework: NSString */
	SetInstallName(value objc.IObject /* cross-framework: NSString */)
	LanguageVersion() LanguageVersion
	SetLanguageVersion(value LanguageVersion)
	Libraries() []objc.ID
	SetLibraries(value []objc.ID)
	LibraryType() LibraryType
	SetLibraryType(value LibraryType)
	MathFloatingPointFunctions() MathFloatingPointFunctions
	SetMathFloatingPointFunctions(value MathFloatingPointFunctions)
	MathMode() MathMode
	SetMathMode(value MathMode)
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	OptimizationLevel() LibraryOptimizationLevel
	SetOptimizationLevel(value LibraryOptimizationLevel)
	PreprocessorMacros() foundation.IDictionary
	SetPreprocessorMacros(value foundation.IDictionary)
	PreserveInvariance() bool
	SetPreserveInvariance(value bool)
	RequiredThreadsPerThreadgroup() objc.IObject /* cross-framework: MTLSize */
	SetRequiredThreadsPerThreadgroup(value objc.IObject /* cross-framework: MTLSize */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CompileOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CompileOptions */
// Alloc allocates a new instance without initialization.
func (cc _CompileOptionsClass) Alloc() CompileOptions {
	rv := objc.Send[CompileOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CompileOptions */
// Compilation settings for a Metal shader library.
//
// You can configure the Metal compiler’s options by setting any or all of an instance’s properties, including the following: Target previous OS releases by assigning the property to an case. Set preprocessor macros for the Metal compiler by assigning a dictionary to the property. Choose what the Metal compiler’s optimizer prioritizes by setting the property to an case. Allow the compiler to optimize for floating-point arithmetic that may violate the IEEE 754 standard by setting to . You can compile a library with your compile options instance by calling an instance’s or method.


// Compilation settings for a Metal shader library.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CompileOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CompileOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CompileOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CompileOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CompileOptions */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/allowReferencingUndefinedSymbols
func (c_ CompileOptions) AllowReferencingUndefinedSymbols() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowReferencingUndefinedSymbols"))
	return rv
}/* debug [instance_properties/getter]: allowReferencingUndefinedSymbols */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/allowReferencingUndefinedSymbols
func (c_ CompileOptions) SetAllowReferencingUndefinedSymbols(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowReferencingUndefinedSymbols:"), value)
}/* debug [instance_properties/setter]: allowReferencingUndefinedSymbols */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/compileSymbolVisibility
func (c_ CompileOptions) CompileSymbolVisibility() CompileSymbolVisibility {
	rv := objc.Send[CompileSymbolVisibility](c_.ID, objc.Sel("compileSymbolVisibility"))
	return rv
}/* debug [instance_properties/getter]: compileSymbolVisibility */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/compileSymbolVisibility
func (c_ CompileOptions) SetCompileSymbolVisibility(value CompileSymbolVisibility) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompileSymbolVisibility:"), value)
}/* debug [instance_properties/setter]: compileSymbolVisibility */


// A Boolean value that enables shader logging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/enableLogging
func (c_ CompileOptions) EnableLogging() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enableLogging"))
	return rv
}/* debug [instance_properties/getter]: enableLogging */


// A Boolean value that enables shader logging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/enableLogging
func (c_ CompileOptions) SetEnableLogging(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnableLogging:"), value)
}/* debug [instance_properties/setter]: enableLogging */


// A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/fastMathEnabled
func (c_ CompileOptions) FastMathEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("fastMathEnabled"))
	return rv
}/* debug [instance_properties/getter]: fastMathEnabled */


// A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/fastMathEnabled
func (c_ CompileOptions) SetFastMathEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFastMathEnabled:"), value)
}/* debug [instance_properties/setter]: fastMathEnabled */


// For a dynamic library, the name to use when installing the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/installName
func (c_ CompileOptions) InstallName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("installName"))
	return rv
}/* debug [instance_properties/getter]: installName */


// For a dynamic library, the name to use when installing the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/installName
func (c_ CompileOptions) SetInstallName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstallName:"), value)
}/* debug [instance_properties/setter]: installName */


// The language version for interpreting the library source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/languageVersion
func (c_ CompileOptions) LanguageVersion() LanguageVersion {
	rv := objc.Send[LanguageVersion](c_.ID, objc.Sel("languageVersion"))
	return rv
}/* debug [instance_properties/getter]: languageVersion */


// The language version for interpreting the library source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/languageVersion
func (c_ CompileOptions) SetLanguageVersion(value LanguageVersion) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguageVersion:"), value)
}/* debug [instance_properties/setter]: languageVersion */


// An array of dynamic libraries the Metal compiler links against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/libraries
func (c_ CompileOptions) Libraries() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("libraries"))
	return rv
}/* debug [instance_properties/getter]: libraries */


// An array of dynamic libraries the Metal compiler links against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/libraries
func (c_ CompileOptions) SetLibraries(value []objc.ID) {
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
}/* debug [instance_properties/setter]: libraries */


// The kind of library to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/libraryType
func (c_ CompileOptions) LibraryType() LibraryType {
	rv := objc.Send[LibraryType](c_.ID, objc.Sel("libraryType"))
	return rv
}/* debug [instance_properties/getter]: libraryType */


// The kind of library to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/libraryType
func (c_ CompileOptions) SetLibraryType(value LibraryType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLibraryType:"), value)
}/* debug [instance_properties/setter]: libraryType */


// The FP32 math functions Metal uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/mathFloatingPointFunctions
func (c_ CompileOptions) MathFloatingPointFunctions() MathFloatingPointFunctions {
	rv := objc.Send[MathFloatingPointFunctions](c_.ID, objc.Sel("mathFloatingPointFunctions"))
	return rv
}/* debug [instance_properties/getter]: mathFloatingPointFunctions */


// The FP32 math functions Metal uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/mathFloatingPointFunctions
func (c_ CompileOptions) SetMathFloatingPointFunctions(value MathFloatingPointFunctions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMathFloatingPointFunctions:"), value)
}/* debug [instance_properties/setter]: mathFloatingPointFunctions */


// An indication of whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/mathMode
func (c_ CompileOptions) MathMode() MathMode {
	rv := objc.Send[MathMode](c_.ID, objc.Sel("mathMode"))
	return rv
}/* debug [instance_properties/getter]: mathMode */


// An indication of whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/mathMode
func (c_ CompileOptions) SetMathMode(value MathMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMathMode:"), value)
}/* debug [instance_properties/setter]: mathMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/maxTotalThreadsPerThreadgroup
func (c_ CompileOptions) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: maxTotalThreadsPerThreadgroup */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/maxTotalThreadsPerThreadgroup
func (c_ CompileOptions) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}/* debug [instance_properties/setter]: maxTotalThreadsPerThreadgroup */


// An option that tells the compiler what to prioritize when it compiles Metal shader code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/optimizationLevel
func (c_ CompileOptions) OptimizationLevel() LibraryOptimizationLevel {
	rv := objc.Send[LibraryOptimizationLevel](c_.ID, objc.Sel("optimizationLevel"))
	return rv
}/* debug [instance_properties/getter]: optimizationLevel */


// An option that tells the compiler what to prioritize when it compiles Metal shader code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/optimizationLevel
func (c_ CompileOptions) SetOptimizationLevel(value LibraryOptimizationLevel) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOptimizationLevel:"), value)
}/* debug [instance_properties/setter]: optimizationLevel */


// A list of preprocessor macros to apply when compiling the library source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/preprocessorMacros
func (c_ CompileOptions) PreprocessorMacros() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("preprocessorMacros"))
	return rv
}/* debug [instance_properties/getter]: preprocessorMacros */


// A list of preprocessor macros to apply when compiling the library source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/preprocessorMacros
func (c_ CompileOptions) SetPreprocessorMacros(value foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreprocessorMacros:"), value)
}/* debug [instance_properties/setter]: preprocessorMacros */


// A Boolean value that indicates whether the compiler compiles vertex shaders conservatively to generate consistent position calculations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/preserveInvariance
func (c_ CompileOptions) PreserveInvariance() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preserveInvariance"))
	return rv
}/* debug [instance_properties/getter]: preserveInvariance */


// A Boolean value that indicates whether the compiler compiles vertex shaders conservatively to generate consistent position calculations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/preserveInvariance
func (c_ CompileOptions) SetPreserveInvariance(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreserveInvariance:"), value)
}/* debug [instance_properties/setter]: preserveInvariance */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/requiredThreadsPerThreadgroup
func (c_ CompileOptions) RequiredThreadsPerThreadgroup() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: requiredThreadsPerThreadgroup */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileOptions/requiredThreadsPerThreadgroup
func (c_ CompileOptions) SetRequiredThreadsPerThreadgroup(value objc.IObject /* cross-framework: MTLSize */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}/* debug [instance_properties/setter]: requiredThreadsPerThreadgroup */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLCompileOptions */



