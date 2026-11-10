import { useState, useEffect } from 'react'
import { Navbar } from './components/Navbar'
import { CreateNoteModal } from './components/CreateNoteModal'

export interface Note {
  id: number
  title: string
  content: string
  created_at: string
  updated_at: string
}

function App() {
  const [notes, setNotes] = useState<Note[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [searchQuery, setSearchQuery] = useState('')
  const [activeFilter, setActiveFilter] = useState('all')
  const [isCreateModalOpen, setIsCreateModalOpen] = useState(false)

  // Fetch all notes from API
  const fetchNotes = async () => {
    try {
      setLoading(true)
      const res = await fetch('/api/notes')
      if (!res.ok) throw new Error('Failed to fetch notes')
      const data: Note[] = await res.json()
      setNotes(data)
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message)
      } else {
        setError('An unexpected error occurred')
      }
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchNotes()
  }, [])

  // Create note handler
  const handleCreateNote = async (title: string, content: string) => {
    const res = await fetch('/api/notes', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title, content }),
    })

    if (!res.ok) {
      const errorData = await res.json().catch(() => ({}))
      throw new Error(errorData.error || 'Failed to create note')
    }

    const createdNote: Note = await res.json()
    setNotes((prev) => [createdNote, ...prev])
  }

  // Delete note handler
  const handleDeleteNote = async (id: number) => {
    if (!confirm('Are you sure you want to delete this note?')) return

    try {
      const res = await fetch(`/api/notes/${id}`, {
        method: 'DELETE',
      })
      if (!res.ok) throw new Error('Failed to delete note')
      setNotes((prev) => prev.filter((n) => n.id !== id))
    } catch (err: unknown) {
      alert(err instanceof Error ? err.message : 'Failed to delete note')
    }
  }

  // Filter notes based on search query
  const filteredNotes = notes.filter((note) => {
    const q = searchQuery.toLowerCase().trim()
    if (!q) return true
    return (
      note.title.toLowerCase().includes(q) ||
      note.content.toLowerCase().includes(q)
    )
  })

  return (
    <div className="min-h-screen bg-gray-50 text-gray-900">
      <Navbar
        onSearch={setSearchQuery}
        onNewNote={() => setIsCreateModalOpen(true)}
        activeFilter={activeFilter}
        onFilterChange={setActiveFilter}
      />

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="flex items-center justify-between mb-8">
          <div>
            <h1 className="text-2xl font-bold tracking-tight text-gray-900 capitalize">
              {activeFilter === 'all' ? 'All Notes' : activeFilter}
            </h1>
            <p className="text-sm text-gray-500 mt-1">
              Capture your thoughts, ideas, and tasks.
            </p>
          </div>

          <button
            onClick={() => setIsCreateModalOpen(true)}
            className="inline-flex items-center gap-2 px-4 py-2 bg-indigo-600 hover:bg-indigo-700 active:scale-95 text-white text-sm font-medium rounded-xl shadow-sm shadow-indigo-100 transition cursor-pointer"
          >
            <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2.5} d="M12 4v16m8-8H4" />
            </svg>
            Create Note
          </button>
        </div>

        {/* Error message */}
        {error && (
          <div className="mb-6 p-4 rounded-xl bg-red-50 border border-red-200 text-red-700 text-sm flex items-center justify-between">
            <span>{error}</span>
            <button 
              onClick={fetchNotes} 
              className="text-xs font-semibold underline hover:no-underline"
            >
              Retry
            </button>
          </div>
        )}

        {/* Loading state */}
        {loading ? (
          <div className="flex flex-col items-center justify-center py-20 text-gray-400">
            <svg className="animate-spin h-8 w-8 text-indigo-600 mb-3" fill="none" viewBox="0 0 24 24">
              <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
              <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z"></path>
            </svg>
            <p className="text-sm">Loading your notes...</p>
          </div>
        ) : filteredNotes.length === 0 ? (
          /* Empty placeholder card */
          <div className="rounded-2xl border-2 border-dashed border-gray-200 p-12 text-center bg-white/50">
            <div className="mx-auto h-12 w-12 rounded-xl bg-indigo-50 text-indigo-600 flex items-center justify-center mb-4">
              <svg className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
              </svg>
            </div>
            <h3 className="text-base font-semibold text-gray-900">
              {searchQuery ? 'No notes matched your search' : 'No notes yet'}
            </h3>
            <p className="mt-1 text-sm text-gray-500">
              {searchQuery ? 'Try searching for something else.' : 'Get started by creating your first note.'}
            </p>
            {!searchQuery && (
              <div className="mt-6">
                <button
                  onClick={() => setIsCreateModalOpen(true)}
                  className="inline-flex items-center px-4 py-2 text-sm font-medium rounded-xl text-white bg-indigo-600 hover:bg-indigo-700 shadow-sm transition cursor-pointer"
                >
                  + Create Note
                </button>
              </div>
            )}
          </div>
        ) : (
          /* Notes Grid */
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {filteredNotes.map((note) => (
              <div
                key={note.id}
                className="group relative bg-white p-6 rounded-2xl border border-gray-200/80 shadow-xs hover:shadow-md hover:border-gray-300 transition-all flex flex-col justify-between"
              >
                <div>
                  <div className="flex items-start justify-between gap-3 mb-2">
                    <h3 className="font-semibold text-gray-900 text-lg leading-snug break-words">
                      {note.title}
                    </h3>
                    <button
                      onClick={() => handleDeleteNote(note.id)}
                      className="opacity-0 group-hover:opacity-100 text-gray-400 hover:text-red-600 p-1 rounded-lg hover:bg-red-50 transition cursor-pointer"
                      title="Delete note"
                    >
                      <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                      </svg>
                    </button>
                  </div>
                  <p className="text-gray-600 text-sm whitespace-pre-wrap leading-relaxed line-clamp-6">
                    {note.content || <span className="italic text-gray-400">No content</span>}
                  </p>
                </div>

                <div className="mt-4 pt-3 border-t border-gray-100 flex items-center justify-between text-xs text-gray-400">
                  <span>
                    {new Date(note.created_at).toLocaleDateString(undefined, {
                      month: 'short',
                      day: 'numeric',
                      year: 'numeric',
                    })}
                  </span>
                  <span className="text-[11px] text-gray-400">
                    #{note.id}
                  </span>
                </div>
              </div>
            ))}
          </div>
        )}
      </main>

      {/* Modal Dialog for Note Creation */}
      <CreateNoteModal
        isOpen={isCreateModalOpen}
        onClose={() => setIsCreateModalOpen(false)}
        onSubmit={handleCreateNote}
      />
    </div>
  )
}

export default App
