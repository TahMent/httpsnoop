// +build go1.8
// Code generated by "httpsnoop/codegen"; DO NOT EDIT

package httpsnoop

import (
	"bufio"
	"io"
	"net"
	"net/http"
)

// HeaderFunc is part of the http.ResponseWriter interface.
type HeaderFunc func() http.Header

// WriteHeaderFunc is part of the http.ResponseWriter interface.
type WriteHeaderFunc func(code int)

// WriteFunc is part of the http.ResponseWriter interface.
type WriteFunc func(b []byte) (int, error)

// FlushFunc is part of the http.Flusher interface.
type FlushFunc func()

// CloseNotifyFunc is part of the http.CloseNotifier interface.
type CloseNotifyFunc func() <-chan bool

// HijackFunc is part of the http.Hijacker interface.
type HijackFunc func() (net.Conn, *bufio.ReadWriter, error)

// ReadFromFunc is part of the io.ReaderFrom interface.
type ReadFromFunc func(src io.Reader) (int64, error)

// PushFunc is part of the http.Pusher interface.
type PushFunc func(target string, opts *http.PushOptions) error

// Hooks defines a set of method interceptors for methods included in
// http.ResponseWriter as well as some others. You can think of them as
// middleware for the function calls they target. See Wrap for more details.
type Hooks struct {
	Header      func(HeaderFunc) HeaderFunc
	WriteHeader func(WriteHeaderFunc) WriteHeaderFunc
	Write       func(WriteFunc) WriteFunc
	Flush       func(FlushFunc) FlushFunc
	CloseNotify func(CloseNotifyFunc) CloseNotifyFunc
	Hijack      func(HijackFunc) HijackFunc
	ReadFrom    func(ReadFromFunc) ReadFromFunc
	Push        func(PushFunc) PushFunc
}

type Unwrapper interface {
	Unwrap() http.ResponseWriter
}

// combination 1/32
type s0 struct {
	http.ResponseWriter
}

// Unwrap returns the underlying http.ResponseWriter
func (s s0) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 2/32
type s1 struct {
	http.ResponseWriter
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s1) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 3/32
type s2 struct {
	http.ResponseWriter
	io.ReaderFrom
}

// Unwrap returns the underlying http.ResponseWriter
func (s s2) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 4/32
type s3 struct {
	http.ResponseWriter
	io.ReaderFrom
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s3) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 5/32
type s4 struct {
	http.ResponseWriter
	http.Hijacker
}

// Unwrap returns the underlying http.ResponseWriter
func (s s4) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 6/32
type s5 struct {
	http.ResponseWriter
	http.Hijacker
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s5) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 7/32
type s6 struct {
	http.ResponseWriter
	http.Hijacker
	io.ReaderFrom
}

// Unwrap returns the underlying http.ResponseWriter
func (s s6) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 8/32
type s7 struct {
	http.ResponseWriter
	http.Hijacker
	io.ReaderFrom
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s7) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 9/32
type s8 struct {
	http.ResponseWriter
	http.CloseNotifier
}

// Unwrap returns the underlying http.ResponseWriter
func (s s8) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 10/32
type s9 struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s9) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 11/32
type s10 struct {
	http.ResponseWriter
	http.CloseNotifier
	io.ReaderFrom
}

// Unwrap returns the underlying http.ResponseWriter
func (s s10) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 12/32
type s11 struct {
	http.ResponseWriter
	http.CloseNotifier
	io.ReaderFrom
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s11) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 13/32
type s12 struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Hijacker
}

// Unwrap returns the underlying http.ResponseWriter
func (s s12) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 14/32
type s13 struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Hijacker
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s13) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 15/32
type s14 struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Hijacker
	io.ReaderFrom
}

// Unwrap returns the underlying http.ResponseWriter
func (s s14) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 16/32
type s15 struct {
	http.ResponseWriter
	http.CloseNotifier
	http.Hijacker
	io.ReaderFrom
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s15) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 17/32
type s16 struct {
	http.ResponseWriter
	http.Flusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s16) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 18/32
type s17 struct {
	http.ResponseWriter
	http.Flusher
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s17) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 19/32
type s18 struct {
	http.ResponseWriter
	http.Flusher
	io.ReaderFrom
}

