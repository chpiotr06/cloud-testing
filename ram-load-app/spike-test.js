import http from "k6/http";
import { sleep, check } from "k6";
import { randomString } from "https://jslib.k6.io/k6-utils/1.2.0/index.js";

const BASE_URL = "http://localhost:4000";

export let options = {
  stages: [
    { duration: "30s", target: 20 },
    { duration: "2m", target: 20 },

    // First spike: 200 VUs
    { duration: "2m", target: 200 }, // Spike to 200 within 2 minutes
    { duration: "5m", target: 200 }, // Hold at 200 for 5 minutes
    { duration: "2m", target: 20 }, // Reduce back to 20 within 2 minutes
    { duration: "7m", target: 20 }, // Wait for 7 minutes

    // Second spike: 400 VUs
    { duration: "2m", target: 400 },
    { duration: "5m", target: 400 },
    { duration: "2m", target: 20 },
    { duration: "7m", target: 20 },

    { duration: "2m", target: 800 },
    { duration: "5m", target: 800 },
    { duration: "2m", target: 20 },
    { duration: "7m", target: 20 },

    // Final spike: 1400 VUs
    { duration: "2m", target: 1000 },
    { duration: "5m", target: 1000 },
    { duration: "2m", target: 20 },
    { duration: "7m", target: 0 },
  ],

  systemTags: [
    "proto",
    "subproto",
    "status",
    "method",
    "name",
    "group",
    "check",
    "error",
    "error_code",
    "tls_version",
    "scenario",
    "service",
    "expected_response",
  ],
};

function generateRandomData() {
  const username = `user_${randomString(8)}`;
  const email = `${randomString(6)}@example.com`;
  return { username, email };
}

const timeout = "1s";

export default function () {
  const sessionData = generateRandomData();

  const createSessionPayload = JSON.stringify({
    username: sessionData.username,
    email: sessionData.email,
  });

  const res = http.post(`${BASE_URL}/session`, createSessionPayload, {
    headers: {
      "Content-Type": "application/json",
    },
    timeout,
  });

  check(res, {
    "✅ status is 200": (r) => r.status === 200,
    "✅ response body is not empty": (r) => r.body && r.body.length > 0,
    "✅ response time under 1s": (r) => r.timings.duration < 1000,
  });

  sleep(1);
}
