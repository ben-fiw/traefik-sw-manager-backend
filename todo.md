# TODOs

## Core Features

### Service connection

-   [ ] Implement service connection logging
    -   [x] Create service connection when a service request is made
    -   [ ] Add a connection ping route
    -   [ ] Add active connection check to report route
    -   [ ] Add connection close route
    -   [ ] Handle connection states

### Cron

-   [ ] Add cron logics
    -   [ ] Add cron entrypoint
    -   [ ] Add cron job to check for inactive connections
    -   [ ] Add cron job to check for inactive services

### Load balancing

-   [ ] Add load balancing logics
    -   [x] Actively support multiple service instances
    -   [x] Consider open connections and resource rating when selecting a service
    -   [ ] Add statistics route for an external dashboard

## Documentation

-   [ ] Document the local development setup
-   [ ] Document the service API

## Active discovery

-   [ ] Add an optionlal active discovery method
    -   [ ] Scan for services in provided IP ranges
    -   [ ] Add an identification route module for services

## Out of scope

-   [ ] Create a module to handle service-side connection management
    -   [ ] Extend the module to auto-adjust the resource rating

## Bugs

**Currently, no known bugs**

# Change log

| Author                                                                | Date       | Changes         |
| --------------------------------------------------------------------- | ---------- | --------------- |
| Ben Fischer [\<root@ben-fischer.dev\>](mailto://root@ben-fischer.dev) | 2021-01-01 | Initial version |
