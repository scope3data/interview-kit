from dataclasses import dataclass
from typing import List
from dataclasses_json import DataClassJsonMixin, dataclass_json

@dataclass_json
@dataclass
class AdFormats:
    generic: int
    metric: str
    unknown: int
    vendorSpecific: int

@dataclass_json
@dataclass
class Channels:
    deprecated: int
    metric: str
    modeled: int
    unknown: int

@dataclass_json
@dataclass
class BaseMetric:
    metric: str
    modeled: int
    unknown: int

@dataclass_json
@dataclass
class TotalMetric:
    metric: str
    modeled: int
    skipped: int

@dataclass_json
@dataclass
class Coverage:
    adFormats: AdFormats
    channels: Channels
    mediaOwners: BaseMetric
    properties: BaseMetric
    sellers: BaseMetric
    totalImpressions: TotalMetric
    totalRows: TotalMetric

@dataclass_json
@dataclass
class Policy:
    compliant: int
    noncompliant: int
    policy: str
    policyOwner: str

@dataclass_json
@dataclass
class RowPolicy:
    compliant: bool
    policy: str
    policyOwner: str

@dataclass_json
@dataclass
class EmissionsTotals:
    adSelection: float
    creativeDelivery: float
    mediaDistribution: float

@dataclass_json
@dataclass
class EmissionsBreakdown:
    framework: str
    totals: EmissionsTotals

@dataclass_json
@dataclass
class AdFormatCoverage:
    name: str
    value: str
    verified: bool

@dataclass_json
@dataclass
class SimpleValue:
    value: str

@dataclass_json
@dataclass
class Impressions:
    modeled: int
    processed: int
    skipped: int

@dataclass_json
@dataclass
class SupplyGraphDepth:
    averageDepth: float
    maxDepth: int
    minDepth: int

@dataclass_json
@dataclass
class SupplyGraph:
    logical: SupplyGraphDepth
    technical: SupplyGraphDepth
    totalCount: int

@dataclass_json
@dataclass
class RowCoverage:
    adFormat: AdFormatCoverage
    channel: SimpleValue
    compensationProvider: SimpleValue
    impressions: Impressions
    property: SimpleValue
    supplyGraph: SupplyGraph

@dataclass_json
@dataclass
class EmissionsComponent:
    emissions: float

@dataclass_json
@dataclass
class CompensationBreakdown:
    emissions: float
    provider: str

@dataclass_json
@dataclass
class CompensatedBreakdownStruct:
    compensation: CompensationBreakdown

@dataclass_json
@dataclass
class Compensated:
    breakdown: CompensatedBreakdownStruct
    total: float

@dataclass_json
@dataclass
class SimpleBreakdown:
    adPlatform: EmissionsComponent
    dataTransfer: EmissionsComponent

@dataclass_json
@dataclass
class MediaDistributionBreakdown:
    corporate: EmissionsComponent
    dataTransfer: EmissionsComponent

@dataclass_json
@dataclass
class ComponentWithBreakdown:
    breakdown: SimpleBreakdown
    total: float

@dataclass_json
@dataclass
class MediaDistributionComponent:
    breakdown: MediaDistributionBreakdown
    total: float

@dataclass_json
@dataclass
class RowEmissionsBreakdownDetail:
    adSelection: ComponentWithBreakdown
    compensated: Compensated
    creativeDelivery: ComponentWithBreakdown
    mediaDistribution: MediaDistributionComponent

@dataclass_json
@dataclass
class RowEmissionsBreakdown:
    breakdown: RowEmissionsBreakdownDetail
    framework: str

@dataclass_json
@dataclass
class PolicyEvaluationData:
    propertyId: int
    isMFA: bool
    isInventory: bool
    channel: str
    channelStatus: str
    benchmarksPercentile: int

@dataclass_json
@dataclass
class Internal:
    countryRegionGCO2PerKwh: int
    countryRegionCountry: str
    channel: str
    deviceType: str
    propertyId: int
    propertyInventoryType: str
    propertyName: str
    benchmarkPercentile: int
    isMFA: bool
    policyEvaluationData: PolicyEvaluationData

@dataclass_json
@dataclass
class Row:
    coverage: RowCoverage
    emissionsBreakdown: RowEmissionsBreakdown
    inventoryCoverage: str
    policies: List[RowPolicy]
    rowIdentifier: str
    totalEmissions: float
    internal: Internal

@dataclass
class MeasureResponse(DataClassJsonMixin):
    coverage: Coverage
    policies: List[Policy]
    requestId: str
    totalEmissions: float
    totalEmissionsBreakdown: EmissionsBreakdown
    rows: List[Row]

    def __str__(self) -> str:
        return (
            f"RequestID: {self.requestId}\n"
            f"Total Emissions: {self.totalEmissions:.4f}\n"
            f"Breakdown:\n"
            f"  - Ad Selection: {self.totalEmissionsBreakdown.totals.adSelection:.4f}\n"
            f"  - Creative Delivery: {self.totalEmissionsBreakdown.totals.creativeDelivery:.4f}\n"
            f"  - Media Distribution: {self.totalEmissionsBreakdown.totals.mediaDistribution:.4f}\n"
        )
