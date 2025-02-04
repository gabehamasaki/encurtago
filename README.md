# EncurtaGo

EncurtaGo is a URL shortener built with **Go (Gin Gonic)** for the backend and **Vite with React** for the frontend. It provides a RESTful API, a fast frontend interface, and a PostgreSQL database for persistent URL storage.

The application is live at: [EncurtaGo](https://encurtago.onrender.com/).

---

## 🚀 Features

- **URL Shortening** – Convert long URLs into short, shareable links.
- **REST API** – Provides endpoints for creating and retrieving shortened URLs.
- **Frontend with Vite & React** – Fast, modern UI for users.
- **PostgreSQL Storage** – Persistent URL mappings using a relational database.
- **Migrations with Tern** – Efficient database version control.
- **Code Generation with SQLC** – Type-safe queries and better performance.
- **Docker Support** – Easy deployment and environment consistency.

---

## 🛠️ Technologies Used

### **Backend:**
- **Go** – Fast and efficient backend service.
- **Gin Gonic** – Lightweight and high-performance web framework.
- **SQLC** – Type-safe SQL query generation.
- **PostgreSQL** – Relational database for storing URLs.
- **Tern** – Database migration tool.
- **Docker** – Containerized environment for deployment.

### **Frontend:**
- **Vite** – Fast development build tool.
- **React** – Frontend library for dynamic UI.
- **TypeScript** – Ensures type safety and maintainability.
- **Tailwind CSS** – Utility-first CSS for styling.

---

## 📂 Project Structure

```
encurtago/
├── client/             # Frontend (Vite + React)
├── cmd/api/            # Backend (Gin Gonic API)
├── internal/           # Internal Go modules and utilities
├── test/               # Test cases
├── Dockerfile          # Docker container setup
├── docker-compose.yml  # Docker Compose configuration
├── .env.example        # Example environment variables
└── Makefile            # Build automation commands
```

---

## 🔧 Setup & Installation

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/gabehamasaki/encurtago.git
cd encurtago
```

### 2️⃣ Set Up Environment Variables

- Copy `.env.example` and rename it to `.env`:

  ```bash
  cp .env.example .env
  ```

- Update the `.env` file with the correct database URL and configurations.

### 3️⃣ Run the Project with Docker

```bash
docker-compose up --build
```

This command builds and starts the services defined in `docker-compose.yml`.

### 4️⃣ Access the Application

- **URL**: [http://localhost:8080](http://localhost:8080)

---

## 🧪 Running Tests

To run backend tests:

```bash
go test ./...
```

---

## 🤝 Contributing

Contributions are welcome! Feel free to:

- Report issues
- Submit feature requests
- Fork the repository and create pull requests

---

## 📜 License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for more details.

---

*Built with ❤️ by [Gabriel Hamasaki](https://github.com/gabehamasaki).*
