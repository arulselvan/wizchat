import { Router } from "express";
import { logger } from './lib/logger'
const routes = Router();

routes.get("/chat/initiate", (req, res) => {
  logger.info("inside chat initiate")
  return res.json({ message:'Chat Request Accepted, will connect once Agent available' });
});

export default routes;