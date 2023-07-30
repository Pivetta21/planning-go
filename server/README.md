```bash
docker build -t planninggo .
```

```bash
docker run --restart unless-stopped --name planning_go -p 9000:9000 -d planninggo
```
