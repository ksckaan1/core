// Code generated by "core generate -add-types -add-funcs"; DO NOT EDIT.

package main

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "main.Config", IDName: "config", Doc: "Config is the configuration information for the cosh cli.", Directives: []types.Directive{{Tool: "go", Directive: "generate", Args: []string{"core", "generate", "-add-types", "-add-funcs"}}}, Fields: []types.Field{{Name: "Input", Doc: "The input file to run/compile.\nIf this is provided as the first argument,\nthen the program will exit after running,\nunless the Interactive mode is flagged."}, {Name: "Output", Doc: "the Go file to output the transpiled Input file to,\nas an optional second argument in build mode.\nIt defaults to the input file with .cosh changed to .go."}, {Name: "Expr", Doc: "an optional expression to evaluate, which can be used\nin addition to the Input file to run, to execute commands\ndefined within that file for example, or as a command to run\nprior to starting interactive mode if no Input is specified."}, {Name: "Interactive", Doc: "runs the interactive command line after processing an Input file.\nInteractive mode is the default for all cases except when\nan Input file is specified, and is not available\nif an Output file is specified for transpiling."}}})

var _ = types.AddFunc(&types.Func{Name: "main.Run", Doc: "Run runs the specified cosh file. If no file is specified,\nit runs an interactive shell that allows the user to input cosh.", Directives: []types.Directive{{Tool: "cli", Directive: "cmd", Args: []string{"-root"}}}, Args: []string{"c"}, Returns: []string{"error"}})

var _ = types.AddFunc(&types.Func{Name: "main.Interactive", Doc: "Interactive runs an interactive shell that allows the user to input cosh.", Args: []string{"c"}, Returns: []string{"error"}})

var _ = types.AddFunc(&types.Func{Name: "main.Build", Doc: "Build builds the specified input cosh file to the specified output Go file.", Args: []string{"c"}, Returns: []string{"error"}})
