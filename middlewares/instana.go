package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// InstanaTracerMiddleware responsable to spaw tracer to all requests
func InstanaTracerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		// Collect request parameters to add them to the entry HTTP span. We also need to make
		// sure that a proper span kind is set for the entry span, so that Instana could combine
		// it and its children into a call.
		opts := []opentracing.StartSpanOption{
			ext.SpanKindRPCServer,
			opentracing.Tags{
				"http.method":  c.Request.Method,
				"http.headers": c.Request.Header,
				"http.host":    c.Request.Host,
				"http.uri":     c.Request.RequestURI,
				"http.query":   c.Request.URL.Query(),
			},
		}

		// If there is any header information comming from other span, this one will be
		//  a child span
		wireContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		if err != nil {
			opts = append(opts, ext.RPCServerOption(wireContext))
		}

		// Start the entry span adding collected tags and optional parent. The span name here
		// matters, as it allows Instana backend to classify the call as an HTTP one.
		span := opentracing.GlobalTracer().StartSpan("g.http", opts...)
		defer span.Finish()
		c.Next()
		// after request
	}
}
