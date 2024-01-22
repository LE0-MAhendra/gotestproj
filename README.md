
## Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

- Go installed
- PostgreSQL installed
- Make installed (optional but recommended)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/your-go-project.git
   cd your-go-project
   ```

2. Copy the `.env.copy` file:

   ```bash
   cp .env.copy .env
   ```

3. Open the `.env` file and update the `DB_URL` with your database connection details, including username and password.

### Running Migrations

If you have Make installed, you can use the provided Makefile to simplify the process.

1. Open the Makefile and add your database URL or change upon your requirements

2. Run the migration:

   ```bash
   make migrate
   ```

### Starting the Project

1. Start the project:

   ```bash
   make start
   ```

The project should now be running locally.
