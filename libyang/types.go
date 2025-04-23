// File: types.go
package libyang

/*
#include <libyang/libyang.h>
*/
import "C"

type DataFormat int

const (
	DataFormatXML  DataFormat = iota // LYD_XML
	DataFormatJSON                   // LYD_JSON
)

func lyd_format(format DataFormat) C.LYD_FORMAT {
	switch format {
	case DataFormatXML:
		return C.LYD_XML
	case DataFormatJSON:
		return C.LYD_JSON
	default:
		return C.LYD_XML
	}
}
