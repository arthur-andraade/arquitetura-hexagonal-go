This repository is where I have saved the code that I developed during the 'Hexagonal Architecture' course from FullCycle. To make the project I used the Golang.

## Hexagonal architecture
Hexagonal architecture is a software architecture pattern that aims to isolate business logic from the external environment, promoting testability and flexibility. It consists of three main elements: adapters, ports, and the application core.

1. Application Core:
   - Contains business rules and the main logic of the application.
   - Does not depend on external implementation details.

2. Ports:
   - Defines interfaces for communication between the application core and the external environment.
   - May include interfaces for input (e.g., user interfaces, APIs) and output (e.g., database persistence).
   - Primary ports is where external environment make communication with application core
   - Secondary ports is where the application core make communication with external enviroment

3. Adapters:
   - Implement the interfaces defined in the ports.
   - Connect the application core to the external environment.
   - May include database adapters, API adapters, user interface adapters, etc.
