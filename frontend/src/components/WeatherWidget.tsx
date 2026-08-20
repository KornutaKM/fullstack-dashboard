import { useQuery } from '@tanstack/react-query'
import type { WeatherResponse } from '../types'
import { api } from '../api/client'


export function WeatherWidget() {
	const { data, isLoading, error } = useQuery<WeatherResponse>({
		queryKey: ['weather', 'Perm'],
		queryFn: async () => {
			const response = await api.getWeather('Perm')
			return response.data
		},
	})

	if (isLoading) {
		return <div className='p-4 bg-white rounded-lg shadow'>
			Загрузка погоды...
		</div>
	}

	if (error) {
		return <div className='p-4 bg-white rounded-lg shadow text-red-500'>
			Ошибка: {error.message}
		</div>
	}

	return (
		<div className="p-4 bg-white rounded-lg shadow">
			<h2 className="text-xl font-bold mb-2">🌤️ Погода в {data?.name}</h2>
			<div className="flex items-center gap-4">
				<div className="text-4xl">{Math.round(data?.main.temp ?? 0)}°C</div>
				<div className="text-gray-600">
					{data?.weather[0]?.description}
					<br />
					Ощущается как {Math.round(data?.main.feels_like ?? 0)}°C
				</div>
			</div>
		</div>
	)

}
