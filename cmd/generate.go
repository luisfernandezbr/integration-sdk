package cmd

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/integration-sdk/gen"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate <indir> <outdir>",
	Aliases: []string{"gen", "g"},
	Short:   "Generate code",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		indir, outdir := args[0], args[1]
		absdir, err := filepath.Abs(indir)
		if err != nil {
			log.Fatal(err)
		}
		os.MkdirAll(outdir, 0777)
		files, err := fileutil.FindFiles(absdir, regexp.MustCompile("\\.y[a]?ml$"))
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			of, err := os.Open(file)
			if err != nil {
				log.Fatalf("error opening %s. %v", file, err)
			}
			defer of.Close()
			var schema gen.Schema
			if err := schema.DecodeYAML(of); err != nil {
				log.Fatalf("error parsing %s. %v", file, err)
			}
			for name := range schema.Models {
				fn := schema.GetSchemaFilePath(outdir, name)
				os.MkdirAll(filepath.Dir(fn), 0777)
				wf, err := os.Create(fn)
				if err != nil {
					log.Fatalf("error creating %s. %v", fn, err)
				}
				defer wf.Close()
				if err := gen.Generate(schema, name, wf); err != nil {
					log.Fatalf("error generating %s. %v", fn, err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
