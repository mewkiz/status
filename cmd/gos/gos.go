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
	{Path: "code.google.com/p/cascadia"}, // +get +build
	{Path: "github.com/404user/404repo"}, // -get
	{Path: "github.com/BurntSushi/xgb"},  // +get +build
	{Path: "github.com/burntsushi/xgb"},  // -get
	{Path: "github.com/mewmew/broken"},   // +get -build
}

func gos() (err error) {
	for _, pkg := range pkgs {
		// get
		fmt.Println("get started:", pkg.Path)
		err = pkg.Get()
		if err != nil {
			log.Println(err)
			fmt.Println()
			continue
		} else {
			fmt.Println("get succeeded for:", pkg.Path)
		}
		fmt.Println()
		// build
		fmt.Println("build started:", pkg.Path)
		err = pkg.Build()
		if err != nil {
			log.Println(err)
			fmt.Println()
			continue
		} else {
			fmt.Println("build succeeded for:", pkg.Path)
		}
		fmt.Println()
	}
	var getable, buildable int
	for _, pkg := range pkgs {
		// count how many packages that are "go getable".
		if pkg.CanGet {
			getable++
		}
		// count how many packages that are "go buildable".
		if pkg.CanBuild {
			buildable++
		}
	}
	fmt.Printf("%d/%d packages are \"go getable\".\n", getable, len(pkgs))
	fmt.Printf("%d/%d packages are \"go buildable\".\n", buildable, len(pkgs))
	return nil
}
