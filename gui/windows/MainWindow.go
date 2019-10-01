package windows

import (
	. "github.com/lxn/walk/declarative"
)

type MainWindows struct {
	Initialized bool
	InitWidth   int
	InitHeight  int
	Size        Size
	MinSize     Size
}
