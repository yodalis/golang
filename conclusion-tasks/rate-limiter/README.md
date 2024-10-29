# Bem-vindo à Minha Aplicação!

Atualmente, tudo está sendo inicializado por meio do Docker Compose, tanto o redis quanto a app. Para começar, você precisará executar o seguinte comando:

```bash
docker-compose up
```

Atualmente estamos configurados para ter um limite de 5 requisições por IP e 2 por token, você pode testar nossa aplicação por meio desses dois exemplos:

Por IP:
```curl
curl --location 'http://localhost:8080/'
```

Pelo token:
```curl
curl --location 'http://localhost:8080/' \
--header 'API_KEY: abc123'
```

A resposta retornará uma saudação:

```json
Welcome!
```

E após o limite ser excedido, vamos receber a seguinte resposta:
```json
you have reached the maximum number of requests or actions allowed within a certain time frame
```

Qualquer dúvida, estou à disposição. Obrigada!
