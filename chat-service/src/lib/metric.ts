import { DiagConsoleLogger, DiagLogLevel, diag } from "@opentelemetry/api";
import {
  MeterProvider,
  MeterProviderOptions,
  PeriodicExportingMetricReader,
} from "@opentelemetry/sdk-metrics";
import { OTLPMetricExporter } from "@opentelemetry/exporter-metrics-otlp-http";
import { Resource } from "@opentelemetry/resources";
import { SemanticResourceAttributes } from "@opentelemetry/semantic-conventions";
//import { OTLPMetricExporter } from "@opentelemetry/exporter-metrics-otlp-grpc";

//diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.DEBUG);

const collectorOptions = {
  // url is optional and can be omitted - default is grpc://localhost:4317
  // url: 'grpc://<IP of signoz backend>:4317',
  //http://otel-collector:4318/v1/metrics"
  url: "http://otel-collector:4318/v1/metrics",
};
const metricExporter = new OTLPMetricExporter(collectorOptions);

// Register the exporter
const meterProvider = new MeterProvider({
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: "chat-service",
  }),
});

meterProvider.addMetricReader(
  new PeriodicExportingMetricReader({
    exporter: metricExporter,
    exportIntervalMillis: 1000,
  })
);

//opentelemetry.metrics.setGlobalMeterProvider(meterProvider);

const meter = meterProvider.getMeter("chat_service_meter");

const requestCount = meter.createCounter("requests_count", {
  description: "Count all incoming requests",
});

requestCount.add(10, { 'key': 'value' });

//const boundInstruments = new Map();

export const countAllRequests = () => {
  return (req: any, res: any, next: any) => {
    const attributes = { pid: process.pid, environment: "poc" };
    requestCount.add(1, attributes);
    next();
  };
};
