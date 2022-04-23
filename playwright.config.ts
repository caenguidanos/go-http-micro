import { PlaywrightTestConfig } from "@playwright/test";

const config: PlaywrightTestConfig = {
  use: {
    baseURL: `http://localhost:8080`,
    browserName: "chromium",
  },
  testDir: "test/e2e",
  testMatch: "*.spec.ts",
  timeout: 30000,
  retries: 2,
};

export default config;
