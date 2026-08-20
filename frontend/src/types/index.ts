export interface WeatherResponse {
	name: string
	main: {
		temp: number
		feels_like: number
	}
	weather: Array<{
		description: string
		icon: string
	}>
}

export interface NewsResponse {
	status: string
	totalResults: number
	articles: Array<{
		source: {
			name: string
		}
		author: string
		title: string
		description: string
		url: string
		urlToImage: string
		publishedAt: string
	}>
}

export interface ExchangeResponse {
	result: string
	base_code: string
	conversion_rates: {
		USD: number
		EUR: number
		GBP: number
		RUB: number
	}
	time_last_update_unix: number
	time_last_update_utc: string
	time_next_update_unix: number
	time_next_update_utc: string
}