//go:build !windows

package creacktty

import (
	"os"

	"github.com/creack/pty"

	"github.com/srlehn/termimg/internal"
	"github.com/srlehn/termimg/internal/errors"
	"github.com/srlehn/termimg/term"
)

type ttyCreack struct {
	master   *os.File
	slave    *os.File
	fileName string
}

var _ term.TTY = (*ttyCreack)(nil)
var _ term.TTYProvider = New

func New(ttyFile string) (term.TTY, error) {
	if !internal.IsDefaultTTY(ttyFile) {
		return nil, errors.New(`only default tty supported`)
	}
	p, t, err := pty.Open()
	if err != nil {
		return nil, errors.New(err)
	}
	return &ttyCreack{
		master:   p,
		slave:    t,
		fileName: ttyFile,
	}, nil
}

func (t *ttyCreack) Write(b []byte) (n int, err error) {
	if t == nil {
		return 0, errors.NilReceiver()
	}
	if t.master == nil {
		return 0, errors.New(`nil tty`)
	}
	return t.master.Write(b)
}

func (t *ttyCreack) Read(p []byte) (n int, err error) {
	if t == nil || t.master == nil {
		return 0, errors.NilReceiver()
	}
	return t.master.Read(p)
}
func (t *ttyCreack) TTYDevName() string {
	if t == nil {
		return internal.DefaultTTYDevice()
	}
	return t.fileName
}

func (t *ttyCreack) SizePixel() (cw int, ch int, pw int, ph int, e error) {
	if t == nil || t.master == nil {
		return 0, 0, 0, 0, errors.NilReceiver()
	}
	sz, err := pty.GetsizeFull(t.master)
	if err != nil {
		return 0, 0, 0, 0, errors.New(err)
	}
	return int(sz.Cols), int(sz.Rows), int(sz.X), int(sz.Y), nil
}

// Close ...
func (t *ttyCreack) Close() error {
	if t == nil {
		return nil
	}
	var errM, errS error
	if t.master != nil {
		errM = t.master.Close()
		t.master = nil
	}
	if t.slave != nil {
		errS = t.slave.Close()
		t.slave = nil
	}
	defer func() { t = nil }()
	return errors.Join(errM, errS)
}
