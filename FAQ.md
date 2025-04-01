# FAQ

## Tools

1. What DBMS is chosen for this pet project ?
    - Postgres RDBMS  is chosen for simplicity on current stage.
2. What tool is chosen to run DB migrations and why ?
    - goose cli tool is chosen for simplicity on current stage.
3. Sqlx db driver (adapter) is chosen for simplicity on current stage.

## Flow

0. What happened if change name or email after participant was added to any meet ?

## GitHub issues

Tasks described below should be synced with issues:
https://github.com/users/ipavlov93/projects/1/views/1

### Tech Tasks

0. Add audit fields to models: createdAt, deletedAt, etc.
1. Add repository, service, facade, controller (handlers) layers.
2. Add unit tests for each of the layers.


### Business Logic Tasks