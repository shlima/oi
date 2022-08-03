package closer

import "os"

var Global = New(os.Interrupt)

func Add(fn CloseFn) {
	Global.Add(fn)
	Global.AutoWatchAsync()
}

func AddE(fn CloseFnE) {
	Global.AddE(fn)
	Global.AutoWatchAsync()
}
