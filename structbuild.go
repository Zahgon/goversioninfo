package goversioninfo

import (
	"reflect"
)

// *****************************************************************************
// Structure Building
// *****************************************************************************

/*
Version Information Structures
http://msdn.microsoft.com/en-us/library/windows/desktop/ff468916.aspx

VersionInfo Names
http://msdn.microsoft.com/en-us/library/windows/desktop/aa381058.aspx#string-name

Translation: LangID
http://msdn.microsoft.com/en-us/library/windows/desktop/aa381058.aspx#langid

Translation: CharsetID
http://msdn.microsoft.com/en-us/library/windows/desktop/aa381058.aspx#charsetid

*/

// VSVersionInfo is the top level version container.
type VSVersionInfo struct {
	WLength      uint16
	WValueLength uint16
	WType        uint16
	SzKey        []byte
	Padding1     []byte
	Value        VSFixedFileInfo
	Padding2     []byte
	Children     VSStringFileInfo
	Children2    VSVarFileInfo
}

// VSFixedFileInfo - most of these should be left at the defaults.
type VSFixedFileInfo struct {
	DwSignature        uint32
	DwStrucVersion     uint32
	DwFileVersionMS    uint32
	DwFileVersionLS    uint32
	DwProductVersionMS uint32
	DwProductVersionLS uint32
	DwFileFlagsMask    uint32
	DwFileFlags        uint32
	DwFileOS           uint32
	DwFileType         uint32
	DwFileSubtype      uint32
	DwFileDateMS       uint32
	DwFileDateLS       uint32
}

// VSStringFileInfo holds multiple collections of keys and values,
// only allows for 1 collection in this package.
type VSStringFileInfo struct {
	WLength      uint16
	WValueLength uint16
	WType        uint16
	SzKey        []byte
	Padding      []byte
	Children     VSStringTable
}

// VSStringTable holds a collection of string keys and values.
type VSStringTable struct {
	WLength      uint16
	WValueLength uint16
	WType        uint16
	SzKey        []byte
	Padding      []byte
	Children     []VSString
}

// VSString holds the keys and values.
type VSString struct {
	WLength      uint16
	WValueLength uint16
	WType        uint16
	SzKey        []byte
	Padding      []byte
	Value        []byte
}

// VSVarFileInfo holds the translation collection of 1.
type VSVarFileInfo struct {
	WLength      uint16
	WValueLength uint16
	WType        uint16
	SzKey        []byte
	Padding      []byte
	Value        VSVar
}

// VSVar holds the translation key.
type VSVar struct {
	WLength      uint16
	WValueLength uint16
	WType        uint16
	SzKey        []byte
	Padding      []byte
	Value        uint32
}

func buildString(i int, v reflect.Value) (VSString, bool) {
	_ = "STUB: not implemented"
	return *new(VSString), false
}

// If the value is set

// 0 for binary, 1 for text

// Create key

// Align to 32-bit boundary

// Align zeros to 32-bit boundary

// Create value

// Length of text in words (2 bytes) plus zero terminate word

// Length of structure
//ss.WLength = 6 + uint16(soFar) + (ss.WValueLength * 2)

func buildStringTable(vi *VersionInfo) VSStringTable {
	_ = "STUB: not implemented"
	return *

	// Always set to 0
	new(VSStringTable)
}

// 0 for binary, 1 for text

// Language identifier and Code page

// Align to 32-bit boundary

// Loop through the struct fields

// If the struct is valid

func buildStringFileInfo(vi *VersionInfo) VSStringFileInfo {
	_ = "STUB: not implemented"
	return *

	// Always set to 0
	new(VSStringFileInfo)
}

// 0 for binary, 1 for text

// Align to 32-bit boundary

// Allows for more than one string table

func buildVar(vfi VarFileInfo) VSVar {
	_ = "STUB: not implemented"

	// 0 for binary, 1 for text
	return *new(VSVar)
}

// Create key

// Align to 32-bit boundary

// Create value

// Length of text in bytes

// Length of structure

func buildVarFileInfo(vfi VarFileInfo) VSVarFileInfo {
	_ = "STUB: not implemented"
	return *

	// Always set to 0
	new(VSVarFileInfo)
}

// 0 for binary, 1 for text

// Align to 32-bit boundary

// TODO Allow for more than one var table

func buildFixedFileInfo(vi *VersionInfo) VSFixedFileInfo {
	_ = "STUB: not implemented"
	return *new(VSFixedFileInfo)
}

// According to the spec, these should be zero...ugh
/*if vi.Timestamp {
	now := syscall.NsecToFiletime(time.Now().UnixNano())
	ff.DwFileDateMS = now.HighDateTime
	ff.DwFileDateLS = now.LowDateTime
}*/

// Build fills the structs with data from the config file
func (v *VersionInfo) Build() { _ = "STUB: not implemented"; return }

// 0 for binary, 1 for text

// Align to 32-bit boundary
// 6 is for the size of WLength, WValueLength, and WType (each is 1 word or 2 bytes: FF FF)

// Length of VSFixedFileInfo (always the same)

// Never needs padding, not included in WLength

// Build strings

// Build translation

// Calculate the total size
