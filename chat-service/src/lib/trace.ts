const opentelemetry = require("@opentelemetry/sdk-node");
const {
  getNodeAutoInstrumentations,
} = require("@opentelemetry/auto-instrumentations-node");
const { diag, DiagConsoleLogger, DiagLogLevel } = require("@opentelemetry/api");
const {
  OTLPTraceExporter,
} = require("@opentelemetry/exporter-trace-otlp-http");
const { Resource } = require("@opentelemetry/resources");
const {
  SemanticResourceAttributes,
} = require("@opentelemetry/semantic-conventions");

// For troubleshooting, set the log level to DiagLogLevel.DEBUG
diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.INFO);

//http://otel-collector:4318/v1/traces
const traceExporter = new OTLPTraceExporter({
  url: "http://otel-collector:4318/v1/traces",
});

const sdk = new opentelemetry.NodeSDK({
  traceExporter: traceExporter,
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: "chat-service",
  }),
  instrumentations: [
    getNodeAutoInstrumentations({
      "@opentelemetry/instrumentation-http": {
        applyCustomAttributesOnSpan: (span: any) => {
          span.setAttribute("chat-service", "arul-test");
        },
      },
    }),
  ],
});

sdk
  .start()
  .then(() => console.log("Tracing initialized"))
  .catch((error: any) => console.log("Error initializing tracing", error));

// gracefully shut down the SDK on process exit
process.on("SIGTERM", () => {
  sdk
    .shutdown()
    .then(() => console.log("Tracing terminated"))
    .catch((error: any) => console.log("Error terminating tracing", error))
    .finally(() => process.exit(0));
});
