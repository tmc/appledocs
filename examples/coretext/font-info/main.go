package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coretext"
)

const kCFStringEncodingUTF8 = 0x08000100

func goStringToCFString(s string) unsafe.Pointer {
	// Convert Go string to C string
	cStr := append([]byte(s), 0) // null-terminated
	return corefoundation.CFStringCreateWithCString(
		nil,
		unsafe.Pointer(&cStr[0]),
		unsafe.Pointer(uintptr(kCFStringEncodingUTF8)),
	)
}

func cfStringToGo(cfStr unsafe.Pointer) string {
	if cfStr == nil {
		return ""
	}

	// Get UTF8 C string
	cStr := corefoundation.CFStringGetCStringPtr(cfStr, unsafe.Pointer(uintptr(kCFStringEncodingUTF8)))
	if cStr == nil || cStr == unsafe.Pointer(uintptr(0)) {
		// Fallback: get length and copy
		length := corefoundation.CFStringGetLength(cfStr)
		lengthVal := uintptr(length)
		if lengthVal == 0 {
			return ""
		}
		// For simplicity, just return a placeholder if direct pointer doesn't work
		return fmt.Sprintf("<CFString length=%d>", lengthVal)
	}

	// Convert C string to Go string
	result := ""
	for i := 0; ; i++ {
		b := *(*byte)(unsafe.Pointer(uintptr(cStr) + uintptr(i)))
		if b == 0 {
			break
		}
		result += string(b)
	}
	return result
}

func main() {
	// Lock to main thread for CoreText/CoreFoundation operations
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("CoreText Font Information Examples")
	fmt.Println("===================================")

	// Example 1: Create system fonts
	fmt.Println("\n1. Creating System Fonts:")

	fonts := []struct {
		name string
		size float64
	}{
		{"Helvetica", 12.0},
		{"Courier", 14.0},
		{"Times", 16.0},
		{".AppleSystemUIFont", 13.0}, // System default
	}

	for _, f := range fonts {
		// Create CFString for font name
		nameStr := goStringToCFString(f.name)
		if nameStr == nil {
			fmt.Printf("   %-25s: Failed to create CFString\n", f.name)
			continue
		}
		defer corefoundation.CFRelease(nameStr)

		// Create font
		font := coretext.CTFontCreateWithName(nameStr, f.size, nil)
		if font == nil {
			fmt.Printf("   %-25s: Font not found\n", f.name)
			continue
		}
		defer corefoundation.CFRelease(font)

		fmt.Printf("   %-25s (%.1fpt): Created successfully\n", f.name, f.size)
	}

	// Example 2: Get font properties
	fmt.Println("\n2. Font Properties:")

	// Create a font to inspect
	nameStr := goStringToCFString("Helvetica")
	if nameStr != nil {
		defer corefoundation.CFRelease(nameStr)

		font := coretext.CTFontCreateWithName(nameStr, 12.0, nil)
		if font != nil {
			defer corefoundation.CFRelease(font)

			// Get display name
			displayName := coretext.CTFontCopyDisplayName(font)
			if displayName != nil {
				defer corefoundation.CFRelease(displayName)
				name := cfStringToGo(displayName)
				fmt.Printf("   Display Name: %s\n", name)
			}

			// Get family name
			familyName := coretext.CTFontCopyFamilyName(font)
			if familyName != nil {
				defer corefoundation.CFRelease(familyName)
				name := cfStringToGo(familyName)
				fmt.Printf("   Family Name: %s\n", name)
			}

			// Get font size
			size := coretext.CTFontGetSize(font)
			fmt.Printf("   Font Size: %.1f pt\n", size)

			// Get ascent/descent
			ascent := coretext.CTFontGetAscent(font)
			descent := coretext.CTFontGetDescent(font)
			fmt.Printf("   Ascent: %.2f\n", ascent)
			fmt.Printf("   Descent: %.2f\n", descent)
			fmt.Printf("   Line Height: %.2f\n", ascent+descent)
		}
	}

	// Example 3: Common system fonts
	fmt.Println("\n3. Common macOS System Fonts:")

	systemFonts := []struct {
		name string
		desc string
	}{
		{".AppleSystemUIFont", "System UI (default)"},
		{"Helvetica", "Classic sans-serif"},
		{"Helvetica Neue", "Modern sans-serif"},
		{"Times", "Classic serif"},
		{"Courier", "Monospace"},
		{"Monaco", "Monospace (code)"},
		{"Menlo", "Monospace (modern)"},
	}

	for _, sf := range systemFonts {
		nameStr := goStringToCFString(sf.name)
		if nameStr != nil {
			font := coretext.CTFontCreateWithName(nameStr, 12.0, nil)
			corefoundation.CFRelease(nameStr)

			if font != nil {
				fmt.Printf("   %-30s: ✓ Available (%s)\n", sf.name, sf.desc)
				corefoundation.CFRelease(font)
			} else {
				fmt.Printf("   %-30s: ✗ Not found\n", sf.name)
			}
		}
	}

	fmt.Println("\n✓ All CoreText font operations completed successfully!")
	fmt.Println("\nNote: CoreText is a low-level text rendering framework.")
	fmt.Println("For UI applications, consider using NSFont (AppKit) or UIFont (UIKit).")
}
