## Auth Manager

This is a demo project to see how open telemetry works by passing the context from the controller down to the repository level.
I also tried to use this to implement clean code where each domain are self contained.

In the user domain, there is the:
* `interfaces` that defines exported services.
* `mocks` a handwritten mock to simulate both happy path and error path.
* `entities` contains the DTO. This should be totally different from the views or database entities.
* `repository` houses the different datasources
* `repository/database` houses the database manipulation resources.
* `service` houses the implementation of the exported interfaces and serves as the gateway for the domain.
