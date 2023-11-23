## PostgreSQL Database Interaction README

This Go program demonstrates basic interaction with a PostgreSQL database. It performs the following operations:

1. **Connection Setup:**
   - Reads database connection details (host, port, dbname, user, password) from a `.env` file using `github.com/joho/godotenv`.
   - Connects to the PostgreSQL database using `database/sql` and `github.com/jackc/pgx/v5/stdlib`.

2. **Database Ping:**
   - Tests the database connection by pinging it.

3. **Retrieve and Display Rows:**
   - Retrieves all rows from the "users" table and displays the data.

4. **Insert a Row:**
   - Inserts a new row into the "users" table with first name "Mike" and last name "Mom."

5. **Update a Row:**
   - Updates the first name of a user with ID 3 to "Terraform."

6. **Retrieve and Display Rows Again:**
   - Retrieves all rows from the "users" table after the insert and update operations.

7. **Query a Single Row:**
   - Retrieves and displays details of a user with ID 3 using `QueryRow`.

8. **Delete a Row:**
   - Deletes a row from the "users" table with ID 7.

9. **Final Rows Display:**
   - Retrieves and displays all rows from the "users" table after the delete operation.

### Running the Program

1. Ensure you have a PostgreSQL database set up.
2. Create a `.env` file in the same directory as the program with the following variables:

   ```
   DB_HOST=your_host
   DB_PORT=your_port
   DB_NAME=your_db_name
   DB_USER=your_user
   DB_PASSWORD=your_password
   ```

3. Run the program:

   ```bash
   go run main.go
   ```

   The program will connect to the specified PostgreSQL database, perform the operations, and display the results.

### Note

- Ensure the necessary PostgreSQL driver (`github.com/jackc/pgx/v5/stdlib`) is installed:

  ```bash
  go get github.com/jackc/pgx/v5/stdlib
  ```

- Modify the database connection details in the `.env` file according to your PostgreSQL setup.
