# running in local 
1. Create docker network
```
docker network create eld-network
```

2. Run database 
The database can be run using the this [repo].

3. Create your database 
```
docker exec -it database sh

psql -U postgres

create database your_database_name
```

4. Run service
First make sure service docker compose files is correctly configured. 
Thing that should be taken into account: 
> credentials
> ports (just changing outer port is enough the inner ports can remain the same)

Run the service with the following code
```
docker compose up -d --build
```

5. Stop service
```
docker compose down
```


# Needed doc's
- 
