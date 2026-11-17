//go:build tk_integration

package main

import (
	"fmt"
	"testing"

	tk9 "modernc.org/tk9.0"
	"modernc.org/tk9.0/extensions/eval"
)

func TestImageLayout(t *testing.T) {
	canvas := createUI()
	t.Cleanup(func() { tk9.Destroy(tk9.App) })
	tk9.Update()
	for key, extension := range tk9.Extensions {
		if key.Name == "eval" {
			if err := extension.Initialize(nil); err != nil {
				t.Fatal(err)
			}
		}
	}
	query := fmt.Sprintf("list [winfo x %s] [winfo y %s] [%s canvasx 0] [%s canvasy 0] [%s bbox all]", canvas, canvas, canvas, canvas, canvas)
	result, err := eval.Eval(query)
	if err != nil {
		t.Fatal(err)
	}
	if want := "0 0 0.0 0.0 {0 0 601 631}"; result != want {
		t.Errorf("canvas position, origin and image bounds = %q, want %q", result, want)
	}
}
