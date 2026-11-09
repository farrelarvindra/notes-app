import { useState } from 'react'
import { Navbar } from './components/Navbar'

function App() {
  const [activeFilter, setActiveFilter] = useState('all')

  const handleSearch = (query: string) => {
    console.log('Searching for:', query)
  }

  const handleNewNote = () => {
    console.log('Create new note clicked')
  }

  return (
    <div className="min-h-screen bg-gray-50 text-gray-900">
      <Navbar
        onSearch={handleSearch}
        onNewNote={handleNewNote}
        activeFilter={activeFilter}
        onFilterChange={setActiveFilter}
      />

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="flex items-center justify-between mb-6">
          <div>
            <h1 className="text-2xl font-bold tracking-tight text-gray-900 capitalize">
              {activeFilter === 'all' ? 'All Notes' : activeFilter}
            </h1>
            <p className="text-sm text-gray-500 mt-1">
              Capture your thoughts, ideas, and tasks.
            </p>
          </div>
        </div>

        {/* Empty placeholder card */}
        <div className="rounded-2xl border-2 border-dashed border-gray-200 p-12 text-center">
          <div className="mx-auto h-12 w-12 rounded-xl bg-indigo-50 text-indigo-600 flex items-center justify-center mb-4">
            <svg
              className="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M12 4v16m8-8H4"
              />
            </svg>
          </div>
          <h3 className="text-base font-semibold text-gray-900">No notes yet</h3>
          <p className="mt-1 text-sm text-gray-500">
            Get started by creating your first note.
          </p>
          <div className="mt-6">
            <button
              onClick={handleNewNote}
              className="inline-flex items-center px-4 py-2 text-sm font-medium rounded-xl text-white bg-indigo-600 hover:bg-indigo-700 shadow-sm transition cursor-pointer"
            >
              + Create Note
            </button>
          </div>
        </div>
      </main>
    </div>
  )
}

export default App
