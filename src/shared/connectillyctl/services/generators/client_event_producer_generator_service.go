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

type ClientEventProducerGeneratorService interface {
	Generate(
		packageName string,
		eventType string,
		outputPath string) error
}

type clientEventProducerGeneratorService struct {
	logger   *zap.SugaredLogger
	osHelper os.OsHelper
}

func NewClientEventProducerGeneratorService(
	logger *zap.SugaredLogger,
	osHelper os.OsHelper) (ClientEventProducerGeneratorService, error) {
	return &clientEventProducerGeneratorService{
		logger:   logger,
		osHelper: osHelper,
	}, nil
}

func (cehgs *clientEventProducerGeneratorService) Generate(
	packageName string,
	eventType string,
	outputPath string) error {
	type data struct {
		PackageName string
		EventType   string
	}

	tmpl, err := template.New("").Parse(clienteventtemplates.GetProducer())
	if err != nil {
		return err
	}

	var processed bytes.Buffer
	if err = tmpl.ExecuteTemplate(&processed, "", &data{PackageName: packageName, EventType: eventType}); err != nil {
		return err
	}

	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		return fmt.Errorf("could not format processed template: %v", err)
	}

	err = cehgs.osHelper.CreateDir(outputPath)
	if err != nil {
		return err
	}

	return cehgs.osHelper.CreateBinaryFile(
		path.Join(outputPath, "producer_eventgen.go"),
		formatted)
}
