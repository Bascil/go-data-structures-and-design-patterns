# structs
- a struct is an object in go
- has the same similarities like classes in java
- can implement interfaces and declare methods
- structs prefer composition over inheritace
- composition over inheritance is a requirement of the go language

# pointer receivers
## why use pointers
- you can pass a person by value by removing the pointer signature
- a copy of the value of person is passed to the function
- it has limitations: any change in p wont be reflected in the source p
- if the person object is modified, the source person wont be affected