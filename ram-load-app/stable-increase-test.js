import http from "k6/http";
import { sleep } from "k6";
import { randomString } from "https://jslib.k6.io/k6-utils/1.2.0/index.js";

export let options = {
  stages: [{ target: 1000, duration: "2h" }],
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

const BASE_URL = "http://localhost:4000";

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
