// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginReply struct {
	Token string `json:"token"`
}
