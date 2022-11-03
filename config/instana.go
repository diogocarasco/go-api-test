package configs

import (
	"os"

	instana "github.com/instana/go-sensor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func init() {

	if os.Getenv("ENVIRONMENT") == "prod" {
		tracer := instana.NewTracerWithOptions(instana.DefaultOptions())
		opentracing.SetGlobalTracer(tracer)
		instana.SetLogger(logrus.StandardLogger())
	}
}
