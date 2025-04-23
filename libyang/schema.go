// File: schema.go
package libyang

/*
#include <libyang/libyang.h>
*/
import "C"

type SchemaNode struct {
    ptr *C.struct_lysc_node
}

func (s *SchemaNode) Name() string {
    return C.GoString(s.ptr.name)
}

func (s *SchemaNode) Nodetype() int {
    return int(s.ptr.nodetype)
}