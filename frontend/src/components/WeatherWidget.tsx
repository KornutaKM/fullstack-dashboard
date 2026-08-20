import { useQuery } from '@tanstack/react-query'
import { useState } from 'react'
import { api } from '../api/client'
import type { WeatherResponse } from '../types'

export function WeatherWidget() {
  // Время последнего обновления
  const [lastUpdated, setLastUpdated] = useState<string>(
    new Date().toLocaleTimeString()
  )

  const { data, isLoading, error, refetch, isFetching } = useQuery<WeatherResponse>({
    queryKey: ['weather', 'Moscow'],
    queryFn: async () => {
      const response = await api.getWeather('Moscow')
      return response.data
    },
    refetchInterval: 1000 * 30,
  })

  // Функция обновления
  const handleRefresh = async () => {
    await refetch() // ждём, пока данные обновятся
    setLastUpdated(new Date().toLocaleTimeString()) // обновляем время
  }

  if (isLoading) {
    return <div className="p-4 bg-white rounded-lg shadow">Загрузка погоды...</div>
  }

  if (error) {
    return <div className="p-4 bg-white rounded-lg shadow text-red-500">Ошибка: {error.message}</div>
  }

  return (
    <div className="p-4 bg-white rounded-lg shadow">
      <div className="flex justify-between items-center mb-2">
        <h2 className="text-xl font-bold">🌤️ Погода в {data?.name}</h2>
        <button
          onClick={handleRefresh}
          className="text-sm text-blue-500 hover:text-blue-700 disabled:opacity-50"
          disabled={isFetching}
        >
          {isFetching ? '⏳ Обновление...' : '🔄 Обновить'}
        </button>
      </div>
      <div className="flex items-center gap-4">
        <div className="text-4xl">{Math.round(data?.main.temp ?? 0)}°C</div>
        <div className="text-gray-600">
          {data?.weather[0]?.description}
          <br />
          Ощущается как {Math.round(data?.main.feels_like ?? 0)}°C
        </div>
      </div>
      <p className="text-xs text-gray-400 mt-3">
        Обновлено: {lastUpdated}
        {isFetching && ' (обновляется...)'}
      </p>
    </div>
  )
}