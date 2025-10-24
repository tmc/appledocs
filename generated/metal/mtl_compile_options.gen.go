// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AllowReferencingUndefinedSymbols() bool
	SetAllowReferencingUndefinedSymbols(value bool)
	CompileSymbolVisibility() CompileSymbolVisibility /* not a class type */
	SetCompileSymbolVisibility(value CompileSymbolVisibility /* not a class type */)
	EnableLogging() bool
	SetEnableLogging(value bool)
	FastMathEnabled() bool
	SetFastMathEnabled(value bool)
	InstallName() objc.IObject /* cross-framework: NSString */
	SetInstallName(value objc.IObject /* cross-framework: NSString */)
	LanguageVersion() LanguageVersion /* not a class type */
	SetLanguageVersion(value LanguageVersion /* not a class type */)
	Libraries() DynamicLibrary /* not a class type */
	SetLibraries(value DynamicLibrary /* not a class type */)
	LibraryType() LibraryType /* not a class type */
	SetLibraryType(value LibraryType /* not a class type */)
	MathFloatingPointFunctions() MathFloatingPointFunctions
	SetMathFloatingPointFunctions(value MathFloatingPointFunctions)
	MathMode() MathMode /* not a class type */
	SetMathMode(value MathMode /* not a class type */)
	MaxTotalThreadsPerThreadgroup() int
	SetMaxTotalThreadsPerThreadgroup(value int)
	OptimizationLevel() LibraryOptimizationLevel /* not a class type */
	SetOptimizationLevel(value LibraryOptimizationLevel /* not a class type */)
	PreprocessorMacros() objc.IObject /* cross-framework: NSObject */
	SetPreprocessorMacros(value objc.IObject /* cross-framework: NSObject */)
	PreserveInvariance() bool
	SetPreserveInvariance(value bool)
	RequiredThreadsPerThreadgroup() objc.IObject /* cross-framework: Size */
	SetRequiredThreadsPerThreadgroup(value objc.IObject /* cross-framework: Size */)
	// methods:
}

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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/allowreferencingundefinedsymbols
func (c_ CompileOptions) AllowReferencingUndefinedSymbols() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowReferencingUndefinedSymbols"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/allowreferencingundefinedsymbols
func (c_ CompileOptions) SetAllowReferencingUndefinedSymbols(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowReferencingUndefinedSymbols:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/compilesymbolvisibility
func (c_ CompileOptions) CompileSymbolVisibility() CompileSymbolVisibility /* not a class type */ {
	rv := objc.Send[CompileSymbolVisibility](c_.ID, objc.Sel("compileSymbolVisibility"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/compilesymbolvisibility
func (c_ CompileOptions) SetCompileSymbolVisibility(value CompileSymbolVisibility /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompileSymbolVisibility:"), value)
}


// A Boolean value that enables shader logging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/enablelogging
func (c_ CompileOptions) EnableLogging() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enableLogging"))
	return rv
}


// A Boolean value that enables shader logging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/enablelogging
func (c_ CompileOptions) SetEnableLogging(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnableLogging:"), value)
}


// A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/fastmathenabled
func (c_ CompileOptions) FastMathEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("fastMathEnabled"))
	return rv
}


// A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/fastmathenabled
func (c_ CompileOptions) SetFastMathEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFastMathEnabled:"), value)
}


// For a dynamic library, the name to use when installing the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/installname
func (c_ CompileOptions) InstallName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("installName"))
	return rv
}


// For a dynamic library, the name to use when installing the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/installname
func (c_ CompileOptions) SetInstallName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstallName:"), value)
}


// The language version for interpreting the library source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/languageversion
func (c_ CompileOptions) LanguageVersion() LanguageVersion /* not a class type */ {
	rv := objc.Send[LanguageVersion](c_.ID, objc.Sel("languageVersion"))
	return rv
}


// The language version for interpreting the library source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/languageversion
func (c_ CompileOptions) SetLanguageVersion(value LanguageVersion /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguageVersion:"), value)
}


// An array of dynamic libraries the Metal compiler links against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/libraries
func (c_ CompileOptions) Libraries() DynamicLibrary /* not a class type */ {
	rv := objc.Send[DynamicLibrary](c_.ID, objc.Sel("libraries"))
	return rv
}


// An array of dynamic libraries the Metal compiler links against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/libraries
func (c_ CompileOptions) SetLibraries(value DynamicLibrary /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLibraries:"), value)
}


// The kind of library to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/librarytype
func (c_ CompileOptions) LibraryType() LibraryType /* not a class type */ {
	rv := objc.Send[LibraryType](c_.ID, objc.Sel("libraryType"))
	return rv
}


// The kind of library to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/librarytype
func (c_ CompileOptions) SetLibraryType(value LibraryType /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLibraryType:"), value)
}


// The FP32 math functions Metal uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/mathfloatingpointfunctions
func (c_ CompileOptions) MathFloatingPointFunctions() MathFloatingPointFunctions {
	rv := objc.Send[MathFloatingPointFunctions](c_.ID, objc.Sel("mathFloatingPointFunctions"))
	return rv
}


// The FP32 math functions Metal uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/mathfloatingpointfunctions
func (c_ CompileOptions) SetMathFloatingPointFunctions(value MathFloatingPointFunctions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMathFloatingPointFunctions:"), value)
}


// An indication of whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/mathmode
func (c_ CompileOptions) MathMode() MathMode /* not a class type */ {
	rv := objc.Send[MathMode](c_.ID, objc.Sel("mathMode"))
	return rv
}


// An indication of whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/mathmode
func (c_ CompileOptions) SetMathMode(value MathMode /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMathMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/maxtotalthreadsperthreadgroup
func (c_ CompileOptions) MaxTotalThreadsPerThreadgroup() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/maxtotalthreadsperthreadgroup
func (c_ CompileOptions) SetMaxTotalThreadsPerThreadgroup(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}


// An option that tells the compiler what to prioritize when it compiles Metal shader code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/optimizationlevel
func (c_ CompileOptions) OptimizationLevel() LibraryOptimizationLevel /* not a class type */ {
	rv := objc.Send[LibraryOptimizationLevel](c_.ID, objc.Sel("optimizationLevel"))
	return rv
}


// An option that tells the compiler what to prioritize when it compiles Metal shader code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/optimizationlevel
func (c_ CompileOptions) SetOptimizationLevel(value LibraryOptimizationLevel /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOptimizationLevel:"), value)
}


// A list of preprocessor macros to apply when compiling the library source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/preprocessormacros
func (c_ CompileOptions) PreprocessorMacros() objc.IObject /* cross-framework: NSObject */ {
	rv := objc.Send[foundation.NSObject](c_.ID, objc.Sel("preprocessorMacros"))
	return rv
}


// A list of preprocessor macros to apply when compiling the library source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/preprocessormacros
func (c_ CompileOptions) SetPreprocessorMacros(value objc.IObject /* cross-framework: NSObject */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreprocessorMacros:"), value)
}


// A Boolean value that indicates whether the compiler compiles vertex shaders conservatively to generate consistent position calculations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/preserveinvariance
func (c_ CompileOptions) PreserveInvariance() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preserveInvariance"))
	return rv
}


// A Boolean value that indicates whether the compiler compiles vertex shaders conservatively to generate consistent position calculations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/preserveinvariance
func (c_ CompileOptions) SetPreserveInvariance(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreserveInvariance:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/requiredthreadsperthreadgroup
func (c_ CompileOptions) RequiredThreadsPerThreadgroup() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](c_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcompileoptions/requiredthreadsperthreadgroup
func (c_ CompileOptions) SetRequiredThreadsPerThreadgroup(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}



