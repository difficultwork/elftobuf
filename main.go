package main

import (
	"fmt"
	"os"
	"path/filepath"

	"elftobuf/internal/lang"

	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		Lang         string `short:"l" long:"language" description:"specify the buffer which language to be generated. support c/go"`
		ElfFile      string `short:"e" long:"elf-file" description:"specify elf file to be converted."`
		TargetFile   string `short:"t" long:"target-file" description:"specify target file to be generated."`
		VariableName string `short:"v" long:"variable-name" description:"specify the buffer name to be generated."`
		PackageName  string `short:"p" long:"package-name" description:"specify the package name to which the target file belongs."`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
			return
		}
		fmt.Println("parse params failed: ", err)
		os.Exit(1)
	}

	if opts.ElfFile, err = filepath.Abs(opts.ElfFile); err != nil {
		fmt.Println("invalid elf-file: ", opts.ElfFile)
		os.Exit(1)
	}

	if opts.TargetFile, err = filepath.Abs(opts.TargetFile); err != nil {
		fmt.Println("invalid target-file: ", opts.TargetFile)
		os.Exit(1)
	}

	if len(opts.VariableName) == 0 {
		fmt.Println("empty variable-name: ", opts.VariableName)
		os.Exit(1)
	}

	switch opts.Lang {
	case "c":
		if err = lang.ElfToC(opts.ElfFile, opts.TargetFile, opts.VariableName); err != nil {
			fmt.Println("elf-to-c failed: ", err)
			os.Exit(1)
		}
	case "go":
		if len(opts.PackageName) == 0 {
			fmt.Println("empty package-name: ", opts.PackageName)
			os.Exit(1)
		}
		if err = lang.ElfToGo(opts.ElfFile, opts.TargetFile, opts.PackageName, opts.VariableName); err != nil {
			fmt.Println("elf-to-go failed: ", err)
			os.Exit(1)
		}
	default:
		fmt.Println("unsupported language: ", opts.Lang)
		os.Exit(1)
	}
}
