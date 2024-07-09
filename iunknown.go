package ole

import (
	"syscall"
	"unsafe"
)

type IUnknown struct {
	// RawVTable *[1024]uintptr
	RawVTable *interface{}
}

type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type UnknownLike interface {
	QueryInterface(iid *GUID) (disp *IDispatch, err error)
	AddRef() int32
	Release() int32
}

type IUnknownInterface interface {
	QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) HRESULT
	AddRef() uint32
	Release() uint32
}

func (v *IUnknown) VTable() *IUnknownVtbl {
	return (*IUnknownVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUnknown) PutQueryInterface(interfaceID *GUID, obj interface{}) error {
	return reflectQueryInterface(v, v.VTable().QueryInterface, interfaceID, obj)
}

func (v *IUnknown) IDispatch(interfaceID *GUID) (dispatch *IDispatch, err error) {
	err = v.PutQueryInterface(interfaceID, &dispatch)
	return
}

func (v *IUnknown) IEnumVARIANT(interfaceID *GUID) (enum *IEnumVARIANT, err error) {
	err = v.PutQueryInterface(interfaceID, &enum)
	return
}

func (v *IUnknown) QueryInterface(iid *GUID) (*IDispatch, error) {
	return queryInterface(v, iid)
}

func (v *IUnknown) MustQueryInterface(iid *GUID) (disp *IDispatch) {
	unk, err := queryInterface(v, iid)
	if err != nil {
		panic(err)
	}
	return unk
}

func (v *IUnknown) AddRef() uintptr {
	return addRef(v)
}

func (v *IUnknown) Release() uintptr {
	return release(v)
}
