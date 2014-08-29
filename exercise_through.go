package main

import "fmt"

type write func(data int)
type end func(args interface{})
type events map[string]interface{}

type Stream struct {
	readable               bool
	writable               bool
	paused                 bool
	autoDestroy            bool
	write                  func(data int) bool
	queue                  func(data ...int)
	destroy, pause, resume func()
	end                    func(arg ...int)
	events
}

func (s *Stream) on(name string, data interface{}) {

}

func (s *Stream) emit(name string, data ...interface{}) interface{} {
	return nil
}

func Through(w write, e end, opts ...bool) *Stream {

	var (
		ended       = false
		destroyed   = false
		buffer      = []int{}
		_ended      = false
		autoDestroy = false
		s           *Stream
	)
	if opts != nil {
		autoDestroy = true
	} else {
		autoDestroy = !opts[0]
	}

	s = &Stream{
		readable:    true,
		writable:    true,
		paused:      false,
		autoDestroy: autoDestroy,
	}

	s.write = func(data int) bool {
		w(data)
		return !s.paused
	}

	var drain = func() interface{} {
		for len(buffer) > 0 && !s.paused {
			data := buffer[0:1]

			if data == nil {
				return s.emit("end")
			} else {
				return s.emit("data", data)
			}
		}
		return nil
	}

	s.queue = func(data ...int) {
		if _ended {
			return
		}
		var d int
		if len(data) == 0 {
			_ended = true
			d = 0
		} else {
			d = data[0]
		}
		buffer = append(buffer, d)
		drain()
	}

	s.on("end", func() {
		s.readable = false
		if !s.writable && s.autoDestroy {
			go s.destroy()
		}
	})

	var _end = func() {
		s.writable = false
		e(s)
		if !s.writable && s.autoDestroy {
			go s.destroy()
		}
	}

	s.end = func(arg ...int) {
		if ended {
			return
		}

		ended = true
		if len(arg) > 0 {
			s.write(arg[0])
		}
		_end()
	}

	s.destroy = func() {
		if destroyed {
			return
		}
		destroyed = true
		ended = true
		buffer = []int{}
		s.writable = false
		s.readable = false
		s.emit("close")
	}

	s.pause = func() {
		if s.paused {
			return
		}
		s.paused = true
	}

	s.resume = func() {
		if s.paused {
			s.paused = false
			s.emit("resume")
		}
		drain()
		//may have become paused again,
		//as drain emits 'data'.
		if !s.paused {
			s.emit("drain")
		}
	}

	fmt.Println(s)

	w(1)
	return s
}

func main() {
	fmt.Println("Hello, 世界")

	Through(func(data int) {
		fmt.Println("call")
	}, func(args interface{}) {

	}, true)
}
