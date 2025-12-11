Spotify Dashboard Setup

Live Demo:
https://spotify-dashboard-phi.vercel.app/

About:
A Music Searching App that lets you search for tracks and save favorites.

Features:
- Search for music using Spotify's API
- Save favorite tracks to PostgreSQL database
- Status monitoring page
- FULL REST API backend

Current Status:
- Backend: Fully functional with search, favorite and delete endpoints
- Frontend: Complete UI with search, favorites, and delete functionality
- Database: PostgreSql integration with proper framework

Tech Stack:
- Backend: Go, PostgreSQL
- Frontend: React, Vite
- API: Spotify Web API

PostgreSQL Configuration
This project uses PostgreSQL. You can use any PostgreSQL user and database name.

Default setup (if you just installed PostgreSQL):
If you just installed PostgreSQL, use the default configuration:
- User: postgres (default superuser)
- Database: postgres (default database)
- Port: 5432 (default port)

Spotify Developer API Setup:
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

Frontend Setup (React)
- npm install

Running the project:

Backend:
- Run with go run main.go
- runs localhost:8080

Frontend:
- Run npm run dev
- runs localhost:5173
