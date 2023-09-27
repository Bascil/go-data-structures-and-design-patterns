# stack
- stacks and queues are very basic data structures
- the difference is basically how data is removed
- stacks have LIFO data structure
- stacks have two main operations push, pop and peek
- peek returns the top element from the stack without removing it
- every programming language, compiler and OS uses stacks internally
- queues have FIFO data structure
- queues have two main operations enqueue & dequeue
- enqueue pushes/adds an item to the queue
- dequeu takes out the oldest item first(FIFO)
- stacks & queues can be implemented using arrays, linked list, ring buffer, e.t.c
- each approach has its own advantages & tradeoffs interms of performance & memory usage
- the code will implement a basic queue and stack.
- for more performance critical scenarios consider using a linked list or ring buffer

## ring buffer
A ring buffer, also known as a circular buffer or cyclic buffer, is a fixed-size data structure that efficiently manages a collection of items. It is designed for scenarios where you need to efficiently read and write data in a circular manner, such as in streaming applications

Ring buffers are a versatile data structure that find applications in various domains, particularly in embedded systems, real-time systems, multimedia processing, and network protocols where efficient handling of continuous streams of data is crucial.

## stack applications
- function calls - keep track of function calls and returns
- undo operation - implement undo feature in a program

## queue applications
- event handling - used in event driven programming 
- message queues in IPC 
- used in async communication between processes(message passing) e.g kafka, pubsub, rabbitmq
- call centres - to manage incoming calls
- process scheduling- to determine the order in which processes are executed
- printer queue - to manage print jobs
- buffering in networking - to manage data packets waiting to be transmitted or received


```
    |_____|
    |_____|         
    |_____|
    |_____|        
    stack

    |     |      |     |
    |_____|______|_____|
    queue

```
### Memory management
## stack memory
- Is a region of memory used to store local variables and function calls.
- When a function is called, space for its local variables is allocated on the stack.
- This space/memory is deallocated when the function returns.
- Allows functions to have their own private data not accessible outside the function (local variables)
- Go uses a stack for managing function calls and local variables.
- Each goroutine (concurrent thread in Go) has its own stack.
- The size of a stack is typically fixed and relatively small compared to the heap
- The stack memory is managed by the program's runtime environment 
- The Go runtime automatically handles the management of memory on both the stack and the heap memory,
- This make memory management more convenient for the developer compared to languages like C or C++.

## advantages
- automatic allocation and deallovation of memory
- makes it easier for the programmer
- faster than heap or dynamic memory

## disadvantages
- lack of control; sometimes you need to directly manage memory e.g when reading data from a file, you dont know how big the file will be. memory may grow based on size of the file. 

## call stack and stack frame
- The call stack is a data structure that keeps track of all the active functions or methods in a program.
- The call stack is responsible for keeping track of the local variables in your function
- The call stack follows the Last-In, First-Out (LIFO) principle, which means that the last function called is the first one to be completed.
- A stack frame, also known as an activation record or activation frame, is a data structure within the call stack that represents a single function call.
- Each function call creates a new "stack frame" on top of the stack.
- It contains all the information needed for a specific function call to execute, including local variables, parameters, return addresses, and sometimes other bookkeeping information.
- When a function is called, a new stack frame is created and pushed onto the call stack. When the function completes, its stack frame is popped off the call stack.

## heap memory
- Heap memory is a region of a computer's memory (RAM) that is used for dynamic memory allocation.
- this means that memory can be allocated and deallocated at any point during a program's execution.
- programmer has direct control of memory
- It is a data structure allows you to store items in memory in any order.
- Information is fragmented resulting to slower access time
- Adding items to the heap has a higher overhead than adding items to the stack
- In some programming languages like C and C++, developers are responsible for manually managing heap memory 
- This includes allocating memory using functions like malloc, and deallocating it using free.
- Heap memory is unstructured, meaning that memory blocks can be allocated in any order 
- Heap memory is typically larger in size but slower compared to stack memory. 
- programmer must allocate and deallocate memory
- It is used for storing larger data structures and objects.
- If not managed properly, heap memory can lead to memory leaks,
- A memory leak is where memory that is no longer in use is not properly deallocated.
- Accessing data in heap memory can be slower compared to stack memory
- This is because it involves accessing a location in RAM rather than using CPU registers.

### heap memory in go
- The heap is where objects like structs, slices, and maps are typically allocated.
- Go also uses a heap for dynamic memory allocation, including objects with unknown or changing size.
- Memory in the heap is managed by Go's garbage collector
- The garbage collector automatically reclaims memory that is no longer in use.
- The Go runtime automatically handles the management of memory on both the stack and the heap,

### heap summary
Heap memory is crucial for managing data structures and objects that have a dynamic or unpredictable size. It provides flexibility in memory allocation, allowing programs to adapt to different runtime conditions and handle varying amounts of data. However, it also requires careful management to avoid issues

## garbage collection
In most runtimes, there is a garbage collector. It goes around and looks at the heap and reclaims memory that is not in use(a bunch  of data dangling around) and cleans them up (goes  around collecting garbage). 

Some modern programming languages, like Go, Java and C#, use automatic garbage collection to manage heap memory. Garbage collection automatically reclaims memory that is no longer in use, reducing the risk of memory leaks.

## asynchronous methods
whenever you run anything asynchronously, its going to run on a different thread. When you are running your code it has multiple threads. Each thread has its own call stack. To access the result of your async method, the result needs to be stored on the heap.
