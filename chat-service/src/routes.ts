import { Router } from "express";

const routes = Router();

routes.get("/chat/initate", (req, res) => {
  return res.json({ message:'Awaiting for available Agent' });
});

export default routes;