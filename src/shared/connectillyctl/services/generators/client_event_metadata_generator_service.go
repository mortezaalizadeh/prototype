package generators

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"path"

	clienteventtemplates "github.com/Connectilly/connectilly/src/shared/connectillyctl/templates/client/event"
	"github.com/Connectilly/connectilly/src/shared/enterprise/os"
	"go.uber.org/zap"
)

type ClientEventMetadataGeneratorService interface {
	Generate(
		packageName string,
		topicName string,
		outputPath string) error
}

type clientEventMetadataGeneratorService struct {
	logger   *zap.SugaredLogger
	osHelper os.OsHelper
}

func NewClientEventMetadataGeneratorService(
	logger *zap.SugaredLogger,
	osHelper os.OsHelper) (ClientEventMetadataGeneratorService, error) {
	return &clientEventMetadataGeneratorService{
		logger:   logger,
		osHelper: osHelper,
	}, nil
}

func (cemgs *clientEventMetadataGeneratorService) Generate(
	packageName string,
	topicName string,
	outputPath string) error {
	type data struct {
		PackageName string
		TopicName   string
	}

	tmpl, err := template.New("").Parse(clienteventtemplates.GetMetadata())
	if err != nil {
		return err
	}

	var processed bytes.Buffer
	if err = tmpl.ExecuteTemplate(&processed, "", &data{PackageName: packageName, TopicName: topicName}); err != nil {
		return err
	}

	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		return fmt.Errorf("could not format processed template: %v", err)
	}

	err = cemgs.osHelper.CreateDir(outputPath)
	if err != nil {
		return err
	}

	return cemgs.osHelper.CreateBinaryFile(
		path.Join(outputPath, "metadata_eventgen.go"),
		formatted)
}
