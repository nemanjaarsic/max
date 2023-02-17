# max-services 

### Run app
From root run following commands
```
make proto
docker compose up --build
```

***Note:** db service might fail to connect to database on first compose, depending on the order of execution. If that happens stop it and run it again.*

### Use case
1. Send login request to get token. `host:8000/api/login` with body:
```
{ 
  "username": "username1",
  "password": "user1pass"
}
```
On db initialization 5 users are created, there data is in same format as the example above. [1-5]

2. Insert received token in your next request whether it is deposit or withdraw. `host:8000/api/deposit` with body:
```
{ 
  "id": "ab6b1cc7-d3e3-446c-bd87-24d742557e83",
  "token": "aa87dfasd9098sfasasd",
  "amount": 10,
  "timestamp": "2023-02-20 12:28:00"
}
```

## version 2

Adding new features on branch v2

* Added configuration loader. There is default configuration in config file for local run and if app is started via docker, variables will be loaded from dockerfile.
