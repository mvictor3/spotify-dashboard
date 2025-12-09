Spotify Dashboard Setup

PostgreSQL Configuration
This project uses PostgreSQL. You can use any PostgreSQL user and database name.

Default setup (if you just installed PostgreSQL):
- User: postgres (default superuser)
- Database: postgres (default database)
- Port: 5432 (default port)

Custom setup(optional, project uses default):
- Use any existing PostgreSQL user
- Create a new database: CREATE DATABASE spotify_dashboard;
- Update environment variables accordingly

Spotfiy Developer API Setup:
- Go to https://developer.spotify.com/ and setup an account
- Create a project 
- Name it Wrapped
- Description can be what you please
- Make http://127.0.0.1:8080 your website
- Make http://127.0.0.1:8080/callback your Redirect URIs
- Acquire your client ID and client secret when you click on your project
- Set APIs used to Web API & Web Playback SDK

Environment Variables
Set these based on YOUR PostgreSQL setup:
- set POSTGRES_HOST=localhost
- set POSTGRES_PORT=5432
- set POSTGRES_USER=postgres       ← Change to your PostgreSQL username
- set POSTGRES_PASSWORD=yourpass   ← Your PostgreSQL password
- set POSTGRES_DB=postgres         ← Or create a custom database

Set these up for your spotify auth:
- set SPOTIFY_CLIENT_ID=your client id
- set SPOTIFY_CLIENT_SECRET=your client secret


Don't have PostgreSQL?
- Download: https://www.postgresql.org/download/
- Install with default settings
- Remember the password you set during installation
- Use postgres as the user (default)

Frontend Setup (Vite + React)
- npm create vite@latest . -- --template react
- npm install
- npm install react-router-dom


Running the project:

Backend:
- Run with go run main.go

Frontend:
- Run npm run dev

Favoriting Songs:

Favoriting is not yet implement in Frontend UI

To favorite a song via API:
- Use curl to favroite a song in the CMD terminal of backend using the track name
- Use this following command in your terminal:
- example: curl -X POST http://localhost:8080/favorites/save \ -H "Content-Type: application/json" \ -d '{"track_id":"Levitating Dua Lipa"}'

Make sure your backend server is running on localhost:8080 first.
