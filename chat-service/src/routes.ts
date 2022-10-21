import { Router } from "express";
import { logger } from "./lib/logger";
import axios from "axios";

const routes = Router();

routes.get("/chat/initiate", async (req, res) => {
  logger.info("inside chat initiate");

  logger.info("routing chat request");
  const response = await axios.post("http://router-service:8080/route", {
    reqType: "chat",
    userId: "101",
    businessLine: "sales",
  });

  logger.info("router response", response.data);

  return res.json({
    message: "Chat Request Accepted, will connect once Agent available",
  });
});

export default routes;
