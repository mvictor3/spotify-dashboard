import { useState, useEffect } from 'react'
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom"
import './App.css'

const Home = () => {
    return (
        <div>
            <h1>Welcome to Spotify Dashboard</h1>
            <p>Explore tracks and find new favorites!</p>
        </div>
    )
}

const Favorites = () => {
    const [favoritesData, setFavoritesData] = useState(null)
    useEffect(() => {
        fetch('/favorites')
            .then(res => res.json())
            .then(data => setFavoritesData(data))
    }, [])

    return (
        <div>
            <h1>Liked Songs</h1>
            {favoritesData ? (
                <div>
                    <h2>Your {favoritesData.count} Favorite Songs</h2>
                    {favoritesData.tracks && favoritesData.tracks.map((track) => (
                        <div key={track.id} className="track-card">
                            <h3>{track.name}</h3>
                            <p><strong>Artist:</strong> {track.artists[0]?.name}</p>
                            <p><strong>Album:</strong> {track.album?.name}</p>
                        </div>
                    ))}
                </div>
            ) : (
                <p>Loading favorites...</p>
            )}
        </div>
    )
}

const Tracks = () => {
    const [tracksData, setTracksData] = useState(null)
    const [searchQuery, setSearchQuery] = useState('Homecoming')

    useEffect(() => {
        fetch(`/spotify/tracks?q=${searchQuery}`)
            .then(res => res.json())
            .then(data => setTracksData(data))
    }, [searchQuery])

    return (
        <div>
            <h1>Search Tracks</h1>
            <input
                type="text"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                placeholder="Search for a song..."
                style={{ padding: "8px", marginBottom: "1rem", width: "300px" }}
            />
            {tracksData ? (
                <div>
                    <h2>Found {tracksData.count} tracks for "{tracksData.query}"</h2>
                    {tracksData.tracks.map((track) => (
                        <div key={track.id} className="track-card">
                            <h3>{track.name}</h3>
                            <p><strong>Artist:</strong> {track.artists[0].name}</p>
                            <p><strong>Album:</strong> {track.album.name}</p>
                            <p><strong>Popularity:</strong> {track.popularity}</p>
                        </div>
                    ))}
                </div>
            ) : (
                <p>Loading tracks...</p>
            )}
        </div>
    )
}

const Status = () => {
    const [statusData, setStatusData] = useState(null)
    useEffect(() => {
        fetch('/spotify/status')
            .then(res => res.json())
            .then(data => setStatusData(data))
    }, [])

    return (
        <div>
            <h1>Status</h1>
            {statusData ? (
                <div>
                    <h2>System Status: {statusData.status}</h2>
                    <p><strong>Spotify Connection:</strong> {statusData.spotify}</p>
                </div>
            ) : (
                <p>Checking status...</p>
            )}
        </div>
    )
}

function App() {
    return (
        <Router>
            <div className="app">
                <header style={{ padding: "1rem", background: "#1db954" }}>
                    <nav style={{ display: "flex", gap: "1rem" }}>
                        <Link to="/" style={{ color: "white" }}>Home</Link>
                        <Link to="/favorites" style={{ color: "white" }}>Liked Songs</Link>
                        <Link to="/tracks" style={{ color: "white" }}>Search</Link>
                        <Link to="/status" style={{ color: "white" }}>Status</Link>
                    </nav>
                </header>
                <main style={{ padding: "1rem" }}>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/favorites" element={<Favorites />} />
                        <Route path="/tracks" element={<Tracks />} />
                        <Route path="/status" element={<Status />} />
                    </Routes>
                </main>
            </div>
        </Router>
    )
}

export default App