export interface User {

}

export interface SignUpUserDTO {
  name: string
  email: string
  password: string
}

export interface SignInUserDTO {
  username: string
  password: string
}

export interface AuthResponse {
  access_token: string
  user: User
}