package rules

import (
	"github.com/yoheimuta/go-protoparser/v4/lexer"
	"github.com/yoheimuta/go-protoparser/v4/lexer/scanner"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/autodisable"
	"github.com/yoheimuta/protolint/linter/fixer"
	"github.com/yoheimuta/protolint/linter/rule"

	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/strs"
	"github.com/yoheimuta/protolint/linter/visitor"
)

type FieldNamesLowerCamelCaseRule struct {
	fixMode         bool
	autoDisableType autodisable.PlacementType
}

// NewFieldNamesLowerCamelCaseRule creates a new FieldNamesLowerCamelCaseRule.
func NewFieldNamesLowerCamelCaseRule() FieldNamesLowerCamelCaseRule {
	return FieldNamesLowerCamelCaseRule{
		fixMode:         false,
		autoDisableType: autodisable.Noop,
	}
}

// ID returns the ID of this rule.
func (r FieldNamesLowerCamelCaseRule) ID() string {
	return "FIELD_NAMES_LOWER_CAMEL_CASE"
}

// Purpose returns the purpose of this rule.
func (r FieldNamesLowerCamelCaseRule) Purpose() string {
	return "Verifies that all field names are lowerCamel case."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r FieldNamesLowerCamelCaseRule) IsOfficial() bool {
	return true
}

func (r FieldNamesLowerCamelCaseRule) Severity() rule.Severity {
	return rule.SeverityWarning
}

// Apply applies the rule to the proto.
func (r FieldNamesLowerCamelCaseRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	base, err := visitor.NewBaseFixableVisitor(r.ID(), r.fixMode, proto, string(r.Severity()))
	if err != nil {
		return nil, err
	}

	v := &FieldNamesLowerCamelCaseVisitor{
		BaseFixableVisitor: base,
	}
	return visitor.RunVisitorAutoDisable(v, proto, r.ID(), r.autoDisableType)
}

type FieldNamesLowerCamelCaseVisitor struct {
	*visitor.BaseFixableVisitor
}

// VisitField checks the field.
func (v *FieldNamesLowerCamelCaseVisitor) VisitField(field *parser.Field) bool {
	name := field.FieldName
	if !strs.IsLowerCamelCase(name) {
		expected := strs.ToLowerCamelCase(name)
		v.AddFailuref(field.Meta.Pos, "Field name %q must be lowerCamel like %q", name, expected)

		err := v.Fixer.SearchAndReplace(field.Meta.Pos, func(lex *lexer.Lexer) fixer.TextEdit {
			lex.NextKeyword()
			switch lex.Token {
			case scanner.TREPEATED, scanner.TREQUIRED, scanner.TOPTIONAL:
			default:
				lex.UnNext()
			}
			parseType(lex)
			lex.Next()
			return fixer.TextEdit{
				Pos:     lex.Pos.Offset,
				End:     lex.Pos.Offset + len(lex.Text) - 1,
				NewText: []byte(expected),
			}
		})
		if err != nil {
			panic(err)
		}
	}
	return false
}

// VisitMapField checks the map field.
func (v *FieldNamesLowerCamelCaseVisitor) VisitMapField(field *parser.MapField) bool {
	name := field.MapName
	if !strs.IsLowerCamelCase(name) {
		expected := strs.ToLowerCamelCase(name)
		v.AddFailuref(field.Meta.Pos, "Field name %q must be lowerCamel like %q", name, expected)

		err := v.Fixer.SearchAndReplace(field.Meta.Pos, func(lex *lexer.Lexer) fixer.TextEdit {
			lex.NextKeyword()
			lex.Next()
			lex.Next()
			lex.Next()
			parseType(lex)
			lex.Next()
			lex.Next()
			return fixer.TextEdit{
				Pos:     lex.Pos.Offset,
				End:     lex.Pos.Offset + len(lex.Text) - 1,
				NewText: []byte(expected),
			}
		})
		if err != nil {
			panic(err)
		}
	}
	return false
}

// VisitOneofField checks the oneof field.
func (v *FieldNamesLowerCamelCaseVisitor) VisitOneofField(field *parser.OneofField) bool {
	name := field.FieldName
	if !strs.IsLowerCamelCase(name) {
		expected := strs.ToLowerCamelCase(name)
		v.AddFailuref(field.Meta.Pos, "Field name %q must be lowerCamel like %q", name, expected)

		err := v.Fixer.SearchAndReplace(field.Meta.Pos, func(lex *lexer.Lexer) fixer.TextEdit {
			parseType(lex)
			lex.Next()
			return fixer.TextEdit{
				Pos:     lex.Pos.Offset,
				End:     lex.Pos.Offset + len(lex.Text) - 1,
				NewText: []byte(expected),
			}
		})
		if err != nil {
			panic(err)
		}
	}
	return false
}
