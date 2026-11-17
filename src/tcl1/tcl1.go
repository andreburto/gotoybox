package main

import (
	"strconv"

	tk9 "modernc.org/tk9.0"
)

func main() {
	value := 0
	entry := tk9.Entry(tk9.Textvariable(strconv.Itoa(value)))
	buttons := tk9.Frame()

	tk9.Pack(entry, tk9.Fill("x"))

	decrease := buttons.Button(tk9.Txt("-"), tk9.Command(func() {
		value--
		entry.Configure(tk9.Textvariable(strconv.Itoa(value)))
	}))
	tk9.Pack(decrease, tk9.Side("left"), tk9.Expand(true), tk9.Fill("x"))

	increase := buttons.Button(tk9.Txt("+"), tk9.Command(func() {
		value++
		entry.Configure(tk9.Textvariable(strconv.Itoa(value)))
	}))
	tk9.Pack(increase, tk9.Side("right"), tk9.Expand(true), tk9.Fill("x"))

	tk9.Pack(buttons, tk9.Fill("x"))
	tk9.App.Wait()
}
