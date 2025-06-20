package logger

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/go-stack/stack"
)

type LevelFilterHandler struct {
	Level LogLevel
	h     LogHandler
}

// Filters out records which do not match the level
// Uses the `log15.FilterHandler` to perform this task.
func LevelHandler(lvl LogLevel, h LogHandler) LogHandler {
	return &LevelFilterHandler{lvl, h}
}

// The implementation of the Log.
func (h LevelFilterHandler) Log(r *Record) error {
	if r.Level == h.Level {
		return h.h.Log(r)
	}
	return nil
}

// Filters out records which do not match the level
// Uses the `log15.FilterHandler` to perform this task.
func MinLevelHandler(lvl LogLevel, h LogHandler) LogHandler {
	return FilterHandler(func(r *Record) (pass bool) {
		return r.Level <= lvl
	}, h)
}

// Filters out records which match the level
// Uses the `log15.FilterHandler` to perform this task.
func NotLevelHandler(lvl LogLevel, h LogHandler) LogHandler {
	return FilterHandler(func(r *Record) (pass bool) {
		return r.Level != lvl
	}, h)
}

func CallerFileHandler(h LogHandler) LogHandler {
	return FuncHandler(func(r *Record) error {
		r.Context.Add("caller", PrettyCallerPrint(r.Call.(stack.Call)))
		return h.Log(r)
	})
}

var isClosureReg = regexp.MustCompile(`(\.func\d+|\-range\d+)(\.\d+)?$`)

func isClosure(fn string) bool {
	return isClosureReg.MatchString(fn)
}
func PrettyStackPrint(cs []stack.Call) string {
	if len(cs) == 1 {
		return cs[0].String()
	}
	var buf strings.Builder
	buf.Grow(len(cs) * 32)
	buf.WriteByte('[')
	for i := len(cs) - 1; i >= 0; i-- {
		buf.WriteString(cs[i].String())
		if i > 0 {
			buf.WriteString("->")
		}
	}
	buf.WriteByte(']')
	return buf.String()
}
func PrettyCallerPrint(c stack.Call) string {
	// For closures, walk three additional stack frames upward and skip frames from the Go standard library
	if isClosure(c.Frame().Function) {
		s := stack.Trace().TrimBelow(c).TrimRuntime()
		stackDepth := 3
		if len(s) < stackDepth {
			stackDepth = len(s)
		}

		s2 := make([]stack.Call, 0, stackDepth)
		for i := 0; i < stackDepth && i < len(s); i++ {
			cs := s[i : i+1]
			if len(cs.TrimRuntime()) != 0 {
				s2 = append(s2, s[i])
			} else {
				// The Go standard-library frames are already skipped, so we need to unwind one more level
				stackDepth++
			}
		}

		if len(s2) > 0 {
			return PrettyStackPrint(s2)
		}
	}
	return c.String()
}

// Adds in a context called `caller` to the record (contains file name and line number like `foo.go:12`)
// Uses the `log15.CallerFuncHandler` to perform this task.
func CallerFuncHandler(h LogHandler) LogHandler {
	return FuncHandler(func(r *Record) error {
		r.Context.Add("fn", fmt.Sprintf("%+n", r.Call))
		return h.Log(r)
	})
}

// Filters out records which match the key value pair
// Uses the `log15.MatchFilterHandler` to perform this task.
func MatchHandler(key string, value interface{}, h LogHandler) LogHandler {
	return MatchFilterHandler(key, value, h)
}

// MatchFilterHandler returns a Handler that only writes records
// to the wrapped Handler if the given key in the logged
// context matches the value. For example, to only log records
// from your ui package:
//
//	log.MatchFilterHandler("pkg", "app/ui", log.StdoutHandler)
func MatchFilterHandler(key string, value interface{}, h LogHandler) LogHandler {
	return FilterHandler(func(r *Record) (pass bool) {
		return r.Context.Data[key] == value
	}, h)
}

// If match then A handler is called otherwise B handler is called.
func MatchAbHandler(key string, value interface{}, a, b LogHandler) LogHandler {
	return FuncHandler(func(r *Record) error {
		if r.Context.Data[key] == value {
			return a.Log(r)
		} else if b != nil {
			return b.Log(r)
		}

		return nil
	})
}

// The nil handler is used if logging for a specific request needs to be turned off.
func NilHandler() LogHandler {
	return FuncHandler(func(r *Record) error {
		return nil
	})
}

// Match all values in map to log.
func MatchMapHandler(matchMap map[string]interface{}, a LogHandler) LogHandler {
	return matchMapHandler(matchMap, false, a)
}

// Match !(Match all values in map to log) The inverse of MatchMapHandler.
func NotMatchMapHandler(matchMap map[string]interface{}, a LogHandler) LogHandler {
	return matchMapHandler(matchMap, true, a)
}

// Rather then chaining multiple filter handlers, process all here.
func matchMapHandler(matchMap map[string]interface{}, inverse bool, a LogHandler) LogHandler {
	return FuncHandler(func(r *Record) error {
		matchCount := 0
		for k, v := range matchMap {
			value, found := r.Context.Data[k]
			if !found {
				return nil
			}
			// Test for two failure cases
			if value == v && inverse || value != v && !inverse {
				return nil
			} else {
				matchCount++
			}
		}
		if matchCount != len(matchMap) {
			return nil
		}
		return a.Log(r)
	})
}

// Filters out records which do not match the key value pair
// Uses the `log15.FilterHandler` to perform this task.
func NotMatchHandler(key string, value interface{}, h LogHandler) LogHandler {
	return FilterHandler(func(r *Record) (pass bool) {
		return r.Context.Data[key] != value
	}, h)
}

func MultiHandler(hs ...LogHandler) LogHandler {
	return FuncHandler(func(r *Record) error {
		for _, h := range hs {
			// what to do about failures?
			h.Log(r)
		}
		return nil
	})
}

// StreamHandler writes log records to an io.Writer
// with the given format. StreamHandler can be used
// to easily begin writing log records to other
// outputs.
//
// StreamHandler wraps itself with LazyHandler and SyncHandler
// to evaluate Lazy objects and perform safe concurrent writes.
func StreamHandler(wr io.Writer, fmtr LogFormat) LogHandler {
	h := FuncHandler(func(r *Record) error {
		_, err := wr.Write(fmtr.Format(r))
		return err
	})
	return LazyHandler(SyncHandler(h))
}

// Filter handler.
func FilterHandler(fn func(r *Record) bool, h LogHandler) LogHandler {
	return FuncHandler(func(r *Record) error {
		if fn(r) {
			return h.Log(r)
		}
		return nil
	})
}

// List log handler handles a list of LogHandlers.
type ListLogHandler struct {
	handlers []LogHandler
}

// Create a new list of log handlers.
func NewListLogHandler(h1, h2 LogHandler) *ListLogHandler {
	ll := &ListLogHandler{handlers: []LogHandler{h1, h2}}
	return ll
}

// Log the record.
func (ll *ListLogHandler) Log(r *Record) (err error) {
	for _, handler := range ll.handlers {
		if err == nil {
			err = handler.Log(r)
		} else {
			handler.Log(r)
		}
	}
	return
}

// Add another log handler.
func (ll *ListLogHandler) Add(h LogHandler) {
	if h != nil {
		ll.handlers = append(ll.handlers, h)
	}
}

// Remove a log handler.
func (ll *ListLogHandler) Del(h LogHandler) {
	if h != nil {
		for i, handler := range ll.handlers {
			if handler == h {
				ll.handlers = append(ll.handlers[:i], ll.handlers[i+1:]...)
			}
		}
	}
}
