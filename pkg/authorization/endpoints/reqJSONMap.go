package endpoints

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Err   string `json:"err:omitempty"`
}

type LogoutRequest struct {
	Account string `json:"account"`
	Token   string `json:"token"`
}

type LogoutResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err:omitempty"`
}

type ServiceStatusRequest struct {
}

type ServiceStatusResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err,omitempty"`
}
