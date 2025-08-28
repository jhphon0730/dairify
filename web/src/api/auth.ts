import type { Response } from "@/type/api"
import type { AuthResponse, SignInUserDTO } from '@/type/auth'
import { FetchWithOutAuth } from "@/api/api";

const AUTH_SIGNIN_ENDPOINT = "api/v1/users/signin/";

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