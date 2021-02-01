package main

import (
	"os"

	"github.com/pborman/getopt"

	"github.com/roflcopter4/x4c-go/myxml"
	"github.com/roflcopter4/x4c-go/translation"
	"github.com/roflcopter4/x4c-go/util"
)

var opt struct {
	infname string
	outfile struct {
		name string
		fp   *os.File
	}
	help     bool
	dump     bool
	validate bool

	operation string
}

func init() {
	getopt.BoolVarLong(&opt.help, "help", 'h', "Display help")
	getopt.BoolVarLong(&opt.dump, "dump", 'D', "Dump the tree")
	getopt.BoolVarLong(&opt.validate, "validate", 'v', "Validate input with schema")
	getopt.StringVarLong(&opt.infname, "file", 'f', "Input filename")
	getopt.StringVarLong(&opt.outfile.name, "out", 'o', "Output filename")

	getopt.EnumVarLong(&opt.operation, "operation", 'x', []string{"u", "t", "q", "Q", "s", "S"}, "Operation")
}

func main() {
	args := handle_opts()

	switch opt.operation {
	case "u":
		translation.UnTranslate(opt.outfile.fp, opt.infname)
	case "t":
		do_translate()
	case "q":
		translation.TestLexer(args[0], true)
	case "Q":
		translation.TestLexer(args[0], false)
	case "s":
		translation.TestScriptLexer(args[0], true)
	case "S":
		translation.TestScriptLexer(args[0], false)
	}
}

func handle_opts() []string {
	getopt.Parse()
	args := getopt.Args()

	if opt.help {
		getopt.Usage()
		os.Exit(0)
	}

	if opt.infname == "" {
		if getopt.NArgs() == 0 {
			getopt.Usage()
			os.Exit(1)
		}
		opt.infname = args[0]
	}

	if util.StrEqAny(opt.outfile.name, "", "-") {
		opt.outfile.fp = os.Stdout
	} else {
		fp, err := os.Create(opt.outfile.name)
		if err != nil {
			util.PanicE(err)
		}
		opt.outfile.fp = fp
	}

	return args
}

func do_translate() {
	doc, err := myxml.New_Document(opt.infname)
	if err != nil {
		util.DieE(2, err)
		// panic(err)
	}
	defer doc.Free()

	if opt.validate {
		err = doc.GetSchema()
		if err != nil {
			util.DieE(2, err)
			// panic(err)
		}

		err = doc.ValidateSchema()
		if err != nil {
			util.DieE(2, err)
			// panic(err)
		}
	}

	// translate.TestReader(opt.outfile.fp, doc)
	// translation.TestTranslate(opt.outfile.fp, doc)
	translation.Translate(opt.outfile.fp, doc)
}
