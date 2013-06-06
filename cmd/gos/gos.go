// gos - Go documentation status.
package main

import (
	"fmt"
	"log"

	"github.com/mewkiz/status"
)

func main() {
	err := gos()
	if err != nil {
		log.Fatalln(err)
	}
}

var pkgs = []*status.Package{
	{Path: "code.google.com/p/cascadia"},     // +get +parse +build
	{Path: "github.com/404user/404repo"},     // -get
	{Path: "github.com/BurntSushi/xgb"},      // +get +parse +build
	{Path: "github.com/burntsushi/xgb"},      // -get
	{Path: "github.com/mewmew/broken/get1"},  // -get
	{Path: "github.com/mewmew/broken/get2"},  // -get
	{Path: "github.com/mewmew/broken/parse"}, // +get -parse
	{Path: "github.com/mewmew/broken/build"}, // +get +parse -build
}

func gos() (err error) {
	for _, pkg := range pkgs {
		fmt.Printf("--- [ import path: %s ] ---\n", pkg.Path)
		fmt.Println()
		// get
		fmt.Println("get started.")
		err = pkg.Get()
		if err != nil {
			log.Println(err)
			fmt.Println()
			continue
		} else {
			fmt.Println("get succeeded.")
		}
		fmt.Println()
		// parse
		fmt.Println("parse started.")
		err = pkg.Parse()
		if err != nil {
			log.Println(err)
			fmt.Println()
			continue
		} else {
			fmt.Println("parse succeeded.")
		}
		fmt.Println()
		// build
		fmt.Println("build started.")
		err = pkg.Build()
		if err != nil {
			log.Println(err)
			fmt.Println()
			continue
		} else {
			fmt.Println("build succeeded.")
		}
		fmt.Println()
	}
	var getable, parsable, buildable int
	for _, pkg := range pkgs {
		// count how many packages that are "go getable".
		if pkg.CanGet {
			getable++
		}
		// count how many packages that are "parsable".
		if pkg.CanParse {
			parsable++
		}
		// count how many packages that are "go buildable".
		if pkg.CanBuild {
			buildable++
		}
	}
	fmt.Printf("%d/%d packages are \"go getable\".\n", getable, len(pkgs))
	fmt.Printf("%d/%d packages are \"parsable\".\n", parsable, len(pkgs))
	fmt.Printf("%d/%d packages are \"go buildable\".\n", buildable, len(pkgs))
	return nil
}
