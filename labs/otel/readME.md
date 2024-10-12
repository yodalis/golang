# Bem-vindo à Minha Aplicação!

Atualmente, tudo está sendo inicializado por meio do Docker Compose. Para começar, você precisará executar o seguinte comando:

```bash
docker-compose up
```

Após a inicialização, você pode fazer uma chamada para obter a temperatura com o CEP, utilizando o corpo da requisição como no exemplo abaixo:

```json
{"cep": "02461011"}
```

A resposta retornará um JSON com a temperatura da sua localização, como este exemplo:

```json
{
    "temp_c": 20.1,
    "temp_f": 68.2,
    "temp_k": 293.1
}
```

Além disso, você pode acessar o site do Zipkin em [http://localhost:9411/zipkin](http://localhost:9411/zipkin) para visualizar os traces. Você encontrará o `serviceName` como `weather-service` ou `cep-service`, onde poderá verificar os tracers e spans que criei.

Qualquer dúvida, estou à disposição. Obrigada!
