import type { Response } from "@/type/api"
import type { AuthResponse, SignInUserDTO } from '@/type/auth'
import { FetchWithOutAuth } from "@/api/api";

export const SignIn = async (signInProps: SignInUserDTO): Promise<Response<AuthResponse>> => {
  const res = await FetchWithOutAuth('api/v1/users/signin/', {
    method: 'POST',
    body: JSON.stringify({ ...signInProps }),
  })
  return {
    data: res.data,
    error: res.error,
  }
}