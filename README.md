# Go Num 

`"Go Num"` - A number management application.

## Requirements

  * **Go**
  * **Terminal** for execution

## Installation

1.  Download or clone the project to your computer.
2.  Open the terminal or command prompt and navigate to the project directory.
3.  Run the program with the command `go run main.go`.

## Usage

> A mini-project developed as part of the leveling course in the Distributed Programming discipline, using the Go language, taught by Professor Ruan Delgado Gomes, at the Federal Institute of Education, Science and Technology of Paraíba (IFPB). [description](Miniprojeto_Go.pdf)

## Project Requirements

  * `Command Line:` An interface for the application to be executed.
  * `Slices:` Use slices to store numbers.
  * `Functions:` Code modularization.
  * `Multiple Returns:` Use functions with multiple returns to handle errors and statistics.

## How it Works

1.  **Objective:** Explore basic concepts of the Go language, such as variables, conditionals, loops, functions, and slices.
2.  **Tasks:** The application should present a menu to the user with the following options:
      * **1) Add number:** Prompts for an integer and adds it to a slice.
      * **2) List numbers:** Displays all stored numbers.
      * **3) Remove by index:** Removes a number from the slice based on a given index.
      * **4) Statistics:** Calculates the minimum, maximum, and average of the numbers.
      * **5) Safe division:** Performs division between two numbers with error handling.
      * **6) Clear list:** Empties the number slice.
      * **0) Exit:** Closes the application.
3.  **Technical Requirements:**
      * The program must be implemented in a single `main.go` file.
      * Error handling must follow the standard Go pattern (`if err != nil`).

## File Description

| Name | Description |
| :--- | :--- |
| `main.go` | The main file that contains all the logic for the number manager, including the options menu and data manipulation functions. |

### Bonus Features

  * Prevent adding negative numbers.
  * Implement ascending and descending list sorting.
  * An option to display only even numbers.
  * Export the list to a text file.

### Dev

  * Ananda Guedes