// Unwrap returns the underlying http.ResponseWriter
func (s s18) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 20/32
type s19 struct {
	http.ResponseWriter
	http.Flusher
	io.ReaderFrom
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s19) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 21/32
type s20 struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
}

// Unwrap returns the underlying http.ResponseWriter
func (s s20) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 22/32
type s21 struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s21) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 23/32
type s22 struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	io.ReaderFrom
}

// Unwrap returns the underlying http.ResponseWriter
func (s s22) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 24/32
type s23 struct {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	io.ReaderFrom
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s23) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 25/32
type s24 struct {
	http.ResponseWriter
	http.Flusher
	http.CloseNotifier
}

// Unwrap returns the underlying http.ResponseWriter
func (s s24) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 26/32
type s25 struct {
	http.ResponseWriter
	http.Flusher
	http.CloseNotifier
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s25) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 27/32
type s26 struct {
	http.ResponseWriter
	http.Flusher
	http.CloseNotifier
	io.ReaderFrom
}

// Unwrap returns the underlying http.ResponseWriter
func (s s26) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 28/32
type s27 struct {
	http.ResponseWriter
	http.Flusher
	http.CloseNotifier
	io.ReaderFrom
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s27) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 29/32
type s28 struct {
	http.ResponseWriter
	http.Flusher
	http.CloseNotifier
	http.Hijacker
}

// Unwrap returns the underlying http.ResponseWriter
func (s s28) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 30/32
type s29 struct {
	http.ResponseWriter
	http.Flusher
	http.CloseNotifier
	http.Hijacker
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s29) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 31/32
type s30 struct {
	http.ResponseWriter
	http.Flusher
	http.CloseNotifier
	http.Hijacker
	io.ReaderFrom
}

// Unwrap returns the underlying http.ResponseWriter
func (s s30) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// combination 32/32
type s31 struct {
	http.ResponseWriter
	http.Flusher
	http.CloseNotifier
	http.Hijacker
	io.ReaderFrom
	http.Pusher
}

// Unwrap returns the underlying http.ResponseWriter
func (s s31) Unwrap() http.ResponseWriter {
	return s.ResponseWriter.(*rw).w
}

