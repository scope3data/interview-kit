import json
import logging
from typing import List
from dataclasses import dataclass
import requests

from config import config
from .response import MeasureResponse

# Configure logging
logger = logging.getLogger(__name__)

BASE_URL = "https://api.scope3.com/v2"

@dataclass
class MeasureRow:
    """Represents a full measure row with all possible fields."""
    row_identifier: str
    impressions: int
    utc_datetime: str
    inventory_id: str
    app_store: str
    country: str
    region: str
    device_type: str
    channel: str
    network: str

@dataclass
class RequestRow:
    """Represents a row in the measure request with required fields."""
    inventory_id: str
    impressions: int
    device_type: str
    row_identifier: str
    utc_datetime: str

    @classmethod
    def new(cls, inventory_id: str, utc_datetime: str) -> 'RequestRow':
        """Creates a new RequestRow with default values.

        Args:
            inventory_id (str): The inventory ID to measure
            utc_datetime (str): The UTC datetime for the measurement

        Returns:
            RequestRow: A new request row with default values set
        """
        return cls(
            inventory_id=inventory_id,
            impressions=1,
            device_type="pc",
            row_identifier=inventory_id,
            utc_datetime=utc_datetime
        )

@dataclass
class MeasureQueryParams:
    """Query parameters for the measure endpoint."""
    include_rows: bool = True
    latest: bool = False
    fields: str = "all"
    framework: str = "scope3"

    def to_dict(self) -> dict:
        """Converts the parameters to a dictionary for requests."""
        return {
            "includeRows": self.include_rows,
            "latest": self.latest,
            "fields": self.fields,
            "framework": self.framework
        }

class Client:
    def __init__(self):
        self.http_client = requests.Session()
        self.api_key = config.api_key

    @classmethod
    def new(cls) -> 'Client':
        """Factory method to create a new client instance."""
        return cls()

    def measure(self, inventory_ids: List[str], request_date: str) -> MeasureResponse:
            """Makes a measure request to the Scope3 API.

            Sends a POST request to the /measure endpoint to calculate emissions
            for the provided inventory IDs on the specified date.

            Args:
                inventory_ids (List[str]): List of inventory IDs to measure. Each ID should
                    be a valid domain name (e.g., ["yahoo.com", "google.com"]).
                request_date (str): The request date in YYYY-MM-DD format (e.g., "2025-05-01").

            Returns:
                MeasureResponse: A structured response containing the measurement data.

            Raises:
                Exception: If the API request fails due to network, authentication, or server issues.
            """
            rows = [
                RequestRow.new(inventory_id, request_date)
                for inventory_id in inventory_ids
            ]

            json_data = {
                "rows": [
                    {
                        "inventoryId": row.inventory_id,
                        "impressions": row.impressions,
                        "deviceType": row.device_type,
                        "rowIdentifier": row.row_identifier,
                        "utcDatetime": row.utc_datetime
                    }
                    for row in rows
                ]
            }

            # Build request URL with parameters
            params = MeasureQueryParams()
            request_url = f"{BASE_URL}/measure"

            logger.debug("Request Details", extra={
                "url": request_url,
                "body": json_data
            })

            headers = {
                "Authorization": f"Bearer {self.api_key}",
                "Content-Type": "application/json",
                "Accept": "application/json"
            }

            logger.debug("Request Headers", extra={"headers": headers})

            response = None
            try:
                response = self.http_client.post(
                    request_url,
                    json=json_data,
                    headers=headers,
                    params=params.to_dict()
                )
                response.raise_for_status()
                result = response.json()
                return MeasureResponse.from_dict(result)
            except requests.exceptions.RequestException as e:
                if e.response and hasattr(e.response, 'text'):
                    error_message = f"API request failed: {str(e)}, body: {e.response.text}"
                else:
                    error_message = f"API request failed: {str(e)}"
                raise Exception(error_message) from e
            except json.JSONDecodeError as e:
                if response:
                    error_message = f"Failed to decode response: {str(e)}, body: {response.text}"
                else:
                    error_message = f"Failed to decode response: {str(e)}"
                raise Exception(error_message) from e
