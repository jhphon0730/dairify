export interface User {

}

export interface SignInUserDTO {
  username: string
  password: string
}

export interface AuthResponse {
  access_token: string
  user: User
}

export interface SignUpUserDTO {
  username: string
  nickname: string
  email: string
  password: string
}
