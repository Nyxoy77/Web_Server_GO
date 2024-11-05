# Web Server in Go

A basic backend for a Movies Management System implemented in Go (Golang). This project provides RESTful APIs to manage a collection of movies, supporting operations such as adding, updating, retrieving, and deleting movie records.

## Features
- **Add Movies**: Create and store movie entries with details like title, director, and release year.
- **Retrieve Movies**: Fetch all movies or search for specific titles.
- **Update & Delete**: Modify or remove movie records from the system.

## Prerequisites
- Go (Golang) installed on your system.
- Basic understanding of REST API concepts.

## Setup
1. **Clone the repository**:
   ```bash
   git clone https://github.com/Nyxoy77/Web_Server_GO.git
   cd Web_Server_GO
   ```

2. **Install dependencies**:
   Ensure you have `go mod` initialized. If not, run:
   ```bash
   go mod init
   ```
   Then, use:
   ```bash
   go mod tidy
   ```

3. **Run the server**:
   ```bash
   go run main.go
   ```

## Usage
- The server runs on a specified port (defined in the code). You can make API requests using tools like **Postman**, **Thunder Client**, or via **curl** in the terminal.
- Example endpoints:
  - `GET /movies`: Retrieve all movies.
  - `POST /movies`: Add a new movie.
  - `PUT /movies/{id}`: Update an existing movie.
  - `DELETE /movies/{id}`: Remove a movie.

## Project Structure
- **main.go**: Contains the server logic and route handling.
- **go.mod & go.sum**: Manage dependencies.

## Future Improvements
- Integrate database support.
- Implement user authentication.
- Add more sophisticated error handling.

## License
This project is open-source and available under the MIT License.

