import json
from dataclasses import dataclass
from typing import List, Optional

from dataclasses_json import DataClassJsonMixin, dataclass_json


@dataclass_json
@dataclass
class AdFormats:
    metric: str
    generic: int
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
    verifiedRate: Optional[float] = None


@dataclass_json
@dataclass
class TotalMetric:
    metric: str
    modeled: int
    skipped: int


@dataclass_json
@dataclass
class Coverage:
    totalImpressions: TotalMetric
    totalRows: TotalMetric
    adFormats: Optional[AdFormats] = None
    channels: Optional[Channels] = None
    mediaOwners: Optional[BaseMetric] = None
    properties: Optional[BaseMetric] = None
    sellers: Optional[BaseMetric] = None


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
    adSelection: Optional[float] = None
    compensated: Optional[float] = None
    creativeDelivery: Optional[float] = None
    disposal: Optional[float] = None
    mediaDistribution: Optional[float] = None
    techManipulation: Optional[float] = None


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
    vendor: Optional[str] = None
    verified: Optional[bool] = None


@dataclass_json
@dataclass
class SimpleValue:
    value: str
    verified: Optional[bool] = None


@dataclass_json
@dataclass
class Impressions:
    modeled: int
    processed: int
    skipped: int


@dataclass_json
@dataclass
class SupplyGraphDepth:
    averageDepth: Optional[float] = None
    maxDepth: Optional[int] = None
    minDepth: Optional[int] = None


@dataclass_json
@dataclass
class SupplyGraph:
    logical: SupplyGraphDepth
    technical: SupplyGraphDepth
    totalCount: Optional[int] = None


@dataclass_json
@dataclass
class RowCoverage:
    impressions: Impressions
    adFormat: Optional[AdFormatCoverage] = None
    channel: Optional[SimpleValue] = None
    compensationProvider: Optional[SimpleValue] = None
    inventoryClassification: Optional[SimpleValue] = None
    mediaOwner: Optional[SimpleValue] = None
    property: Optional[SimpleValue] = None
    seller: Optional[SimpleValue] = None
    supplyGraph: Optional[SupplyGraph] = None
    venueCategory: Optional[SimpleValue] = None


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
    adPlatform: Optional[EmissionsComponent] = None
    dataTransfer: Optional[EmissionsComponent] = None


@dataclass_json
@dataclass
class MediaDistributionBreakdown:
    corporate: Optional[EmissionsComponent] = None
    dataTransfer: Optional[EmissionsComponent] = None
    storageAndTransport: Optional[EmissionsComponent] = None


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
    adSelection: Optional[ComponentWithBreakdown] = None
    compensated: Optional[Compensated] = None
    creativeDelivery: Optional[ComponentWithBreakdown] = None
    disposal: Optional[ComponentWithBreakdown] = None
    mediaDistribution: Optional[MediaDistributionComponent] = None
    techManipulation: Optional[ComponentWithBreakdown] = None


@dataclass_json
@dataclass
class RowEmissionsBreakdown:
    breakdown: RowEmissionsBreakdownDetail
    framework: str


@dataclass_json
@dataclass
class Error:
    message: str


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
    countryRegionGCO2PerKwh: float
    countryRegionCountry: str
    channel: str
    deviceType: str
    propertyId: int
    propertyInventoryType: str
    propertyName: str
    matchedInventoryId: str
    benchmarkPercentile: int
    isMFA: bool
    policyEvaluationData: PolicyEvaluationData


@dataclass_json
@dataclass
class Row:
    inventoryCoverage: str
    coverage: Optional[RowCoverage] = None
    emissionsBreakdown: Optional[RowEmissionsBreakdown] = None
    error: Optional[Error] = None
    internal: Optional[Internal] = None
    policies: Optional[List[RowPolicy]] = None
    rowIdentifier: Optional[str] = None
    totalEmissions: Optional[float] = None


@dataclass
class MeasureResponse(DataClassJsonMixin):
    totalEmissions: float
    totalEmissionsBreakdown: EmissionsBreakdown
    coverage: Optional[Coverage] = None
    policies: Optional[List[Policy]] = None
    requestId: Optional[str] = None
    rows: Optional[List[Row]] = None

    def __str__(self) -> str:
        print(json.dumps(self.to_dict(), indent=2))
        return (
            f"RequestID: {self.requestId}\n"
            f"Total Emissions: {self.totalEmissions:.4f}\n"
            f"Breakdown:\n"
            f"  - Ad Selection: {self.totalEmissionsBreakdown.totals.adSelection:.4f}\n"
            f"  - Creative Delivery: {self.totalEmissionsBreakdown.totals.creativeDelivery:.4f}\n"
            f"  - Media Distribution: {self.totalEmissionsBreakdown.totals.mediaDistribution:.4f}\n"
        )
