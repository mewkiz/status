// gos - Go documentation status.

package main

import "fmt"
import "log"

import "github.com/mewmew/status"

func main() {
	err := gos()
	if err != nil {
		log.Fatalln(err)
	}
}

var pkgs = []*status.Package{
	{Path: "code.google.com/p/cascadia"}, // works
	{Path: "github.com/404user/404repo"}, // fails. Outputs "Username for..."
	{Path: "github.com/BurntSushi/xgb"},  // works
	{Path: "github.com/burntsushi/xgb"},  // fails. Outputs "Username for..."
}

func gos() (err error) {
	// The pkgs slice is defined in links.go for brevity.
	for _, pkg := range pkgs {
		fmt.Println("get started:", pkg.Path)
		err = pkg.Get()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("get succeeded for:", pkg.Path)
		}
		fmt.Println()
	}
	// count how many packages that are "go getable".
	var getable int
	for _, pkg := range pkgs {
		if pkg.CanGet {
			getable++
		}
	}
	fmt.Printf("%d/%d packages are \"go getable\".\n", getable, len(pkgs))
	return nil
}
