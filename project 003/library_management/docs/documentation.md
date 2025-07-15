# Console-Based Library Management System

## Objective
This system provides a basic console interface for managing books and members in a library. It demonstrates fundamental Go concepts such as structs, interfaces, methods, slices, and maps.

## Features
- **Add Book:** Allows adding new books to the library with a title, author, and an initial status of "Available".
- **Remove Book:** Enables removal of books by their unique ID, provided the book is not currently borrowed.
- **Borrow Book:** Facilitates borrowing of available books by registered members. Updates book status and member's borrowed list.
- **Return Book:** Allows members to return borrowed books, making them "Available" again and removing them from the member's borrowed list.
- **List Available Books:** Displays all books currently marked as "Available" in the library.
- **List Borrowed Books:** Shows all books borrowed by a specific member.
- **Add Member (Helper):** A utility to quickly add new members for testing the system.

## Folder Structure
The project adheres to the specified clean architecture for better organization and separation of concerns:
- `main.go`: The application's entry point, handling the main menu and user interaction loop.
- `controllers/`: Contains `library_controller.go` which acts as the intermediary between the user interface (console) and the business logic (services). It parses user input and formats output.
- `models/`: Defines the data structures (`Book` and `Member` structs) used throughout the application.
- `services/`: Houses `library_service.go`, which implements the `LibraryManager` interface. This is where all the core business logic and data manipulation (adding, removing, borrowing, returning books; managing members) resides.
- `docs/`: Stores this documentation file.
- `go.mod`: Manages Go module dependencies.

## How to Run

1.  **Navigate to the project directory:**
    ```bash
    cd library_management
    ```
2.  **Run the application:**
    ```bash
    go run main.go
    ```

## Error Handling
The system includes basic error handling for common scenarios:
- Attempting to remove a borrowed book.
- Trying to borrow an already borrowed or non-existent book.
- Trying to return a book that isn't borrowed or doesn't exist.
- Invalid member or book IDs.

## Future Enhancements
- Persistence: Save library data to a file (CSV, JSON, or a simple database) so it's not lost when the application closes.
- More robust input validation.
- Advanced search functionalities.
- User authentication/authorization.
- A more sophisticated UI (e.g., a web interface).