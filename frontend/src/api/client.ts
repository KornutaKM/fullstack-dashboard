import axios from 'axios'
import type { ExchangeResponse, NewsResponse, WeatherResponse } from '../types'

const apiClient = axios.create({
	baseURL: "http://localhost:8080/api",
	headers: {
		"Content-Type": 'application/json',
	},
})

export const api = {

	// Get weather
	getWeather: (city: string = 'Moscow') =>
		apiClient.get<WeatherResponse>('/weather', { params: { city } }),

	getNews: (country: string = 'us', category: string = 'general') =>
		apiClient.get<NewsResponse>('/news', { params: { country, category } }),

	getExchange: (base: string = 'USD') =>
		apiClient.get<ExchangeResponse>('/exchange', { params: { base } }),
}