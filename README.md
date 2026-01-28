# Go Calculator

A simple **command-line calculator** written in Go.

## Description

This program performs basic arithmetic operations with **two integers**.  
It supports the following operators:

- `+` Addition  
- `-` Subtraction  
- `*` Multiplication  
- `/` Division  
- `%` Modulus  

The calculator **handles errors** such as:

- Incorrect number of arguments  
- Invalid operator  
- Division or modulus by zero  
- Non-integer input  

---

## Usage

Run the program with **exactly 3 arguments**


## Imported the following
package - main defines this as an executable program.
os - is used to read command-line arguments.
fmt - is used to print output.
strconv - is used to convert string arguments to integers.

## Check Number of Arguments
The program expects exactly 3 arguments:
number1 operator number2
os.Args always includes the program name as the first element, so we check for len(os.Args) != 4.
Prints Error if the count is incorrect.

## Convert Arguments to Integers
strconv.Atoi converts strings to integers.
os.Args[1] → first number, os.Args[3] → second number
If conversion fails (e.g., "five"), it prints Error.

## Perform the Operation
Uses a switch to choose the operation based on the operator string.
Handles division and modulus by zero to prevent runtime errors.
If the operator is not recognized, prints Error.

## Output
Prints the result of the calculation to the console.
All error cases print Error and stop execution.
