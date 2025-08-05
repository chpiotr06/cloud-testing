import http from "k6/http";
import { sleep, check } from "k6";

export let options = {
  stages: [{ duration: "1h", target: 1300 }],
  thresholds: {
    http_req_failed: ["rate<0.01"],
    http_req_duration: ["p(95)<800"],
    checks: ["rate>0.99"],
  },
};

const BASE_URL = "http://localhost:4000";

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
  });

  sleep(0.1);
}
