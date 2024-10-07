// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package collections

import (
	"sync"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/internal/kernel32"
)

const GUIDVectorChangedEventHandler string = "0c051752-9fbf-4c70-aa0c-0e4c82d9a761"
const SignatureVectorChangedEventHandler string = "delegate({0c051752-9fbf-4c70-aa0c-0e4c82d9a761})"

type VectorChangedEventHandler struct {
	ole.IUnknown
	sync.Mutex
	refs uintptr
	IID  ole.GUID
}

type VectorChangedEventHandlerVtbl struct {
	ole.IUnknownVtbl
	Invoke uintptr
}

type VectorChangedEventHandlerCallback func(instance *VectorChangedEventHandler, sender *IObservableVector, event *IVectorChangedEventArgs)

var callbacksVectorChangedEventHandler = &vectorChangedEventHandlerCallbacks{
	mu:        &sync.Mutex{},
	callbacks: make(map[unsafe.Pointer]VectorChangedEventHandlerCallback),
}

var releaseChannelsVectorChangedEventHandler = &vectorChangedEventHandlerReleaseChannels{
	mu:    &sync.Mutex{},
	chans: make(map[unsafe.Pointer]chan struct{}),
}

func NewVectorChangedEventHandler(iid *ole.GUID, callback VectorChangedEventHandlerCallback) *VectorChangedEventHandler {
	// create type instance
	size := unsafe.Sizeof(*(*VectorChangedEventHandler)(nil))
	instPtr := kernel32.Malloc(size)
	inst := (*VectorChangedEventHandler)(instPtr)

	// get the callbacks for the VTable
	callbacks := delegate.RegisterCallbacks(instPtr, inst)

	// the VTable should also be allocated in the heap
	sizeVTable := unsafe.Sizeof(*(*VectorChangedEventHandlerVtbl)(nil))
	vTablePtr := kernel32.Malloc(sizeVTable)

	inst.RawVTable = (*interface{})(vTablePtr)

	vTable := (*VectorChangedEventHandlerVtbl)(vTablePtr)
	vTable.IUnknownVtbl = ole.IUnknownVtbl{
		QueryInterface: callbacks.QueryInterface,
		AddRef:         callbacks.AddRef,
		Release:        callbacks.Release,
	}
	vTable.Invoke = callbacks.Invoke

	// Initialize all properties: the malloc may contain garbage
	inst.IID = *iid // copy contents
	inst.Mutex = sync.Mutex{}
	inst.refs = 0

	callbacksVectorChangedEventHandler.add(unsafe.Pointer(inst), callback)

	// See the docs in the releaseChannelsVectorChangedEventHandler struct
	releaseChannelsVectorChangedEventHandler.acquire(unsafe.Pointer(inst))

	inst.addRef()
	return inst
}

func (r *VectorChangedEventHandler) GetIID() *ole.GUID {
	return &r.IID
}

// addRef increments the reference counter by one
func (r *VectorChangedEventHandler) addRef() uintptr {
	r.Lock()
	defer r.Unlock()
	r.refs++
	return r.refs
}

// removeRef decrements the reference counter by one. If it was already zero, it will just return zero.
func (r *VectorChangedEventHandler) removeRef() uintptr {
	r.Lock()
	defer r.Unlock()

	if r.refs > 0 {
		r.refs--
	}

	return r.refs
}

func (instance *VectorChangedEventHandler) Invoke(instancePtr, rawArgs0, rawArgs1, rawArgs2, rawArgs3, rawArgs4, rawArgs5, rawArgs6, rawArgs7, rawArgs8 unsafe.Pointer) uintptr {
	senderPtr := rawArgs0
	eventPtr := rawArgs1

	// See the quote above.
	sender := (*IObservableVector)(senderPtr)
	event := (*IVectorChangedEventArgs)(eventPtr)
	if callback, ok := callbacksVectorChangedEventHandler.get(instancePtr); ok {
		callback(instance, sender, event)
	}
	return ole.S_OK
}

func (instance *VectorChangedEventHandler) AddRef() uintptr {
	return instance.addRef()
}

func (instance *VectorChangedEventHandler) Release() uintptr {
	rem := instance.removeRef()
	if rem == 0 {
		// We're done.
		instancePtr := unsafe.Pointer(instance)
		callbacksVectorChangedEventHandler.delete(instancePtr)

		// stop release channels used to avoid
		// https://github.com/golang/go/issues/55015
		releaseChannelsVectorChangedEventHandler.release(instancePtr)

		kernel32.Free(unsafe.Pointer(instance.RawVTable))
		kernel32.Free(instancePtr)
	}
	return rem
}

type vectorChangedEventHandlerCallbacks struct {
	mu        *sync.Mutex
	callbacks map[unsafe.Pointer]VectorChangedEventHandlerCallback
}

func (m *vectorChangedEventHandlerCallbacks) add(p unsafe.Pointer, v VectorChangedEventHandlerCallback) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.callbacks[p] = v
}

func (m *vectorChangedEventHandlerCallbacks) get(p unsafe.Pointer) (VectorChangedEventHandlerCallback, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.callbacks[p]
	return v, ok
}

func (m *vectorChangedEventHandlerCallbacks) delete(p unsafe.Pointer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.callbacks, p)
}

// typedEventHandlerReleaseChannels keeps a map with channels
// used to keep a goroutine alive during the lifecycle of this object.
// This is required to avoid causing a deadlock error.
// See this: https://github.com/golang/go/issues/55015
type vectorChangedEventHandlerReleaseChannels struct {
	mu    *sync.Mutex
	chans map[unsafe.Pointer]chan struct{}
}

func (m *vectorChangedEventHandlerReleaseChannels) acquire(p unsafe.Pointer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	c := make(chan struct{})
	m.chans[p] = c

	go func() {
		// we need a timer to trick the go runtime into
		// thinking there's still something going on here
		// but we are only really interested in <-c
		t := time.NewTimer(time.Minute)
		for {
			select {
			case <-t.C:
				t.Reset(time.Minute)
			case <-c:
				t.Stop()
				return
			}
		}
	}()
}

func (m *vectorChangedEventHandlerReleaseChannels) release(p unsafe.Pointer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if c, ok := m.chans[p]; ok {
		close(c)
		delete(m.chans, p)
	}
}
