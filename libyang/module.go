// File: module.go
package libyang

/*
#include <libyang/libyang.h>
*/
import "C"

type Module struct {
    ptr *C.struct_lys_module
}

func (m *Module) Name() string {
    return C.GoString(m.ptr.name)
}

func (m *Module) Revision() string {
    return C.GoString(m.ptr.revision)
}
