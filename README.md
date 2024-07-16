# Dish Dashboard

## Overview

Dish Dashboard is a web application for managing dishes. The application allows users to view the list of dishes, toggle their publication status, and receive real-time updates via WebSockets. The backend is built with Go using the Gin framework, and the frontend is built with React, vite and Material Tailwind.

## Features

- **View Dishes:** Display a list of dishes with their details.
- **Toggle Publication Status:** Change the publication status of dishes.
- **Real-time Updates:** Receive live updates on dish status changes via WebSockets.

## Technologies Used

- **Frontend:** React, Vite, Material Tailwind for UI components.
- **Backend:** Go (Gin framework), PostgreSQL for database (GORM library).
- **Real-Time:** WebSockets (gorilla/websocket library).

# Backend

## Structure

- **Main Application (main.go):** Initializes the environment, sets up the router, and starts the server.
- **Database Connection (database/database.go):** Handles the connection to the PostgreSQL database using GORM.
- **Models (models/models.go):** Defines the Dish model.
- **Routes (routes/routes.go):** Contains handlers for HTTP endpoints and WebSocket connections.

## Key Functions

- **GetDishes:** Handles GET requests to fetch all dishes from the database.
- **ToggleDishStatus:** Handles PUT requests to toggle the publication status of a dish.
- **HandleConnections:** Manages WebSocket connections and listens for messages.
- **HandleMessages:** Broadcasts updates to all connected WebSocket clients.

## Running the Server

**go run main.go**
The server runs on http://localhost:8005

# Front-End Dashboard

The front-end dashboard is built using React.js and Vite, leveraging Material Tailwind for UI components. It includes:

##

-**Header Component:** Displays header. -**Dishes Component:** Main section for managing dish information. -**Footer Component:** Displays footer.
