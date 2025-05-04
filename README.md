# LAPI

LAPI é uma biblioteca em Go que fornece uma interface simples e intuitiva para fazer requisições HTTP. Ela foi projetada para facilitar a criação e gerenciamento de requisições HTTP com suporte a autenticação, headers, query parameters, body e tratamento de erros.

## Características

- Interface fluente para construção de requisições HTTP
- Suporte a diferentes métodos HTTP (GET, POST, PUT, DELETE, etc.)
- Gerenciamento de headers
- Suporte a query parameters
- Manipulação de body da requisição
- Tratamento de erros personalizado
- Suporte a autenticação
- Contexto para gerenciamento de requisições

## Requisitos

- Go 1.24.2 ou superior

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/pedrofreit4s/lapi.git
cd lapi
```

2. Instale as dependências:

```bash
go mod download
```

## Uso Básico

```go
// Criando uma requisição
r := lapi.NewRequest()

// Configurando a requisição
r.SetMethod("GET")
r.SetBaseURL("https://exemplo.com")
r.SetHeader("Content-Type", "application/json")
r.SetHeader("Authorization", "Bearer seu-token")

// Adicionando query parameters
r.SetQuery(map[string]string{
    "parametro": "valor",
})

// Enviando a requisição
resp, err := r.Send()
if err != nil {
    log.Fatal(err)
}

// Lendo a resposta
body, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(body))
```

## Estrutura do Projeto

```
lapi/
├── cmd/
│   └── examples/        # Exemplos de uso
├── internal/
│   └── lapi/           # Código fonte principal
│       ├── auth.go     # Gerenciamento de autenticação
│       ├── body.go     # Manipulação do body
│       ├── context.go  # Gerenciamento de contexto
│       ├── dest.go     # Configuração de destino
│       ├── error.go    # Tratamento de erros
│       ├── header.go   # Gerenciamento de headers
│       ├── http.go     # Configurações HTTP
│       ├── query.go    # Manipulação de query parameters
│       └── request.go  # Estrutura principal da requisição
├── main.go
├── go.mod
└── README.md
```

## Exemplos

Veja exemplos completos de uso no diretório `cmd/examples/`.

## Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -m 'Add nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
