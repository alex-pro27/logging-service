### LOGGING SERVICE APP ON YANDEX CLICKHOUSE

#### RUN:

```
cp ./example.env ./.env
docker-compose up --build
```

#### Test ping:
<p>
send to GET request http://localhost:8080/api/ping
</p>

<hr/>

#### Example of adding logs:

<p>
send to PUT request http://localhost:8080/api/add-log, 
<br/>content-type: application/json
</p>

```
{
    "text": 404 GET / (127.0.0.1) 457.09ms",
    "app_name": "test",
    "type": "ERROR"
}
```

#### Get log:

<p>
send to GET request http://localhost:8080/api/logs/?app=test&type=DEBUG
</p>

