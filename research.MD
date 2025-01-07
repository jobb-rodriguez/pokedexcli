# Testing
## Unit Testing
It is a form of software testing by which isolated source code is tested to validate expected behavior. (Wikipedia)

## Test-driven Development (TDD)
It is a way of writing code that involves writing an automated unit-level test case that fails, then writing just enough code to make the test pass, then refactoring both the test code and the production code, then repeating with anoter new test case.

> [!NOTE]
> Soon, refactor this project with behavior-driven development (BDD) and Gherkin.

## Black Box Testing vs White Box Testing
- Black-box testing is a type of testing in which it only focuses on the outer structure of the implementation that needs to be tested. 
-  White box testing techniques analyze the internal structures the used data structures, internal design, code structure, and the working of the software rather than just the functionality as in black box testing. 

### Types of Black Box Testing:
1. Functional Testing
2. Regression Testing
3. Nonfunctional Testing

> [!NOTE]
> (1) Equivalence partitioning, (2) boundary value analysis, and (3) error guessing are some methods to create test cases.

### Types of White Box Testing:
1. Unit Testing
2. Integration Testing
3. Regression Testing

> [!NOTE]
> (1) Control flow testing, (2) data flow testing, and (3) statement coverage testing are some methods. 

## Running Tests
```bash
# ./ is the root folder
go test ./
```
