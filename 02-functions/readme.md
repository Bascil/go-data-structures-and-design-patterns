## functions
- named section of a program to perform a specific task
- portion of code that sorrounds some action you want to perform
- returns one or more values or nothing
- main tool for developers to maintain structure and code readability
- allows an experienced programmer to write proper unit tests
- a function must do just one thing but it must do it damn well
- a function block is a piece of code with its own variables and flow

## stack memory
- Is a region of memory used to store local variables and function calls.
- When a function is called, space for its local variables is allocated on the stack.
- This space is deallocated when the function returns.
- Allows functions to have their own private data not accessible outside the function
- Go uses a stack for managing function calls and local variables.
- The stack memory is managed by the program's runtime environment 
- The Go runtime automatically handles the management of stack and heap memory

## call stack and stack frame
- The call stack is a data structure that keeps track of all the active functions or methods in a program.
- The call stack is responsible for keeping track of the local variables in your function
- The call stack follows the Last-In, First-Out (LIFO) principle, which means that the last function called is the first one to be completed.
- A stack frame, also known as an activation record or activation frame, is a data structure within the call stack that represents a single function call.
- Each function call creates a new "stack frame" on top of the stack.
- It contains all the information needed for a specific function call to execute, including local variables, parameters, return addresses, and sometimes other bookkeeping information.
- When a function is called, a new stack frame is created and pushed onto the call stack. When the function completes, its stack frame is popped off the call stack.
