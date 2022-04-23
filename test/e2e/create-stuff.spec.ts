import { test, expect } from "@playwright/test";

test.describe("create-stuff", () => {
  test("should create stuff", async ({ request }) => {
    let response = await request.post("/v1/stuff", {
      data: {},
    });

    expect(response.status()).toBe(201);

    let { data } = await response.json();

    expect(data.id).toBeTruthy();
  });
});
