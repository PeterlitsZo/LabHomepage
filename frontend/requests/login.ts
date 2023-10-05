import { useMutation } from '@tanstack/react-query';
import axios from 'axios';

// Send a HTTP request to login.
export const useLogin = () => {
  type LoginRequest = {
    username: string;
    password: string;
  }
  type LoginData = {
    token: string;
  }

  const mutation = useMutation({
    mutationFn: (req: LoginRequest) => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/login';
      return axios.post<LoginData>(path, req);
    },
  })

  return mutation;
}