# max-services 

### Run app
From root run following commands
```
make proto
docker compose up --build
```

***Note:** db service might fail to connect to database on first compose, depending on the order of execution. If that happens stop it and run it again.*
