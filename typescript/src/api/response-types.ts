interface AdFormats {
	generic: number;
	metric: string;
	unknown: number;
	vendorSpecific: number;
}

interface Channels {
	deprecated: number;
	metric: string;
	modeled: number;
	unknown: number;
}

interface BaseMetric {
	metric: string;
	modeled: number;
	unknown: number;
}

interface TotalMetric {
	metric: string;
	modeled: number;
	skipped: number;
}

interface Coverage {
	adFormats: AdFormats;
	channels: Channels;
	mediaOwners: BaseMetric;
	properties: BaseMetric;
	sellers: BaseMetric;
	totalImpressions: TotalMetric;
	totalRows: TotalMetric;
}

interface Policy {
	compliant: number;
	noncompliant: number;
	policy: string;
	policyOwner: string;
}

interface RowPolicy {
	compliant: boolean;
	policy: string;
	policyOwner: string;
}

interface EmissionsTotals {
	adSelection: number;
	creativeDelivery: number;
	mediaDistribution: number;
}

interface EmissionsBreakdown {
	framework: string;
	totals: EmissionsTotals;
}

interface AdFormatCoverage {
	name: string;
	value: string;
	verified: boolean;
}

interface SimpleValue {
	value: string;
}

interface Impressions {
	modeled: number;
	processed: number;
	skipped: number;
}

interface SupplyGraphDepth {
	averageDepth: number;
	maxDepth: number;
	minDepth: number;
}

interface SupplyGraph {
	logical: SupplyGraphDepth;
	technical: SupplyGraphDepth;
	totalCount: number;
}

interface RowCoverage {
	adFormat: AdFormatCoverage;
	channel: SimpleValue;
	compensationProvider: SimpleValue;
	impressions: Impressions;
	property: SimpleValue;
	supplyGraph: SupplyGraph;
}

interface EmissionsComponent {
	emissions: number;
}

interface CompensationBreakdown {
	emissions: number;
	provider: string;
}

interface CompensatedBreakdownStruct {
	compensation: CompensationBreakdown;
}

interface Compensated {
	breakdown: CompensatedBreakdownStruct;
	total: number;
}

interface SimpleBreakdown {
	adPlatform: EmissionsComponent;
	dataTransfer: EmissionsComponent;
}

interface MediaDistributionBreakdown {
	corporate: EmissionsComponent;
	dataTransfer: EmissionsComponent;
}

interface ComponentWithBreakdown {
	breakdown: SimpleBreakdown;
	total: number;
}

interface MediaDistributionComponent {
	breakdown: MediaDistributionBreakdown;
	total: number;
}

interface RowEmissionsBreakdownDetail {
	adSelection: ComponentWithBreakdown;
	compensated: Compensated;
	creativeDelivery: ComponentWithBreakdown;
	mediaDistribution: MediaDistributionComponent;
}

interface RowEmissionsBreakdown {
	breakdown: RowEmissionsBreakdownDetail;
	framework: string;
}

interface PolicyEvaluationData {
	propertyId: number;
	isMFA: boolean;
	isInventory: boolean;
	channel: string;
	channelStatus: string;
	benchmarksPercentile: number;
}

interface Internal {
	countryRegionGCO2PerKwh: number;
	countryRegionCountry: string;
	channel: string;
	deviceType: string;
	propertyId: number;
	propertyInventoryType: string;
	propertyName: string;
	benchmarkPercentile: number;
	isMFA: boolean;
	policyEvaluationData: PolicyEvaluationData;
}

interface Row {
	coverage: RowCoverage;
	emissionsBreakdown: RowEmissionsBreakdown;
	inventoryCoverage: string;
	policies: RowPolicy[];
	rowIdentifier: string;
	totalEmissions: number;
	internal: Internal;
}

interface MeasureResponse {
	coverage: Coverage;
	policies: Policy[];
	requestId: string;
	totalEmissions: number;
	totalEmissionsBreakdown: EmissionsBreakdown;
	rows: Row[];
}

export type {
	AdFormats,
	Channels,
	BaseMetric,
	TotalMetric,
	Coverage,
	Policy,
	RowPolicy,
	EmissionsTotals,
	EmissionsBreakdown,
	AdFormatCoverage,
	SimpleValue,
	Impressions,
	SupplyGraphDepth,
	SupplyGraph,
	RowCoverage,
	EmissionsComponent,
	CompensationBreakdown,
	CompensatedBreakdownStruct,
	Compensated,
	SimpleBreakdown,
	MediaDistributionBreakdown,
	ComponentWithBreakdown,
	MediaDistributionComponent,
	RowEmissionsBreakdownDetail,
	RowEmissionsBreakdown,
	PolicyEvaluationData,
	Internal,
	Row,
	MeasureResponse,
};
