package suba

import (
	"fmt"
	"strings"
)

func Args(s string) []string { return strings.Split(s, " ") }

func ExampleCompound() {
	raise := NewCompound()
	lower := NewCompound()

	raise.Set(1, func(amount string) (e error) {
		fmt.Printf("Raising volume by %s dB (Default).\n", amount)
		return
	})
	raise.Assign("db", func(amount string) (e error) {
		fmt.Printf("Raising volume by %s dB.\n", amount)
		return
	})
	raise.Assign("percent", func(amount string) (e error) {
		fmt.Printf("Raising volume by %s percent.\n", amount)
		return
	})

	lower.Set(1, func(amount string) (e error) {
		fmt.Printf("Lowering volume by %s dB (Default).\n", amount)
		return
	})
	lower.Assign("db", func(amount string) (e error) {
		fmt.Printf("Lowering volume by %s dB.\n", amount)
		return
	})
	lower.Assign("percent", func(amount string) (e error) {
		fmt.Printf("Lowering volume by %s percent.\n", amount)
		return
	})

	root := Keyed{
		"volume": Keyed{
			"raise": raise,
			"lower": lower,
		},
	}

	root.Handle(Args("volume raise 10")...)
	root.Handle(Args("volume raise percent 10")...)
	root.Handle(Args("volume raise db 10")...)
	root.Handle(Args("volume lower 10")...)
	root.Handle(Args("volume lower percent 10")...)
	root.Handle(Args("volume lower db 10")...)
	// Output: Raising volume by 10 dB (Default).
	// Raising volume by 10 percent.
	// Raising volume by 10 dB.
	// Lowering volume by 10 dB (Default).
	// Lowering volume by 10 percent.
	// Lowering volume by 10 dB.
}

func ExampleKeyed() {
	files := map[string]bool{
		"index.html":   true,
		"main.go":      true,
		".main.go.swp": false,
	}
	f_keys := []string{"index.html", "main.go", ".main.go.swp"}
	current := func() []string {
		fs := []string{}
		for _, key := range f_keys {
			if file, ok := files[key]; ok && file {
				fs = append(fs, key)
			}
		}
		return fs
	}

	crud := Keyed{}
	crud.Assign("create", func(name string) (e error) {
		files[name] = true
		f_keys = append(f_keys, name)
		fmt.Printf("Created %s\n", name)
		return
	})
	crud.Assign("read", func() (e error) {
		fmt.Printf("Current files are %v\n", current())
		return
	})
	crud.Assign("update", func(name string) (e error) {
		if file, ok := files[name]; ok {
			files[name] = !file
			fmt.Printf("Updated %s to %v\n", name, files[name])
			return
		}
		fmt.Printf("Could not update %s, does not exist\n", name)
		return
	})
	crud.Assign("delete", func(name string) {
		if file, ok := files[name]; ok && file {
			delete(files, name)
			fmt.Printf("Deleted %s\n", name)
			return
		}
		fmt.Printf("Could not delete %s, does not exist\n", name)
		return
	})

	crud.Handle(Args("read")...)
	crud.Handle(Args("create suba.go")...)
	crud.Handle(Args("read")...)
	crud.Handle(Args("update .main.go.swp")...)
	crud.Handle(Args("read")...)
	crud.Handle(Args("update .main.go.swp")...)
	crud.Handle(Args("delete .main.go.swp")...)
	// Output: Current files are [index.html main.go]
	// Created suba.go
	// Current files are [index.html main.go suba.go]
	// Updated .main.go.swp to true
	// Current files are [index.html main.go .main.go.swp suba.go]
	// Updated .main.go.swp to false
	// Could not delete .main.go.swp, does not exist

}
