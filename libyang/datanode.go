package libyang

/*
#include <libyang/libyang.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type DataNode struct {
	Ptr *C.struct_lyd_node
}

func (c *Context) ParseData(data string, format DataFormat) (*DataNode, error) {
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))

	var root *C.struct_lyd_node
	r := C.lyd_parse_data_mem(c.ptr, cdata, lyd_format(format), 0, 0, &root)
	if r != 0 || root == nil {
		return nil, errors.New("failed to parse data")
	}
	return &DataNode{Ptr: root}, nil
}

func (d *DataNode) Free() {
	if d.Ptr != nil {
		C.lyd_free_all(d.Ptr)
		d.Ptr = nil
	}
}

func (d *DataNode) Next() DataNode {
	if d.Ptr.next == nil {
		return DataNode{Ptr: nil}
	}
	return DataNode{Ptr: d.Ptr.next}
}

func (d *DataNode) Prev() DataNode {
	return DataNode{Ptr: d.Ptr.prev}
}

func (d *DataNode) Validate(ctx *Context) C.LY_ERR {
	var err C.LY_ERR
	err = C.lyd_validate_all(&d.Ptr, ctx.ptr, 0, nil)
	return err
}
func (d *DataNode) Print(format DataFormat) (string, error) {
	var out *C.char

	ret := C.lyd_print_mem(&out, d.Ptr, lyd_format(format), 0)
	defer C.free(unsafe.Pointer(out))
	if ret != 0 {
		return "", errors.New("failed to print data")
	}

	return C.GoString(out), nil
}

func (d *DataNode) Child() DataNode {
	child := C.lyd_child(d.Ptr)

	return DataNode{Ptr: child}
}

func (d *DataNode) Name() string {
	if d == nil || d.Ptr == nil {
		return ""
	}
	return C.GoString(d.Ptr.schema.name)
}

func (d *DataNode) Value() string {
	if d == nil || d.Ptr == nil {
		return ""
	}
	return C.GoString(C.lyd_get_value(d.Ptr))
}
func (d *DataNode) FirstSibling() DataNode {
	return DataNode{Ptr: C.lyd_first_sibling(d.Ptr)}
}

// Helper functions

// Required for creating DataNode from go-sysrepo
func NewNode(n unsafe.Pointer) DataNode {
	var dnode DataNode
	node := (*C.struct_lyd_node)(n)
	dnode.Ptr = node

	return dnode
}

func (d *DataNode) ChildValue(name string) string {
	child := d.ChildByName(name)
	return child.Value()
}

func (d *DataNode) ChildByName(name string) DataNode {

	for it := d.Child(); it.Ptr != nil; it = it.Next() {
		if it.Ptr.schema != nil {
			if C.GoString(it.Ptr.schema.name) == name {
				return it
			}
		}
	}
	return DataNode{Ptr: nil}
}
