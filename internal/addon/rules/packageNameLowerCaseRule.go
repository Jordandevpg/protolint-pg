package rules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"

	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/strs"
	"github.com/yoheimuta/protolint/linter/visitor"
)

// PackageNameLowerCaseRule verifies that the package name doesn't contain any uppercase letters.
// See https://developers.google.com/protocol-buffers/docs/style#packages.
type PackageNameLowerCaseRule struct{}

// NewPackageNameLowerCaseRule creates a new PackageNameLowerCaseRule.
func NewPackageNameLowerCaseRule() PackageNameLowerCaseRule {
	return PackageNameLowerCaseRule{}
}

// ID returns the ID of this rule.
func (r PackageNameLowerCaseRule) ID() string {
	return "PACKAGE_NAME_LOWER_CASE"
}

// Purpose returns the purpose of this rule.
func (r PackageNameLowerCaseRule) Purpose() string {
	return "Verifies that the package name doesn't contain any uppercase letters."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r PackageNameLowerCaseRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r PackageNameLowerCaseRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &packageNameLowerCaseVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID()),
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type packageNameLowerCaseVisitor struct {
	*visitor.BaseAddVisitor
}

// VisitPackage checks the package.
func (v *packageNameLowerCaseVisitor) VisitPackage(p *parser.Package) bool {
	if !isPackageLowerCase(p.Name) {
		v.AddFailuref(p.Meta.Pos, "Package name %q must not contain any uppercase letter.", p.Name)
	}
	return false
}

func isPackageLowerCase(packageName string) bool {
	return !strs.HasAnyUpperCase(packageName)
}