// Wrap returns a wrapped version of w that provides the exact same interface
// as w. Specifically if w implements any combination of:
//
// - http.Flusher
// - http.CloseNotifier
// - http.Hijacker
// - io.ReaderFrom
// - http.Pusher
//
// The wrapped version will implement the exact same combination. If no hooks
// are set, the wrapped version also behaves exactly as w. Hooks targeting
// methods not supported by w are ignored. Any other hooks will intercept the
// method they target and may modify the call's arguments and/or return values.
// The CaptureMetrics implementation serves as a working example for how the
// hooks can be used.
func Wrap(w http.ResponseWriter, hooks Hooks) http.ResponseWriter {
	rw := &rw{w: w, h: hooks}
	_, i0 := w.(http.Flusher)
	_, i1 := w.(http.CloseNotifier)
	_, i2 := w.(http.Hijacker)
	_, i3 := w.(io.ReaderFrom)
	_, i4 := w.(http.Pusher)
	switch {
	case !i0 && !i1 && !i2 && !i3 && !i4:
		return s0{rw}
	case !i0 && !i1 && !i2 && !i3 && i4:
		return s1{rw, rw}
	case !i0 && !i1 && !i2 && i3 && !i4:
		return s2{rw, rw}
	case !i0 && !i1 && !i2 && i3 && i4:
		return s3{rw, rw, rw}
	case !i0 && !i1 && i2 && !i3 && !i4:
		return s4{rw, rw}
	case !i0 && !i1 && i2 && !i3 && i4:
		return s5{rw, rw, rw}
	case !i0 && !i1 && i2 && i3 && !i4:
		return s6{rw, rw, rw}
	case !i0 && !i1 && i2 && i3 && i4:
		return s7{rw, rw, rw, rw}
	case !i0 && i1 && !i2 && !i3 && !i4:
		return s8{rw, rw}
	case !i0 && i1 && !i2 && !i3 && i4:
		return s9{rw, rw, rw}
	case !i0 && i1 && !i2 && i3 && !i4:
		return s10{rw, rw, rw}
	case !i0 && i1 && !i2 && i3 && i4:
		return s11{rw, rw, rw, rw}
	case !i0 && i1 && i2 && !i3 && !i4:
		return s12{rw, rw, rw}
	case !i0 && i1 && i2 && !i3 && i4:
		return s13{rw, rw, rw, rw}
	case !i0 && i1 && i2 && i3 && !i4:
		return s14{rw, rw, rw, rw}
	case !i0 && i1 && i2 && i3 && i4:
		return s15{rw, rw, rw, rw, rw}
	case i0 && !i1 && !i2 && !i3 && !i4:
		return s16{rw, rw}
	case i0 && !i1 && !i2 && !i3 && i4:
		return s17{rw, rw, rw}
	case i0 && !i1 && !i2 && i3 && !i4:
		return s18{rw, rw, rw}
	case i0 && !i1 && !i2 && i3 && i4:
		return s19{rw, rw, rw, rw}
	case i0 && !i1 && i2 && !i3 && !i4:
		return s20{rw, rw, rw}
	case i0 && !i1 && i2 && !i3 && i4:
		return s21{rw, rw, rw, rw}
	case i0 && !i1 && i2 && i3 && !i4:
		return s22{rw, rw, rw, rw}
	case i0 && !i1 && i2 && i3 && i4:
		return s23{rw, rw, rw, rw, rw}
	case i0 && i1 && !i2 && !i3 && !i4:
		return s24{rw, rw, rw}
	case i0 && i1 && !i2 && !i3 && i4:
		return s25{rw, rw, rw, rw}
	case i0 && i1 && !i2 && i3 && !i4:
		return s26{rw, rw, rw, rw}
	case i0 && i1 && !i2 && i3 && i4:
		return s27{rw, rw, rw, rw, rw}
	case i0 && i1 && i2 && !i3 && !i4:
		return s28{rw, rw, rw, rw}
	case i0 && i1 && i2 && !i3 && i4:
		return s29{rw, rw, rw, rw, rw}
	case i0 && i1 && i2 && i3 && !i4:
		return s30{rw, rw, rw, rw, rw}
	case i0 && i1 && i2 && i3 && i4:
		return s31{rw, rw, rw, rw, rw, rw}
	}
	panic("unreachable")
}

type rw struct {
	w http.ResponseWriter
	h Hooks
}

func (w *rw) Header() http.Header {
	f := w.w.(http.ResponseWriter).Header
	if w.h.Header != nil {
		f = w.h.Header(f)
	}
	return f()
}

func (w *rw) WriteHeader(code int) {
	f := w.w.(http.ResponseWriter).WriteHeader
	if w.h.WriteHeader != nil {
		f = w.h.WriteHeader(f)
	}
	f(code)
}

func (w *rw) Write(b []byte) (int, error) {
	f := w.w.(http.ResponseWriter).Write
	if w.h.Write != nil {
		f = w.h.Write(f)
	}
	return f(b)
}

func (w *rw) Flush() {
	f := w.w.(http.Flusher).Flush
	if w.h.Flush != nil {
		f = w.h.Flush(f)
	}
	f()
}

func (w *rw) CloseNotify() <-chan bool {
	f := w.w.(http.CloseNotifier).CloseNotify
	if w.h.CloseNotify != nil {
		f = w.h.CloseNotify(f)
	}
	return f()
}

func (w *rw) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	f := w.w.(http.Hijacker).Hijack
	if w.h.Hijack != nil {
		f = w.h.Hijack(f)
	}
	return f()
}

func (w *rw) ReadFrom(src io.Reader) (int64, error) {
	f := w.w.(io.ReaderFrom).ReadFrom
	if w.h.ReadFrom != nil {
		f = w.h.ReadFrom(f)
	}
	return f(src)
}

func (w *rw) Push(target string, opts *http.PushOptions) error {
	f := w.w.(http.Pusher).Push
	if w.h.Push != nil {
		f = w.h.Push(f)
	}
	return f(target, opts)
}

// Unwrap returns the underlying http.ResponseWriter from within zero or more
// layers of httpsnoop wrappers.
func Unwrap(w http.ResponseWriter) http.ResponseWriter {
	if rw, ok := w.(Unwrapper); ok {
		// recurse until rw.Unwrap() returns a non-Unwrapper
		return Unwrap(rw.Unwrap())
	} else {
		return w
	}
}
