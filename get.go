// Package status gathers information about the status of Go packages.
package status

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "os/exec"

// Package contains status information about a package.
type Package struct {
	// Path is the import path of the package.
	Path string
	// CanGet specifies if it's possible to "go get" the package.
	CanGet bool
	// CanBuild specifies if it's possible to "go build" the package.
	CanBuild bool
}

// Get downloads the package and updates it's CanGet status.
func (pkg *Package) Get() (err error) {
	pkg.CanGet = false
	// go get flags:
	//    -d (only download)
	//    -u (update)
	cmd := exec.Command("go", "get", "-d", "-u", pkg.Path)
	cmd.Env = append([]string{"GIT_ASKPASS=/usr/bin/echo"}, os.Environ()...)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Get failed for: %s", pkg.Path)
	}
	pkg.CanGet = true
	return nil
}

// init initializes a custom $GOPATH.
func init() {
	gopath, err := ioutil.TempDir("", "pkg_")
	if err != nil {
		log.Fatalln(err)
	}
	os.Setenv("GOPATH", gopath)
	fmt.Printf("Using '%s' as GOPATH\n", gopath)
}