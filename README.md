run the server:

```
docker build . -t memleak
docker run -p 8080:8080 --rm memleak
```

observe it's mem usage:

```
docker stats
```

![](img.png)