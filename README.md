###Objetivo
Desenvolver um rate limiter em Go que possa ser configurado para limitar o número máximo de requisições por segundo com base em um endereço IP específico ou em um token de acesso.

###Descrição
O objetivo deste desafio é criar um rate limiter em Go que possa ser utilizado para controlar o tráfego de requisições para um serviço web. O rate limiter deve ser capaz de limitar o número de requisições com base em dois critérios:

- **Endereço IP: O rate limiter deve restringir o número de requisições recebidas de um único endereço IP dentro de um intervalo de tempo definido.
- **Token de Acesso: O rate limiter deve também poderá limitar as requisições baseadas em um token de acesso único, permitindo diferentes limites de tempo de expiração para diferentes tokens. O Token deve ser informado no header no seguinte formato:
API_KEY: <TOKEN>
As configurações de limite do token de acesso devem se sobrepor as do IP. Ex: Se o limite por IP é de 10 req/s e a de um determinado token é de 100 req/s, o rate limiter deve utilizar as informações do token.

###Requisitos:
- O rate limiter deve poder trabalhar como um middleware que é injetado ao servidor web
- O rate limiter deve permitir a configuração do número máximo de requisições permitidas por segundo.
- O rate limiter deve ter ter a opção de escolher o tempo de bloqueio do IP ou do Token caso a quantidade de requisições tenha sido excedida.
As configurações de limite devem ser realizadas via variáveis de ambiente ou em um arquivo “.env” na pasta raiz.
- Deve ser possível configurar o rate limiter tanto para limitação por IP quanto por token de acesso.
- O sistema deve responder adequadamente quando o limite é excedido:
    - Código HTTP: 429
    - Mensagem: you have reached the maximum number of requests or actions allowed within a certain time frame
- Todas as informações de "limiter” devem ser armazenadas e consultadas de um banco de dados Redis. Você pode utilizar docker-compose para subir o Redis.
- Crie uma “strategy” que permita trocar facilmente o Redis por outro mecanismo de persistência.
- A lógica do limiter deve estar separada do middleware.

Exemplos:
- Limitação por IP: Suponha que o rate limiter esteja configurado para permitir no máximo 5 requisições por segundo por IP. Se o IP 192.168.1.1 enviar 6 requisições em um segundo, a sexta requisição deve ser bloqueada.
- Limitação por Token: Se um token abc123 tiver um limite configurado de 10 requisições por segundo e enviar 11 requisições nesse intervalo, a décima primeira deve ser bloqueada.
- Nos dois casos acima, as próximas requisições poderão ser realizadas somente quando o tempo total de expiração ocorrer. Ex: Se o tempo de expiração é de 5 minutos, determinado IP poderá realizar novas requisições somente após os 5 minutos.

###Dicas:
- Teste seu rate limiter sob diferentes condições de carga para garantir que ele funcione conforme esperado em situações de alto tráfego.

###Entrega:
- O código-fonte completo da implementação.
- Documentação explicando como o rate limiter funciona e como ele pode ser configurado.
- Testes automatizados demonstrando a eficácia e a robustez do rate limiter.
- Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
- O servidor web deve responder na porta 8080.

------ 

### Go
- Clone o repositório;
- Abra o terminal na pasta do projeto;
- Certifique-se de ter o redis rodando localmente em port `6379`;
- Executar `go run cmd/main.go`;

### Docker-compose
- Clone o repositório;
- Executar `docker-compose up --build`;

### Teste
- Certifique-se de que o redis esteja sendo executado localmente em port `6379`;
- Abra a pasta testsno terminal e execute `go test -v`;

### Executar limitador de taxa
- Configure parâmetros no arquivo `configs/config.env`;
- Você pode especificar a lista de tokens e seus limites;
- Você pode especificar o limite de IPs;
- Você pode especificar o tempo de bloqueio para ips e tokens, caso não seja especificado, o padrão é 1 minuto;
- Depois disso, você pode executar o aplicativo usando Go ou docker-compose
- O aplicativo possui 3 rotas de teste para validar o limitador de taxa:
  - `/token`: esta rota precisa de uma API_KEY no cabeçalho. O token precisa estar na lista de tokens configurada no arquivo `configs/config.env`. O limitador usará o token e o limite especificado no arquivo de configuração para o token específico para validar o acesso.
  - `/ip`:  esta rota não precisa de API_KEY no cabeçalho. O limitador utilizará o IP e o limite especificados no arquivo de configuração para validar o acesso.
  - `/both` ou `/`:  esta rota possui limitador configurado para bloquear tanto por ip quanto por token, seguindo as regras do desafio. Caso não seja enviado um token, ele utilizará o bloco por ip. Se um token for enviado, valida o limitador por token.
    - `TIP`: Você pode tentar executar essa rota sem token para validar o limite por ip e quando bloquear seu ip, você pode enviar um token válido, para validar o limite por token.
- Em qualquer rota, você pode verificar os logs para ver qual `origin` está sendo validado: token ou ip, e também a chave, que é o token específico ou o seu ip.
  - Quando o acesso for válido, você verá nos logs a contagem de acessos.
  - Quando o token ou o ip estiver bloqueado, você verá nos logs a mensagem `access is blocked`.
  - Quando o acesso for bloqueado, você receberá uma resposta com o código de `you have reached the maximum number of requests or actions allowed within a certain time frame`.
