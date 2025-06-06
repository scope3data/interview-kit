import axios, { type AxiosInstance } from "axios";
import type { MeasureResponse } from "./response-types";
import { config } from '../config'
import { logger } from '../logger'

const BASE_URL = "https://api.scope3.com/v2";

interface MeasureRow {
	rowIdentifier: string;
	impressions: number;
	utcDatetime: string;
	inventoryId: string;
	appStore: string;
	country: string;
	region: string;
	deviceType: string;
	channel: string;
	network: string;
}

interface RequestRow {
	inventoryId: string;
	impressions: number;
	deviceType: string;
	rowIdentifier: string;
	utcDatetime: string;
}

interface MeasureQueryParams {
	includeRows: boolean;
	latest: boolean;
	fields: string;
	framework: string;
}

const defaultMeasureQueryParams: MeasureQueryParams = {
	includeRows: true,
	latest: false,
	fields: "all",
	framework: "scope3",
};

class Client {
	private readonly httpClient: AxiosInstance;

	constructor() {
		this.httpClient = axios.create({
			baseURL: BASE_URL,
			headers: {
				Authorization: `Bearer ${config.apiKey}`,
				"Content-Type": "application/json",
				Accept: "application/json",
			},
		});
	}

	static new(): Client {
		return new Client();
	}

	private static createRequestRow(
		inventoryId: string,
		utcDatetime: string,
	): RequestRow {
		return {
			inventoryId,
			impressions: 1,
			deviceType: "pc",
			rowIdentifier: inventoryId,
			utcDatetime,
		};
	}

	async measure(
		inventoryIds: string[],
		requestDate: string,
	): Promise<MeasureResponse> {
		const rows = inventoryIds.map((id) =>
			Client.createRequestRow(id, requestDate),
		);

		const jsonData = {
			rows: rows.map((row) => ({
				inventoryId: row.inventoryId,
				impressions: row.impressions,
				deviceType: row.deviceType,
				rowIdentifier: row.rowIdentifier,
				utcDatetime: row.utcDatetime,
			})),
		};

		logger.debug("Request Details", {
			url: `${BASE_URL}/measure`,
			body: jsonData,
		});

		try {
			const response = await this.httpClient.post<MeasureResponse>(
				"/measure",
				jsonData,
				{
					params: defaultMeasureQueryParams,
				},
			);

			return response.data as MeasureResponse;
		} catch (error) {
			if (axios.isAxiosError(error)) {
				const errorMessage = error.response?.data
					? `API request failed: ${error.message}, body: ${JSON.stringify(error.response.data)}`
					: `API request failed: ${error.message}`;
				throw new Error(errorMessage);
			}
			throw error;
		}
	}
}

export default Client.new();
export type { MeasureRow, RequestRow, MeasureQueryParams };
