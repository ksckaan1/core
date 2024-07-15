// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parse does the parsing stage after lexing
package parser

//go:generate core generate

import (
	"fmt"
	"io"

	"cogentcore.org/core/base/indent"
	"cogentcore.org/core/parse/lexer"
	"cogentcore.org/core/parse/syms"
	"cogentcore.org/core/tree"
)

// AST is a node in the abstract syntax tree generated by the parsing step
// the name of the node (from tree.NodeBase) is the type of the element
// (e.g., expr, stmt, etc)
// These nodes are generated by the parser.Rule's by matching tokens
type AST struct {
	tree.NodeBase

	// region in source lexical tokens corresponding to this Ast node -- Ch = index in lex lines
	TokReg lexer.Reg `set:"-"`

	// region in source file corresponding to this Ast node
	SrcReg lexer.Reg `set:"-"`

	// source code corresponding to this Ast node
	Src string `set:"-"`

	// stack of symbols created for this node
	Syms syms.SymStack `set:"-"`
}

func (ast *AST) Destroy() {
	ast.Syms.ClearAst()
	ast.Syms = nil
	ast.NodeBase.Destroy()
}

// ChildAst returns the Child at given index as an Ast.
// Will panic if index is invalid -- use Try if unsure.
func (ast *AST) ChildAst(idx int) *AST {
	return ast.Child(idx).(*AST)
}

// ParentAst returns the Parent as an Ast.
func (ast *AST) ParentAst() *AST {
	if ast.Parent == nil {
		return nil
	}
	pn := ast.Parent.AsTree().This
	if pn == nil {
		return nil
	}
	return pn.(*AST)
}

// NextAst returns the next node in the Ast tree, or nil if none
func (ast *AST) NextAst() *AST {
	nxti := tree.Next(ast)
	if nxti == nil {
		return nil
	}
	return nxti.(*AST)
}

// NextSiblingAst returns the next sibling node in the Ast tree, or nil if none
func (ast *AST) NextSiblingAst() *AST {
	nxti := tree.NextSibling(ast)
	if nxti == nil {
		return nil
	}
	return nxti.(*AST)
}

// PrevAst returns the previous node in the Ast tree, or nil if none
func (ast *AST) PrevAst() *AST {
	nxti := tree.Previous(ast)
	if nxti == nil {
		return nil
	}
	return nxti.(*AST)
}

// SetTokReg sets the token region for this rule to given region
func (ast *AST) SetTokReg(reg lexer.Reg, src *lexer.File) {
	ast.TokReg = reg
	ast.SrcReg = src.TokenSrcReg(ast.TokReg)
	ast.Src = src.RegSrc(ast.SrcReg)
}

// SetTokRegEnd updates the ending token region to given position --
// token regions are typically over-extended and get narrowed as tokens actually match
func (ast *AST) SetTokRegEnd(pos lexer.Pos, src *lexer.File) {
	ast.TokReg.Ed = pos
	ast.SrcReg = src.TokenSrcReg(ast.TokReg)
	ast.Src = src.RegSrc(ast.SrcReg)
}

// WriteTree writes the AST tree data to the writer -- not attempting to re-render
// source code -- just for debugging etc
func (ast *AST) WriteTree(out io.Writer, depth int) {
	ind := indent.Tabs(depth)
	fmt.Fprintf(out, "%v%v: %v\n", ind, ast.Name, ast.Src)
	for _, k := range ast.Children {
		ai := k.(*AST)
		ai.WriteTree(out, depth+1)
	}
}
