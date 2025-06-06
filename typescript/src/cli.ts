import { program } from "commander";

import client from "./api";

program
	.name("typescript-measure-cli")
	.description("CLI for Analyzing Scope3 Emission Data")
	.version("0.0.1");

program.command("probe").action(async () => {
	const response = await client.measure(["yahoo.com"], "2025-05-01");
	const toLog = `
RequestID: ${response.requestId}
Total Emissions: ${response.totalEmissions}
Breakdown:\n
  - Ad Selection: ${response.totalEmissionsBreakdown.totals.adSelection}
  - Creative Delivery: ${response.totalEmissionsBreakdown.totals.creativeDelivery}
  - Media Distribution: ${response.totalEmissionsBreakdown.totals.mediaDistribution}
	`;
	console.log(toLog);
});

program.parse();
