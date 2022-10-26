import express from "express";
import { countAllRequests } from "./lib/metric";
import routes from "./routes";

class App {
  public server;

  constructor() {
    this.server = express();
    this.middlewares();
    this.routes();
  }

  middlewares() {
    this.server.use(express.json());
    this.server.use(countAllRequests());
  }

  routes() {
    this.server.use(routes);
  }
}

export default new App().server;
