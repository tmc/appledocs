// Code generated from Apple documentation for AppleArchive. DO NOT EDIT.

package applearchive
import (
	"unsafe"
)


// C struct types
// AAAccessControlEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAAccessControlEntry
type AAAccessControlEntry struct {
	Flags unsafe.Pointer
	Perms unsafe.Pointer
	Qualifier_type unsafe.Pointer
	Tag unsafe.Pointer
}/* debug [types.gen.go/struct]: AAAccessControlEntry */

// AAEntryAttributes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryAttributes
type AAEntryAttributes struct {
	Bits uint32
	Btm unsafe.Pointer
	BTM uint32
	Ctm unsafe.Pointer
	CTM uint32
	Flg uint32
	FLG uint32
	Gid uint32
	GID uint32
	Mod uint32
	MOD uint32
	Mtm unsafe.Pointer
	MTM uint32
	UID uint32
	Uid uint32
}/* debug [types.gen.go/struct]: AAEntryAttributes */





