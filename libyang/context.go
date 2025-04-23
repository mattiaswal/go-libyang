package libyang

/*
#cgo pkg-config: libyang
#include <stdlib.h>
#include <libyang/libyang.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Context struct {
	ptr *C.struct_ly_ctx
}

func NewContext(searchPath string) (*Context, error) {
	var err C.LY_ERR
	var ctx *C.struct_ly_ctx
	cpath := C.CString(searchPath)

	defer C.free(unsafe.Pointer(cpath))

	err = C.ly_ctx_new(cpath, 0, &ctx)
	if err != C.LY_SUCCESS {
		return nil, errors.New("failed to create context")
	}

	return &Context{ptr: ctx}, nil
}

func (c *Context) Free() {
	if c.ptr != nil {
		C.ly_ctx_destroy(c.ptr)
		c.ptr = nil
	}
}

func (c *Context) LoadModule(name, revision string) (*Module, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var crev *C.char
	if revision != "" {
		crev = C.CString(revision)
		defer C.free(unsafe.Pointer(crev))
	}

	mod := C.ly_ctx_load_module(c.ptr, cname, crev, nil)
	if mod == nil {
		return nil, errors.New("failed to load module")
	}
	return &Module{ptr: mod}, nil
}

func (c *Context) GetModule(name string) *Module {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	mod := C.ly_ctx_get_module_latest(c.ptr, cname)
	if mod == nil {
		return nil
	}
	return &Module{ptr: mod}
}
