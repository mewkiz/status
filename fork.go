package status

import (
	"fmt"
	"path"
	"strings"

	"github.com/mewfork/go-github/github"
)

// CheckFork uses the GitHub API to check if the package is a fork and updates
// its IsFork status accordingly.
func (pkg *Package) CheckFork() (err error) {
	pkg.IsFork = false
	const prefix = "github.com/"
	if !strings.HasPrefix(pkg.Path, prefix) {
		return fmt.Errorf("Package.CheckFork: invalid prefix; expected %q, got %q", prefix, pkg.Path)
	}
	uname, rname := path.Split(pkg.Path[len(prefix):])
	repo, _, err := client.Repositories.Get(uname, rname)
	if err != nil {
		return err
	}
	if *repo.Fork {
		pkg.IsFork = true
	}
	return nil
}

// client is used to communicate with the GitHub API.
var client *github.Client

// init initializes a GitHub client.
func init() {
	client = github.NewClient(nil)
}
