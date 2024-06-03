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
1. Add a `package_path` property of type `string` to the `Repo` object.
1. Add a `functions_region` property of type `string` to the `Site` object.
1. Add a `cdp_enabled_contexts` property of type `array` of `string`s to the `Site` object.
1. Add a `hud_enabled` property of type `boolean` to the `Site` object.
1. Duplicate the `Site` object into `PartialSite` and remove the `required` properties.
1. Change `updateSite` operation to use the `PartialSite` object as the request body schema (NOTE: not the response body schema).
1. Change the type of `LogDrain.id` to `string`.
1. Add the various `log_drains` paths from `bitballoon-openapi`'s `openapi.json` file.
1. Remove the required properties from the `LogDrainServiceConfig` object.
1. Add properties to the `LogDrainServiceConfig` object, by copying it from an earlier version of the `openapi.json`.
1. Change the request body of the `Log Drains-update` operation to use the `LogDrain` object (copy from `Log Drains-create`).