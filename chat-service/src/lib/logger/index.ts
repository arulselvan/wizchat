const winston = require('winston');

export const logger = winston.createLogger({
  level: 'info',
  format: winston.format.json(),
  defaultMeta: { service: 'chat-service' },
  transports: [
    new winston.transports.Console()
  ],
});
