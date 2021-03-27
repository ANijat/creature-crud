# creature-crud

### Usage

Schema:
For initial step init/01.sql should be executed on the database. 

Config:
Mentioned 6 environment variables need to be configured for database configurations.
DBUser
DBPassword
DBName
DBHost
DBPort
DBtype
You may also add export in front of each line
example:
```bash
$ export DBUser=admin
$ export DBHost=localhost
```

API:
/creature/ GET -> Get all of the creatures from db
/creature/:id GET -> Get a creature from db by id
/creature/:id PUT -> Update a creature by id
/creature/ POST -> Create new creature
