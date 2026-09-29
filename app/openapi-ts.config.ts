import { defineConfig } from "@hey-api/openapi-ts";

export default defineConfig({
  input: "../backend/internal/docs/swagger.yaml",
  output: "./client",
  plugins: [
    {
      name: "@hey-api/client-axios",
      baseUrl: "/api/v1" ,
    },
    {
      name: "@hey-api/sdk",
      operations: {
        strategy: "byTags",
        containerName: "{{name}}Service",
        nesting: "operationId",
        methodName: (name) => name.charAt(0).toLowerCase() + name.slice(1),
      },
    },
    {
      name: "@hey-api/schemas",
      type: "json",
    },
    "@tanstack/react-query",
    {
      name: "zod",
    },
  ],
});
