package a

import "sync"

// Plain holds no no-copy field; its pointer-receiver method is flagged.
type Plain struct {
	count int
	err   error
}

func (p *Plain) Inc() { p.count++ } // want `pointer receiver`

// Value receivers are always allowed.
func (p Plain) Get() int { return p.count }

// Guarded holds a sync.Mutex, so a pointer receiver is allowed.
type Guarded struct {
	mu   sync.Mutex
	data int
}

func (g *Guarded) Set(n int) {
	g.mu.Lock()
	g.data = n
	g.mu.Unlock()
}

// Nested transitively contains a Mutex through Guarded, so it is allowed.
type Nested struct {
	inner Guarded
}

func (n *Nested) Touch() { n.inner.Set(1) }

// Scalar is a non-struct named type; its pointer-receiver method is flagged.
type Scalar int

func (s *Scalar) Bump() { *s++ } // want `pointer receiver`

// plainFunc is not a method.
func plainFunc() {}
