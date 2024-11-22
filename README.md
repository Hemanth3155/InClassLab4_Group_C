 InClassLab4_Group_C

 Calculator Project

 Project Overview

This project implements a simple calculator module in Go. The module contains functions for basic arithmetic operations such as division and squaring a number. It includes unit tests and benchmarks to ensure correct functionality and performance.

 Instructions

 Running the Program

To use the calculator functions, you need to import the `calculator` package into your Go program.

Example usage:

go
package main

import (
    "fmt"
    "calculator"
)

func main() {
    // Example of Divide function
    result, err := calculator.Divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Divide result:", result)
    }

    // Example of Square function
    fmt.Println("Square result:", calculator.Square(4))
}

 Running Tests

To run the tests, use the following command:

go test

 Running Benchmarks

To run the benchmarks, use the following command:

go test -bench=.

 Expected Output of the Tests

The tests should pass successfully. The expected output for each test is as follows:

 `TestDivide`: Verifies correct division and error handling for division by zero.
 `TestSquare`: Verifies correct squaring of integers.

 Explanation of Test and Benchmark Design

 Test Design: Tests are designed to cover both normal and edge cases for the `Divide` and `Square` functions. For example, dividing by zero should result in an error, and squaring negative numbers should return a positive result.
 Benchmark Design: Benchmarks are designed to measure the performance of the `Divide` and `Square` functions with representative inputs.

 How to Execute

1. To run the program:
   - Import the `calculator` package and use the functions as shown in the example.

2. To run the tests:
    Navigate to your project folder and run the following command:
     
     go test

3. To run the benchmarks:
    Run the following command:
     
     go test -bench .

 Outputs:
 
Example for test output:
     PASS
     ok  	calculator	0.001s

Example benchmark output:
    BenchmarkDivide-8   	1000000000	         1.0 ns/op
    BenchmarkSquare-8   	1000000000	         0.5 ns/op





