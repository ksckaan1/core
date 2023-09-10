// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on http://github.com/dmarkham/enumer and
// golang.org/x/tools/cmd/stringer:

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package enumgen

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"

	"goki.dev/enums/enumgen/config"
	"goki.dev/gengo"
	"goki.dev/grease"
	"golang.org/x/tools/go/packages"
)

// Generator holds the state of the generator.
// It is primarily used to buffer the output.
type Generator struct {
	Config *config.Config      // The configuration information
	Buf    bytes.Buffer        // The accumulated output.
	Pkgs   []*packages.Package // The packages we are scanning.
	Pkg    *packages.Package   // The packages we are currently on.
	Types  []*Type             // The enum types
}

// NewGenerator returns a new generator with the
// given configuration information.
func NewGenerator(config *Config) *Generator {
	return &Generator{Config: config}
}

// ParsePackage parses the package(s) located in the configuration source directory.
func (g *Generator) ParsePackage() error {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax | packages.NeedTypesInfo,
		// TODO: Need to think about constants in test files. Maybe write type_string_test.go
		// in a separate pass? For later.
		Tests: false,
	}
	pkgs, err := gengo.Load(cfg, g.Config.Dir)
	if err != nil {
		return err
	}
	g.Pkgs = pkgs
	return nil
}

// Printf prints the formatted string to the
// accumulated output in [Generator.Buf]
func (g *Generator) Printf(format string, args ...any) {
	fmt.Fprintf(&g.Buf, format, args...)
}

// PrintHeader prints the header and package clause
// to the accumulated output
func (g *Generator) PrintHeader() {
	// we need a manual import of enums because it is
	// external, but goimports will handle everything else
	gengo.PrintHeader(&g.Buf, g.Pkg.Name, "goki.dev/enums")
}

// FindEnumTypes goes through all of the types in the package
// and finds all integer (signed or unsigned) types labeled with enums:enum
// or enums:bitflag. It stores the resulting types in [Generator.Types].
func (g *Generator) FindEnumTypes() error {
	g.Types = []*Type{}
	gengo.Inspect(g.Pkg, g.InspectForType)
	return nil
}

// AllowedEnumTypes are the types that can be used for enums
// that are not bit flags (bit flags can only be int64s).
// It is stored as a map for quick and convenient access.
var AllowedEnumTypes = map[string]bool{"int": true, "int64": true, "int32": true, "int16": true, "int8": true, "uint": true, "uint64": true, "uint32": true, "uint16": true, "uint8": true}

// InspectForType looks at the given AST node and adds it
// to [Generator.Types] if it is marked with an appropriate
// comment directive. It returns whether the AST inspector should
// continue, and an error if there is one. It should only
// be called in [ast.Inspect].
func (g *Generator) InspectForType(n ast.Node) (bool, error) {
	ts, ok := n.(*ast.TypeSpec)
	if !ok {
		return true, nil
	}
	if ts.Comment == nil {
		return true, nil
	}
	for _, c := range ts.Comment.List {
		tool, directive, args, has, err := grease.ParseDirective(c.Text)
		if err != nil {
			return false, fmt.Errorf("error parsing comment directive %q: %w", c.Text, err)
		}
		if !has {
			continue
		}
		if tool != "enums" {
			continue
		}
		if directive != "enum" && directive != "bitflag" {
			return false, fmt.Errorf("unrecognized enums directive %q (from %q)", directive, c.Text)
		}

		ident, ok := ts.Type.(*ast.Ident)
		if !ok {
			return false, fmt.Errorf("type of enum type (%v) is %T, not *ast.Ident (try using a standard [un]signed integer type instead)", ts.Type, ts.Type)
		}
		cfg := &Config{}
		*cfg = *g.Config
		leftovers, err := grease.SetFromArgs(cfg, args)
		if err != nil {
			return false, fmt.Errorf("error setting config info from comment directive args: %w (from directive %q)", err, c.Text)
		}
		if len(leftovers) > 0 {
			return false, fmt.Errorf("expected 0 positional arguments but got %d (list: %v) (from directive %q)", len(leftovers), leftovers, c.Text)
		}

		typ := g.Pkg.TypesInfo.Defs[ts.Name].Type()
		utyp := typ.Underlying()

		tt := &Type{Name: ts.Name.Name, Type: ts, Config: cfg}
		if ident.String() != utyp.String() { // if our direct type isn't the same as our underlying type, we are extending our direct type
			tt.Extends = ident.String()
		}
		switch directive {
		case "enum":
			if !AllowedEnumTypes[utyp.String()] {
				return false, fmt.Errorf("enum type %s is not allowed; try using a standard [un]signed integer type instead", ident.Name)
			}
			tt.IsBitFlag = false
		case "bitflag":
			if utyp.String() != "int64" {
				return false, fmt.Errorf("bit flag enum type %s is not allowed; bit flag enums must be of type int64", ident.Name)
			}
			tt.IsBitFlag = true
		}
		g.Types = append(g.Types, tt)

	}
	return true, nil
}

// Generate produces the enum methods for the types
// stored in [Generator.Types] and stores them in
// [Generator.Buf]. It returns whether there were
// any enum types to generate methods for, and
// any error that occurred.
func (g *Generator) Generate() (bool, error) {
	if len(g.Types) == 0 {
		return false, nil
	}
	for _, typ := range g.Types {
		values := make([]Value, 0, 100)
		for _, s := range g.Pkg.Syntax {
			if ast.IsGenerated(s) {
				continue
			}
			// Set the state for this run of the walker.
			file := &File{
				Pkg:     g.Pkg,
				File:    s,
				Type:    typ,
				BitFlag: typ.IsBitFlag,
				Values:  nil,
				Config:  typ.Config,
			}
			if file.File != nil {
				var terr error
				ast.Inspect(file.File, func(n ast.Node) bool {
					if terr != nil {
						return false
					}
					cont, err := file.GenDecl(n)
					if err != nil {
						terr = err
					}
					return cont
				})
				if terr != nil {
					return true, fmt.Errorf("Generate: error parsing declaration clauses: %w", terr)
				}
				values = append(values, file.Values...)
			}
		}

		if len(values) == 0 {
			return true, errors.New("no values defined for type " + typ.Name)
		}

		g.TrimValueNames(values, typ.Config)

		err := g.TransformValueNames(values, typ.Config)
		if err != nil {
			return true, fmt.Errorf("error transforming value names: %w", err)
		}

		g.PrefixValueNames(values, typ.Config)

		values = SortValues(values)

		g.BuildBasicMethods(values, typ)
		if typ.IsBitFlag {
			g.BuildBitFlagMethods(values, typ)
		}

		if typ.Config.Text {
			g.BuildTextMethods(values, typ)
		}
		if typ.Config.JSON {
			g.BuildJSONMethods(values, typ)
		}
		if typ.Config.YAML {
			g.BuildYAMLMethods(values, typ)
		}
		if typ.Config.SQL {
			g.AddValueAndScanMethod(typ)
		}
		if typ.Config.GQL {
			g.BuildGQLMethods(values, typ)
		}
	}
	return true, nil
}

// Write formats the data in the the Generator's buffer
// ([Generator.Buf]) and writes it to the file specified by
// [Generator.Config.Output].
func (g *Generator) Write() error {
	return gengo.Write(gengo.Filepath(g.Pkg, g.Config.Output), g.Buf.Bytes(), nil)
}
