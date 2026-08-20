import { ExchangeWidget } from './components/ExchangeWidget'
import { NewsWidget } from './components/NewsWidget'
import { WeatherWidget } from './components/WeatherWidget'

function App() {
  return (
    <div className="min-h-screen bg-gray-100 p-4">
      <div className="max-w-6xl mx-auto">
        <h1 className="text-3xl font-bold text-gray-800 mb-6">
          📊 Личный дэшборд
        </h1>

        {/* Сетка: 1 колонка на мобилках, 2 колонки на планшетах, 3 колонки на десктопе */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          <WeatherWidget />
          <NewsWidget />
          <ExchangeWidget />
        </div>
      </div>
    </div>
  )
}

export default App
