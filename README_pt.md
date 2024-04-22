# Busca de Código Postal

Este é um projeto para recuperar informações de endereço com base em um código postal.

## Pré-requisitos

- Go 1.21.1 instalado
- Dependências do projeto instaladas

## Instalação

Para instalar o projeto, siga estes passos:

1. Clone este repositório e navegue até o diretório principal:
   ```bash
    $ git clone https://github.com/Ze-Victor/search-zip-code.git
    $ cd search-zip-code
   ```

2. Instale as dependências do projeto:

```bash
    $ go mod tidy
```

## Uso

1. Inicie o servidor:
```bash
    $ make
```

2. Faça solicitações para a API usando um cliente HTTP ou navegador da web. Certifique-se de gerar um token de autenticação antes de fazer a solicitação e incluí-lo no cabeçalho da requisição.

Para gerar o token, um requisição para a rota de `\auth` passando `email` e `senha`:

```bash
$ curl -X POST http://localhost:8001/api/v1/auth -H "Content-Type: application/json" -d '{"email": "your_email@test.com", "password": "your_password"}'

```

Faça a requisição para obter os dados da rua. Lembre de substituir `SEU_TOKEN_DE_AUTENTICACAO` pelo token JWT gerado:

```bash
$ curl -X GET http://localhost:8001/api/v1/cep/:cep_buscado -H "Content-Type: application/json" -H "Authorization: SEU_TOKEN_DE_AUTENTICACAO"
```


## Documentação

Documentação da API

A documentação da API está disponível no [Swagger UI](http://localhost:8001/api/v1/swagger/index.html) após iniciar o servidor.

## Testes Unitários

Para executar os testes do projeto, navegue até o diretório principal e use o comando:

```bash
    $ cd search-zip-code
    $ make test
```
## Teste de Ponta a Ponta

Os testes de ponta a ponta requerem a inicialização da API. Para executar esses testes, inicie o servidor, abra um terminal, navegue até o diretório _"test"_, e execute o seguinte comando:

```bash
    $ cd test
    $ go test -v
```

## Endpoint de saúde
Este endpoint pode ser acessado para verificar o estado de saúde do sistema. Retorna um status HTTP 200 se o sistema estiver saudável e um status HTTP 500 se o sistema estiver em um estado não saudável. Acesse:

- http://localhost:8001/api/v1/health


## Endpoint de métricas
Este endpoint fornece métricas sobre o sistema, como uso de CPU e memória. As métricas são retornadas em um formato específico, adequado para monitoramento e análise. Acesse:

- http://localhost:8001/api/v1/metrics


## Justificativa para escolha da tecnologia

Escolhi utilizar Go como minha principal tecnologia devido ao fato de ser amplamente adotada pela equipe para a qual estou buscando uma vaga, além da minha familiaridade com a linguagem. Ademais, Go oferece uma série de benefícios significativos, incluindo sua coerência e eficiência no tratamento de concorrência, compilação rápida, tipagem estática, simplicidade e clareza na sintaxe, bem como sua eficiência em tempo de execução.

Ainda, utilizei o framework gin-gonic para o roteamento. Além de já ter trabalhado com o mesmo, é um framework simples e direto ao ponto, adequado para projetos pequenos.

## Arquitetura do Projeto

Projetos em Go não possuem um padrão específico ou uma arquitetura bem definida, frequentemente deixando para a equipe estabelecer uma estrutura que melhor se adeque ao fluxo de desenvolvimento do projeto. Como se tratava de um projeto pequeno com um único desenvolvedor, busquei criar uma estrutura que fosse fácil de desenvolver e entender, seguindo alguns padrões do ecossistema de projetos em Go. 

Minha estrutura básica foi a seguinte:

- **Diretório Principal**: Arquivos de configuração do projeto são armazenados aqui.

- **cmd**: É onde estão localizadas as principais aplicações do projeto. Neste caso, temos apenas a API, que possui o arquivo principal main.go e o routes.go que possui as rotas da API.

- **config**: Arquivos de configuração global do projeto.

- **internal**: Código interno para a aplicação, não destinado a ser compartilhado com aplicações externas. 

- **docs**: Contém a documentação do projeto, gerada com gin-swagger.

- **test**: Os testes unitários estão no nível do arquivo que estão testando, mas criei um teste de ponta a ponta para verificar toda a funcionalidade da busca de endereço através do CEP. Ele garante que o componente do sistema funcione conforme esperado.

Com essa estrutura, busquei simplicidade, clareza, flexibilidade e a capacidade de separar claramente as responsabilidades dentro do projeto. Isso pode ajudar a facilitar o desenvolvimento, teste e manutenção do código.

## Estratégia para desenvolvimento

### Planejamento de Gerência de Configuração e Mudanças

- Criei a branch main, que será a branch final do projeto, e iniciei nela um projeto Go.

- Em seguida, criei o ambiente de desenvolvimento para receber novas funcionalidades.

- A partir da branch develop, foram criadas as features, que, após finalizadas e testadas, foram integradas ao ambiente de desenvolvimento e, por fim, à branch principal.

- Esse fluxo ajuda a manter a organização do desenvolvimento e facilita a divisão de tarefas.

### Calendário do Projeto

#### Dia 1:
- Validação da requisição e do CEP.
- Busca do CEP no arquivo simulado.
- Substituição dos números por zeros até encontrar o CEP.
- Revisão de código.
#### Dia 02:
- Testes unitários e testes de ponta a ponta.
- Estruturação do projeto para melhor se adequar ao desenvolvimento.
- Revisão de códigos.
#### Dia 03:
- Qualidade do código com análise de código estático (linter).
- Makefile para automatizar tarefas de compilação e execução.
- Documentação com Swagger.
- Estruturação de loggers.
#### Dia 04: 
- Autorização de API.
- Design de API.


_Observação: O arquivo do desafio mencionava receber e retornar JSON. No entanto, a tarefa era BUSCAR um endereço através de um CEP, então fazia sentido ser uma requisição GET. No entanto, os Getters não levam um corpo na requisição. Portanto, executei a busca enviando o CEP como um parâmetro.
Revisão de código.
Revisão de documentação._

### Dia 05:

- Endpoint para saúde da aplicação
- Endpoint para métricas da aplicação
- Reescrevendo teste de ponta a ponta
- Reescrita do README
- Revisão de código

### Dia 06:

- Conexão com banco de dados (SQLite)
- Refatoração de código
- Revisão de código
- Revisão de testes
- Revisão de documentação

## Resposta da questão 02:

Quando digitamos um endereço (http://www.netshoes.com.br) no navegador, ele começa um processo complexo. Primeiro ele precisa saber onde aquele endereço tá hospedado, ou seja, qual o domínio dele. Para isso, ele pede ajuda ao DNS, que funciona como uma lista telefônica, ou seja, um servidor de domínios. Esse servidor de domínios resolve o endereço mandado pelo servidor e retorna  um IP que será usado pelo navegador para fazer a comunicação cliente-servidor.

De posse do IP, o navegador e o servidor estabelecem uma conexão. A partir disso, o navegador (cliente) faz um pedido ao servidor através de uma requisição (HTTP Request) e o servidor, ao receber essa esse pedido, processa as informações e devolve para o cliente uma resposta (HTTP Response). Seguindo um fluxo, seria assim:

- Resolução de DNS.
- Estabelecimento da conexão TCP/IP.
- Envio do HTTP Request.
- Processamento da requisição pelo servidor.
- Envio do HTTP Response.
