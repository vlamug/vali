package main

import (
	"log"
	"os"

	"github.com/vlamug/vali/config"
	"github.com/vlamug/vali/util"
	"github.com/vlamug/vali/validation"
	"github.com/vlamug/vali/validation/data"
	"github.com/vlamug/vali/validation/validator"

	"github.com/alecthomas/kingpin"
	"github.com/smallfish/simpleyaml"
	"go.uber.org/zap"
)

const defaultConfigPath = "etc/config.yaml"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("could not sync logger: %s\n", err)
	}
	defer func() {
		_ = logger.Sync()
	}()

	var (
		app        = kingpin.New("vali", "Validates any yaml file.")
		configPath = app.Flag("config.path", "Path to config file").Default(defaultConfigPath).String()
	)
	_, err = app.Parse(os.Args[1:])
	if err != nil {
		logger.Fatal("could not parse cli args", zap.Error(err))
	}

	cfg, err := config.MakeConfigFromFile(*configPath)
	if err != nil {
		logger.Fatal("could not make file from config file", zap.Error(err))
	}

	bytes, err := util.ReadBytesFromStdin()
	if err != nil {
		logger.Fatal("could not read data from stdin", zap.Error(err))
	}

	node, err := simpleyaml.NewYaml(bytes)
	if err != nil {
		logger.Fatal("could not parse yaml file", zap.Error(err))
	}

	validators := []validator.ValidateFunc{
		validator.Match,
		validator.AnyOf,
		validator.MatchRe,
		validator.IsNumber,
		validator.Required,
		validator.NotEmpty,
		validator.Absent,
		validator.IsMap,
		validator.IsArray,
	}

	runner := validation.NewRunner(cfg.Rules, validators, data.NewWriterReport(os.Stdout))
	if report, err := runner.Run(node); err != nil {
		logger.Fatal("could not run validator", zap.Error(err))
	} else if err := report.Publish(); err != nil {
		logger.Fatal("could not publish validation report", zap.Error(err))
	}
}
