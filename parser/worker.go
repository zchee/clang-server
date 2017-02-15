// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"runtime"
	"sync"
)

// dispatcher represents a management workers.
type dispatcher struct {
	pool    chan *worker
	queue   chan parseArg
	workers []*worker
	wg      sync.WaitGroup
	quit    chan struct{}
}

// worker represents the worker that executes the job.
type worker struct {
	dispatcher *dispatcher
	data       chan interface{}
	quit       chan struct{}
	fn         func(parseArg) error
}

const maxQueues = 10000

// newDispatcher returns a pointer of dispatcher.
func newDispatcher(fn func(parseArg) error) *dispatcher {
	d := &dispatcher{
		pool:  make(chan *worker, runtime.NumCPU()+1),
		queue: make(chan parseArg, maxQueues),
		quit:  make(chan struct{}),
	}
	d.workers = make([]*worker, cap(d.pool))
	for i := 0; i < cap(d.pool); i++ {
		w := worker{
			dispatcher: d,
			data:       make(chan interface{}),
			quit:       make(chan struct{}),
			fn:         fn,
		}
		d.workers[i] = &w
	}
	return d
}

// Add adds a given value to the queue of the dispatcher.
func (d *dispatcher) Add(v parseArg) {
	d.wg.Add(1)
	d.queue <- v
}

// Start starts the specified dispatcher but does not wait for it to complete.
func (d *dispatcher) Start() {
	for _, w := range d.workers {
		w.start()
	}
	go func() {
		for {
			select {
			case v := <-d.queue:
				(<-d.pool).data <- v

			case <-d.quit:
				return
			}
		}
	}()
}

// Wait waits for the dispatcher to exit. It must have been started by Start.
func (d *dispatcher) Wait() {
	d.wg.Wait()
}

// Stop stops the dispatcher to execute. The dispatcher stops gracefully
// if the given boolean is false.
func (d *dispatcher) Stop(immediately bool) {
	if !immediately {
		d.Wait()
	}

	d.quit <- struct{}{}
	for _, w := range d.workers {
		w.quit <- struct{}{}
	}
}

func (w *worker) start() {
	go func() {
		for {
			// register the current worker into the dispatch pool
			w.dispatcher.pool <- w

			select {
			case v := <-w.data:
				if arg, ok := v.(parseArg); ok {
					w.fn(arg)
				}

				w.dispatcher.wg.Done()

			case <-w.quit:
				return
			}
		}
	}()
}
