import http from "k6/http";
import { sleep, check } from "k6";

const BASE_URL = "http://localhost:3000";

export let options = {
  stages: [{ duration: "1h", target: 800 }],
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
  const res = http.post(`${BASE_URL}/upload`, makeFormData());

  check(res, {
    "✅ status is 201": (r) => r.status === 201,
    "✅ upload succeeded": (r) => r.body && r.body.includes("Saved"),
  });

  sleep(0.1);
}
