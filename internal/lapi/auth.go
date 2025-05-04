package lapi

// SetAccessToken define o token de acesso para a requisição.
func (m *model) SetAccessToken(token string) {
	m.Auth.Token = token
}

// SetRefreshToken define o token de atualização para a requisição.
func (m *model) SetRefreshToken(refreshToken string) {
	m.Auth.RefreshToken = refreshToken
}

// SetAuth define o token de acesso e o token de atualização para a requisição.
func (m *model) SetAuth(token, refreshToken string) {
	m.Auth.Token = token
	m.Auth.RefreshToken = refreshToken
}

// RevalidateToken revalida o token de acesso.
func (m *model) RevalidateToken() string {
	// TODO: Implement the revalidation of the token
	return m.Auth.Token
}
