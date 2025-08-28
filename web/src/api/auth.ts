import type { Response } from "@/type/api"
import type { AuthResponse, SignInUserDTO, SignUpUserDTO } from '@/type/auth'
import { FetchWithOutAuth } from "@/api/api";

const AUTH_SIGNIN_ENDPOINT = "api/v1/users/signin/";
const AUTH_SIGNUP_ENDPOINT = "api/v1/users/signup/";

export const SignIn = async (signInProps: SignInUserDTO): Promise<Response<AuthResponse>> => {
  const res = await FetchWithOutAuth(AUTH_SIGNIN_ENDPOINT, {
    method: 'POST',
    body: JSON.stringify({ ...signInProps }),
  })
  return {
    data: res.data,
    error: res.error,
  }
}

export const SignUp = async (signUpProps: SignUpUserDTO): Promise<Response<AuthResponse>> => {
  const res = await FetchWithOutAuth(AUTH_SIGNUP_ENDPOINT, {
    method: 'POST',
    body: JSON.stringify({ ...signUpProps }),
  })
  return {
    data: res.data,
    error: res.error,
  }
}