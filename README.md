# Mortgage Underwriting Engine

A full-stack web application designed to automate mortgage underwriting decisions based on key financial metrics. The system takes borrower details via a React frontend, processes the logic with a Go backend, and provides an `Approve`, `Refer`, or `Decline` decision with supporting reasons. All services are containerized with Docker.

---

## ✨ Features

-   **Evaluation Form:** Submit mortgage applications via a clean web form.
-   **Automatic Underwriting:** Instantly receive a decision (Approve/Refer/Decline) based on a configurable rules engine.
-   **Evaluation History:** View a complete, paginated history of all past evaluations.
-   **Detailed Breakdown:** Each historical entry shows the inputs, calculated DTI & LTV, and the reasons for the decision.

---

## 🚀 Tech Stack

-   **Frontend:** React, TypeScript, Vite, SCSS, React Router
-   **Backend:** Go, Gin Web Framework
-   **Database:** MySQL 8.0
-   **Containerization:** Docker & Docker Compose

---

## 🏁 Getting Started

Follow these instructions to get the project up and running on your local machine for development and testing purposes.

### Prerequisites

You must have the following software installed on your machine:
-   [Git](https://git-scm.com/)
-   [Docker](https://www.docker.com/products/docker-desktop/)
-   [Docker Compose](https://docs.docker.com/compose/)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/tu-usuario/tu-repositorio.git](https://github.com/tu-usuario/tu-repositorio.git)
    cd your-repository
    ```

2.  **Create the environment file:**
    Copy the example environment file and fill in the required values (especially passwords).
    ```bash
    cp .env-template .env
    ```

3.  **Build and run the containers:**
    This single command will build the frontend and backend images, and start all services (frontend, backend, database) in detached mode.
    ```bash
    docker-compose up --build -d
    ```

### Accessing the Application

-   **Frontend:** Open your browser and navigate to `http://localhost:5173`
-   **Backend API:** The API is accessible at `http://localhost:8080`.

---

## 🏛️ Architecture & Project Structure

The project follows a Clean Architecture / Hexagonal approach to ensure a clean separation of concerns.

-   **Go Backend:** The backend is structured into `domain`, `ports` (interfaces), `services` (business logic), and `adapters` (Gin API, MySQL repository). This makes the core logic independent of the database and the web framework.
-   **React Frontend:** The frontend is organized by features (`pages`, `components`, `layouts`, `hooks`) to ensure scalability and maintainability. Global styles and shared components are kept separate from page-specific logic.

---

## 🧠 Approach, Challenges, and Learnings

### Approach
I chose a decoupled, containerized architecture to mirror modern production environments. The Go backend was designed to be a stateless service whose core logic is completely isolated and unit-tested. The React frontend was built as a standalone Single Page Application (SPA) that communicates with the backend via a proxied API.

### Challenges
-   **Inter-container Communication:** Ensuring the frontend, backend, and database containers could communicate correctly. This was solved by using Docker Compose service names as hostnames and managing dependencies with `depends_on` and `healthcheck`.
-   **Full-Stack Debugging:** Tracing errors through the entire stack, from a frontend `fetch` call failing due to browser caching (`304 Not Modified`) to a backend runtime panic (divide by zero) caused by incorrect data types in the JSON payload. This required inspecting the browser's Network Tab and `docker-compose logs` in tandem.
-   **Development Workflow:** Setting up a smooth development experience with hot-reloading for the frontend while inside Docker, which was accomplished using a `docker-compose.override.yml` file and Docker volumes.

### Learnings
-   **The Power of Decoupling:** The true value of the Hexagonal Architecture became clear when the persistence layer was defined. The core business logic (`service`) depends only on an interface (`port`), making it trivial to swap out database implementations without affecting the core rules engine.
-   **Importance of Data Contracts:** A simple mismatch between the data types expected by the Go backend (numbers) and those sent by the frontend (strings from form inputs) can cause runtime failures. Explicitly parsing and validating data at the boundaries is crucial.
-   **Effective Docker Debugging:** This project reinforced the importance of using `docker ps` to check container status and `docker-compose logs <service_name>` to pinpoint the exact source of an error within a containerized application.