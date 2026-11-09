import { useState } from 'react'
import viteLogo from './assets/vite.svg'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="min-h-screen bg-slate-900 text-white flex flex-col items-center justify-center p-4">
      <div className="max-w-md w-full bg-slate-800 rounded-2xl p-8 shadow-2xl border border-slate-700 text-center">
        <div className="flex justify-center mb-6">
          <a href="https://vite.dev" target="_blank" rel="noreferrer">
            <img src={viteLogo} className="h-16 w-16 hover:scale-110 transition-transform" alt="Vite logo" />
          </a>
        </div>
        <h1 className="text-3xl font-bold mb-2 bg-gradient-to-r from-cyan-400 to-blue-500 bg-clip-text text-transparent">
          Tailwind CSS Ready
        </h1>
        <p className="text-slate-400 mb-6">
          Tailwind CSS v4 is set up and working with Vite and React!
        </p>
        <button
          type="button"
          onClick={() => setCount((c) => c + 1)}
          className="px-5 py-2.5 rounded-lg bg-blue-600 hover:bg-blue-500 active:scale-95 transition text-white font-medium shadow-md"
        >
          Count is {count}
        </button>
      </div>
    </div>
  )
}

export default App
