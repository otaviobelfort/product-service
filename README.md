# Resumo do Projeto Product Service

## O que foi aprendido

### Configuração do Ambiente
- **Docker Compose**: Utilização do `docker-compose.yml` para configurar e gerenciar um container PostgreSQL.
- **Configuração do Banco de Dados**: Criação de tabelas no PostgreSQL utilizando GORM e SQL.

### Desenvolvimento com Go
- **GORM**: Uso do ORM GORM para interagir com o banco de dados PostgreSQL.
- **Estrutura de Projeto**: Organização do projeto em pacotes, incluindo `configs`, `database`, e `handlers`.
- **Chi Router**: Configuração de rotas HTTP utilizando o framework Chi.
- **JWT Authentication**: Implementação de autenticação JWT com `go-chi/jwtauth`.

### Documentação
- **Swagger**: Integração do Swagger para documentação da API, incluindo a configuração de parâmetros principais como `@title`, `@version`, e `@description`.

### Testes de API
- **HTTP Requests**: Criação de requisições HTTP para testar endpoints da API, incluindo métodos GET e POST.

### Boas Práticas
- **Gerenciamento de Conexões**: GORM gerencia automaticamente as conexões com o banco de dados, eliminando a necessidade de fechar manualmente as conexões.
- **Migrações Automáticas**: Uso de `db.AutoMigrate` para gerenciar migrações de banco de dados automaticamente.



## Base de estudos para o projeto
Claro, aqui está a lista de arquivos convertida em texto:

1. Entendendo a primeira linha
2. Declaração e atribuição
3. Criação de tipos
4. Importando fmt e tipagem
5. Percorrendo Arrays, For, Condicionais
6. Slices
7. Maps
8. Funções
9. Funções variádicas
10. Closures
11. Iniciando com Structs
12. Composição de Structs
13. Métodos em Structs
14. Interfaces
15. Ponteiros
16. Quando usar ponteiros
17. Ponteiros e Structs
18. Interfaces vazias
19. Type assertion
20. Generics
21. Pacotes e módulos parte 1
22. Manipulação de arquivos e dados(JSON,...)
23. Chamadas e servidor HTTP, HttpClient, Context
24. Headers
25. ServeMux
26. Banco de dados, CRUD, Relacionamentos
27. API REST
### Multithreading
1. Introdução
2. Mutex. Concorrência vs Paralelismo vs Go
3. Multithreading
4. Scheduler
5. Go e suas green threads
6. Go Routines
7. Wait Groups 
8. Mutex e Operações Atômicas
9. Channel e Buffer
10. Forever
11. With WaitGroup
12. Channel Directions
13. Load Balancer

Referência: 
- [Aprenda Go com Testes](https://larien.gitbook.io/aprenda-go-com-testes/)
- [Go expert](https://curso.fullcycle.com.br/)

