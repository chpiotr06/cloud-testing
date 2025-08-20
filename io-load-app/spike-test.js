import http from "k6/http";
import { sleep, check } from "k6";

const BASE_URL = "http://localhost:4000";

export let options = {
  stages: [
    { duration: "30s", target: 10 },
    { duration: "3m", target: 10 },
    { duration: "2m", target: 100 }, // Spike to 100 within 2 minutes
    { duration: "5m", target: 100 }, // Hold at 100 for 5 minutes
    { duration: "2m", target: 10 }, // Reduce back to 10 within 2 minutes
    { duration: "5m", target: 10 }, // Wait for 5 minutes

    // Second spike: 400 VUs
    { duration: "2m", target: 300 },
    { duration: "5m", target: 300 },
    { duration: "2m", target: 10 },
    { duration: "5m", target: 10 },

    // Third spike: 400 VUs
    { duration: "2m", target: 600 },
    { duration: "5m", target: 600 },
    { duration: "2m", target: 10 },
    { duration: "5m", target: 10 },

    // Final spike: 800 VUs
    { duration: "2m", target: 900 },
    { duration: "5m", target: 900 },
    { duration: "2m", target: 10 },
    { duration: "2m", target: 0 },
  ],
  thresholds: {
    http_req_failed: ["rate<0.01"],
    http_req_duration: ["p(95)<800"],
    checks: ["rate>0.99"],
  },
};

// Create 4KB random payload
function makePayload(size = 4 * 1024) {
  let arr = new Uint8Array(size);
  crypto.getRandomValues(arr);
  return arr.buffer;
}

function makeFormData() {
  return {
    file: http.file(
      makePayload(),
      `test-${__VU}-${Date.now()}.bin`,
      "application/octet-stream"
    ),
  };
}

export default function () {
  const res = http.post(`${BASE_URL}/upload`, makeFormData(), {
    timeout: "2s",
  });

  check(res, {
    "✅ status is 201": (r) => r.status === 201,
  });

  sleep(0.1);
}
