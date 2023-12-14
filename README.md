run the server:

```
docker build . -t memleak
docker run -p 8080:8080 --rm memleak
```

visit http://localhost:8080/run to trigger a run.
and observe it's mem usage:

```
docker stats
```

![](img.png)