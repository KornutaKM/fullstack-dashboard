import { useQuery } from '@tanstack/react-query'
import type { NewsResponse } from '../types'
import { api } from '../api/client'


export function NewsWidget() {
	const { data, isLoading, error, refetch } = useQuery<NewsResponse>({
		queryKey: ["news", 'russia'],
		queryFn: async () => {
			const response = await api.getEverything('Russia', 'ru')
			return response.data
		},
		refetchInterval: 1000 * 30,
	})

	if (isLoading) {
		return <div className='p-4 bg-white rounded-lg shadow'>
			Загрузка новостей...
		</div>
	}

	if (error) {
		return <div className='p-4 bg-white rounded-lg shadow text-red-500'>
			Ошибка: {error.message}
		</div>
	}

	if (!data || data.articles.length === 0) {
		return (
			<div className="p-4 bg-white rounded-lg shadow">
				<h2 className="text-xl font-bold mb-2">📰 Новости</h2>
				<p className="text-gray-500">Новостей пока нет</p>
			</div>
		)
	}
	return (
		<div className="p-4 bg-white rounded-lg shadow">
			<h2 className="text-xl font-bold mb-2">📰 Новости</h2>
			<button
				onClick={() => refetch()}
				className="text-sm text-blue-500 hover:text-blue-700"
			>
				🔄 Обновить
			</button>
			<ul className="space-y-2">
				{data.articles.map((article, index) => (
					<li key={index} className="border-b border-gray-100 pb-2 last:border-0">
						<a
							href={article.url}
							target="_blank"
							rel="noopener noreferrer"
							className="text-blue-600 hover:underline"
						>
							{article.title}
						</a>
						<p className="text-sm text-gray-500 mt-1">
							{article.source.name}
						</p>
					</li>
				))}
			</ul>
			<p className="text-xs text-gray-400 mt-3">
				Обновлено: {new Date().toLocaleTimeString()}
			</p>
		</div>
	)
}