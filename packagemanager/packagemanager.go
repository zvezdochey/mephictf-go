package packagemanager

import "errors"

var (
	ErrCircularDependency = errors.New("circular dependency detected")
	ErrDependencyNotFound = errors.New("dependency not found")
	ErrVersionConflict    = errors.New("version conflict")
)

type Package struct {
	Name    string
	Version int
}

// Represents the repository of packages, in which each package has it's own dependencies.
type Repository struct {
	PackageDependencies map[Package][]Package
}

// Calculates the order in which packages from repository need to be installed to install all
// of the required packages along with their dependencies. If package has any dependencies,
// they must be installed strictly before the package itself.
//
// Returns ErrCircularDependency if some packages form circular dependency.
//
// Returns ErrDependencyNotFound if dependency for some of the packages are missing.
//
// Returns ErrVersionConflict if multiple packages require different version of the same package
// as a dependency.
func GetInstallationOrder(repo Repository, required []Package) ([]Package, error) {
	return nil, errors.New("not implemented")
}
