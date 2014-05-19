package status

import (
	"fmt"
	"os/exec"
)

// Build compiles the package and updates its CanBuild status.
func (pkg *Package) Build() (err error) {
	pkg.CanBuild = false
	cmd := exec.Command("go", "build", pkg.Path)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Build failed for: %s", pkg.Path)
	}
	pkg.CanBuild = true
	return nil
}
