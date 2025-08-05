import http from "k6/http";
import { sleep, check } from "k6";

const BASE_URL = "http://localhost:4000";

export let options = {
  stages: [
    { duration: "30s", target: 30 },

    // First spike: 400 VUs
    { duration: "2m", target: 400 }, // Spike to 400 within 2 minutes
    { duration: "5m", target: 400 }, // Hold at 400 for 5 minutes
    { duration: "2m", target: 30 }, // Reduce back to 30 within 2 minutes
    { duration: "5m", target: 30 }, // Wait for 5 minutes

    // Second spike: 800 VUs
    { duration: "2m", target: 800 },
    { duration: "5m", target: 800 },
    { duration: "2m", target: 30 },
    { duration: "5m", target: 30 },

    // Final spike: 1400 VUs
    { duration: "2m", target: 1400 },
    { duration: "5m", target: 1400 },
    { duration: "2m", target: 30 },
    { duration: "2m", target: 0 },
  ],

  thresholds: {
    http_req_failed: ["rate<0.01"],
    http_req_duration: ["p(95)<800"],
    checks: ["rate>0.99"],
  },
};

export default function () {
  const params = {
    timeout: "1s",
  };

  const res = http.get(`${BASE_URL}/uuid`, params);

  check(res, {
    "✅ status is 200": (r) => r.status === 200,
    "✅ response body is not empty": (r) => r.body && r.body.length > 0,
    "✅ content-type header is correct": (r) =>
      r.headers["Content-Type"] === "application/json",
    "✅ response time under 1s": (r) => r.timings.duration < 1000,
  });

  sleep(0.1);
}
