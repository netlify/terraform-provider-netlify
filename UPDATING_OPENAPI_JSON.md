# Updating swagger.yml

**This is a temporary measure. Feedback will be shared to fix upstream.**

This project uses a modified `openapi.json`. Please maintain these instructions and follow them carefully.

1. Take the latest `openapi-external.json` from [netlify/bitballoon-openapi](https://github.dev/netlify/bitballoon-openapi/blob/main/openapi-external.json).
1. Remove the billing_details property from the `Account` object (also from the `required` array).
1. Fix the type of `Repo.base_rel_dir` to `boolean`.
1. Remove all required properties from the `Repo` object (manual builds).
1. Remove the `domain` property from the `required` array of the `DnsZone` object.
1. Remove the `values`, `scopes` and `is_secret` parameters from the `updateEnvVar` operation.
1. Add a request body schema to the `updateEnvVar` operation, by copying it from an earlier version of the `openapi.json`.