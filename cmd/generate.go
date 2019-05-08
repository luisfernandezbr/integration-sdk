package cmd

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/log"
	"github.com/pinpt/integration-sdk/gen"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate <indir> <outdir>",
	Aliases: []string{"gen", "g"},
	Short:   "Generate code",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		logger := newCommandLogger(cmd)
		defer logger.Close()
		started := time.Now()
		log.Info(logger, "schema generation started")
		indir, outdir := args[0], args[1]
		absdir, err := filepath.Abs(indir)
		if err != nil {
			log.Fatal(logger, "error resolving indir path", "path", indir, "err", err)
		}
		os.MkdirAll(outdir, 0777)
		files, err := fileutil.FindFiles(absdir, regexp.MustCompile("\\.y[a]?ml$"))
		if err != nil {
			log.Fatal(logger, "error finding yaml files in path", "path", absdir, "err", err)
		}
		for _, file := range files {
			of, err := os.Open(file)
			if err != nil {
				log.Fatal(logger, "error opening "+file, "err", err)
			}
			defer of.Close()
			var schema gen.Schema
			if err := schema.DecodeYAML(of); err != nil {
				log.Fatal(logger, "error parsing "+file, "err", err)
			}
			log.Info(logger, "generating schema", "schema", schema.Name)
			fn := filepath.Join(outdir, schema.Name, "v"+strconv.Itoa(schema.Version), schema.Name, "init.go")
			if err := os.MkdirAll(filepath.Dir(fn), 0777); err != nil {
				log.Fatal(logger, "error creaing dir", "fn", fn, "err", err)
			}
			wf, err := os.Create(fn)
			if err != nil {
				log.Fatal(logger, "error creating "+fn, "err", err)
			}
			defer wf.Close()
			if err := gen.Generate(schema, wf); err != nil {
				log.Fatal(logger, "error generating "+fn, "err", err)
			}
			for name := range schema.Models {
				startedgen := time.Now()
				log.Info(logger, "generating model", "model", name)
				fn := schema.GetSchemaFilePath(outdir, name)
				if err := os.MkdirAll(filepath.Dir(fn), 0777); err != nil {
					log.Fatal(logger, "error creaing dir", "fn", fn, "err", err)
				}
				wf, err := os.Create(fn)
				if err != nil {
					log.Fatal(logger, "error creating "+fn, "err", err)
				}
				defer wf.Close()
				if err := gen.GenerateModel(schema, name, wf); err != nil {
					log.Fatal(logger, "error generating "+fn, "err", err)
				}
				log.Info(logger, "done generating", "schema", name, "duration", time.Since(startedgen))
			}
		}
		log.Info(logger, "schema generation ended", "duration", time.Since(started))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
