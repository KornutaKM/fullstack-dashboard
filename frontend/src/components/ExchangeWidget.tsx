import { useQuery } from '@tanstack/react-query'
import type { ExchangeResponse } from '../types'
import { api } from '../api/client'


export function ExchangeWidget() {
	const { data, isLoading, error, refetch } = useQuery<ExchangeResponse>({
		queryKey: ['exchange', 'base'],
		queryFn: async () => {
			const response = await api.getExchange('USD')
			return response.data
		},
		refetchInterval: 1000 * 30,
	})

	if (isLoading) {
		return <div className='p-4 bg-white rounded-lg shadow'>
			Загрузка курсов валют...
		</div>
	}

	if (error) {
		return <div className='p-4 bg-white rounded-lg shadow text-red-500'>
			Ошибка: {error.message}
		</div>
	}

	if (!data) {
		return (
			<div className='p-4 bg-white rounded-lg shadow'>Нет данных о курсах валют</div>
		)
	}

	return (
		<div className="p-4 bg-white rounded-lg shadow">
			<h2 className="text-xl font-bold mb-2">💱 Курс валют</h2>
			<p className="text-sm text-gray-500 mb-3">
				Базовая валюта: {data.base_code}
			</p>
			<button
				onClick={() => refetch()}
				className="text-sm text-blue-500 hover:text-blue-700"
			>
				🔄 Обновить
			</button>
			<ul className="space-y-1">
				<li className="flex justify-between">
					<span>USD → EUR</span>
					<span className="font-mono">{data.conversion_rates.EUR}</span>
				</li>
				<li className="flex justify-between">
					<span>USD → GBP</span>
					<span className="font-mono">{data.conversion_rates.GBP}</span>
				</li>
				<li className="flex justify-between">
					<span>USD → RUB</span>
					<span className="font-mono">{data.conversion_rates.RUB}</span>
				</li>
			</ul>
			<p className="text-xs text-gray-400 mt-3">
				Обновлено: {new Date().toLocaleTimeString()}
			</p>
		</div>
	)
}