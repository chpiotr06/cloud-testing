import http from "k6/http";
import { sleep, check } from "k6";

const BASE_URL = "http://localhost:4000";

export let options = {
  stages: [{ duration: "1h", target: 1300 }],
  thresholds: {
    http_req_failed: ["rate<0.01"],
    http_req_duration: ["p(95)<800"],
    checks: ["rate>0.99"],
  },
};

// File map corresponding to your Go map
const fileMap = ["5kb", "10kb", "25kb", "50kb", "100kb"];

// Get array of file keys for random selection
const fileKeys = Object.keys(fileMap);

export default function () {
  const params = {
    timeout: "1s",
  };

  // Randomly select a file
  const randomFileKey = fileKeys[Math.floor(Math.random() * fileKeys.length)];
  const fileName = fileMap[randomFileKey];

  const res = http.get(`${BASE_URL}/stream/${fileName}`, params);

  check(res, {
    "✅ status is 200": (r) => r.status === 200,
    "✅ response body is not empty": (r) => r.body && r.body.length > 0,
    "✅ content-type header is correct": (r) =>
      r.headers["Content-Type"] === "application/octet-stream",
    "✅ file streamed successfully": (r) => r.status === 200,
  });

  sleep(0.1);
}
