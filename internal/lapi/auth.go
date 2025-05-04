package lapi

// SetAccessToken define o token de acesso JWT para a requisição.
// O token será automaticamente adicionado como Bearer token no header Authorization.
//
// Parâmetros:
//   - token: Token JWT de acesso
//
// Exemplo:
//
//	m.SetAccessToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
func (m *model) SetAccessToken(token string) {
	m.Auth.Token = token
}

// SetRefreshToken define o token de atualização JWT para a requisição.
// Este token pode ser usado para renovar o token de acesso quando expirado.
//
// Parâmetros:
//   - refreshToken: Token JWT de atualização
//
// Exemplo:
//
//	m.SetRefreshToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
func (m *model) SetRefreshToken(refreshToken string) {
	m.Auth.RefreshToken = refreshToken
}

// SetAuth define o token de acesso e o token de atualização JWT para a requisição.
// Esta função é um atalho para definir ambos os tokens de uma vez.
//
// Parâmetros:
//   - token: Token JWT de acesso
//   - refreshToken: Token JWT de atualização
//
// Exemplo:
//
//	m.SetAuth(
//	    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
//	    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
//	)
func (m *model) SetAuth(token, refreshToken string) {
	m.Auth.Token = token
	m.Auth.RefreshToken = refreshToken
}

// RevalidateToken revalida o token de acesso usando o token de atualização.
// Esta função é chamada automaticamente quando o token de acesso expira.
//
// Retorna:
//   - string: Novo token de acesso JWT
//
// Exemplo:
//
//	newToken := m.RevalidateToken()
func (m *model) RevalidateToken() string {
	// TODO: Implement the revalidation of the token
	return m.Auth.Token
}
