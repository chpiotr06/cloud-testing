import { randomUUID } from "node:crypto";
import * as http from "node:http";
import { genSalt, hash } from "bcrypt";
import config from "./config.js";

const server = http.createServer(
  { keepAliveTimeout: 60000 },
  async (req, res) => {
    if (req.url == "/healthz") {
      res.writeHead(200, { "Content-Type": "text/plain" });
      res.end("OK");
      return;
    }

    if (req.method == "GET" && req.url == "/api/test/1") {
      const uuid = randomUUID();
      const salt = await genSalt(13);
      const hashedValue = await hash(uuid, salt);
      const testPayload = { uuid, hash: hashedValue };

      res.writeHead(200, { "Content-Type": "application/json" });
      res.end(JSON.stringify(testPayload));
      return;
    }
  }
);

console.log(`Node is listening on http://0.0.0.0:${config.appPort} ...`);

server.listen(config.appPort);
