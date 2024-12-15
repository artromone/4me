# Task Management Application

## Overview

A robust task management application with:
- RESTful API Server
- Command-Line Interface (CLI)
- PostgreSQL Database Integration
- Docker Support

## Features

- Create, read, update, and delete tasks
- Organize tasks into lists and groups
- CLI for quick task management
- Dockerized deployment

## Prerequisites

- Go 1.23+
- Docker (optional)
- PostgreSQL

## Local Development Setup

1. Clone the repository
   ```bash
   git clone https://github.com/artromone/4me.git
   cd 4me
   ```

2. Create `.env` file
   ```bash
   cp .env.example .env
   ```

3. Edit `.env` with your database credentials

4. Install dependencies
   ```bash
   go mod download
   ```

## Running the Application

### Docker Deployment
```bash
# Build and run with Docker Compose
docker-compose up --build
```

### Local Development
```bash
# Build the cli
make
```

# CLI Commands
```bash
./task-cli create-task "Buy groceries" --list 1
./task-cli list-tasks --list 1
```

## CLI Usage

### Tasks
```bash
# Create a task
./task-cli create-task "Meeting preparation" --list 1 --due 2024-01-15

# List tasks in a list
./task-cli list-tasks --list 1
```

### Lists
```bash
# Create a list
./task-cli create-list "Work Tasks" --group 1

# List groups
./task-cli list-groups
```

## API Endpoints

### Tasks
- `POST /tasks`: Create a task
- `GET /tasks?list_id=1`: List tasks in a list
- `GET /tasks/{id}`: Get a specific task
- `PUT /tasks/{id}`: Update a task
- `DELETE /tasks/{id}`: Delete a task

### Lists & Groups
- Similar CRUD endpoints for lists and groups

## Configuration

Configure via `.env` file or environment variables:
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `SERVER_PORT`: API server port

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
