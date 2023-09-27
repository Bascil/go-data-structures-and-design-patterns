# prototype design pattern
- is a `creational` design pattern
- some patterns overlap a little bit but have small differences
- aims to have object(s) created at compilation time
- main objective is to avoid repititive object creation
- frees cpu of complex object initialization that takes more memory resources
- powerful way to build caches and default objects

## prototype vs singleton

* Prototype - A new object is created each time it is injected/looked up.
* Singleton - The same object is returned each time it is injected/looked up. Here it will instantiate one instance of a struct and then return it each time.

## imaginary shirt shop
- have a shirt cloner objectand interfce to ask for diffent shirt types
- each time we need a new shirt, we will take this prototype and clone it
- SKU will identify items stored in a specific location
- SKU of created object should not affect new object creation
- An info method must give all available information on the instance fields including updated SKU