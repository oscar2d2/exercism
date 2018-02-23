package paasio

import (
	"io"
	"sync"
)

type Counter struct {
	mux  sync.Mutex
	w    io.Writer
	r    io.Reader
	n    int64
	nops int
}

func (c *Counter) Write(p []byte) (n int, err error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	n, err = c.w.Write(p)
	if err == nil {
		c.n += int64(n)
		c.nops++
	}

	return
}

func (c *Counter) WriteCount() (n int64, nops int) {
	return c.n, c.nops
}

func (c *Counter) Read(p []byte) (n int, err error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	n, err = c.r.Read(p)
	if err == nil {
		c.n += int64(n)
		c.nops++
	}

	return
}

func (c *Counter) ReadCount() (n int64, nops int) {
	return c.n, c.nops
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &Counter{w: w}
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &Counter{r: r}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &Counter{r: rw, w: rw}
}
