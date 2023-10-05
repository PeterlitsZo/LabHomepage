import axios from "axios";

import { useAuthStore } from "../states/auth";
import useStore from "../hooks/useStore";

function getPath(url: string) {
  return process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + url;
}

function useGetTokenConfig() {
  const token = useStore(useAuthStore, state => state.token);

  return {
    headers: {
      Authorization: 'Bearer ' + token,
    }
  };
}

export function useGetWithToken<R>(prefix?: string) {
  const tokenConfig = useGetTokenConfig();
  return (url: string) => {
    return axios.get<R>(getPath((prefix || '') + url), tokenConfig);
  }
}

export function usePostWithToken<T>(prefix?: string) {
  const tokenConfig = useGetTokenConfig();
  return (url: string, req: T) => {
    return axios.post(getPath((prefix || '') + url), req, tokenConfig);
  }
}

export function useDeleteWithToken<T>(prefix?: string) {
  const tokenConfig = useGetTokenConfig();
  return (url: string) => {
    return axios.delete(getPath((prefix || '') + url), tokenConfig);
  }
}