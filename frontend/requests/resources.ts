import { useMutation, useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useContext } from 'react';
import { AuthContext } from '../contexts/auth';

export interface Resource {
  id: string;
  title: string;
  content: string;
}

interface ListResourceResponse {
  resources: Resource[];
}

export const useListResource = () => {
  const query = useQuery({
    queryKey: ['resources'],
    queryFn: () => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/resources';
      return axios.get<ListResourceResponse>(path);
    },
  })

  return query;
}

export const useGetResource = (id: string) => {
  const query = useQuery({
    queryKey: ['resources', id],
    queryFn: () => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/resources/' + id;
      return axios.get<Resource>(path);
    },
  })

  return query;
}

export const useCreateResource = () => {
  type CreateResourceRequest = {
    title: string;
    content: string;
  };

  const auth = useContext(AuthContext);

  const mutation = useMutation({
    mutationFn: (req: CreateResourceRequest) => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/resources';
      return axios.post(path, req, {
        headers: {
          Authorization: 'Bearer ' + auth.token,
        }
      });
    },
  })

  return mutation;
}

export const useDeleteResource = () => {
  const auth = useContext(AuthContext);
  const listResourceQuery = useListResource();

  const mutation = useMutation({
    mutationFn: (id: string) => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/resources/' + id;
      return axios.delete(path, {
        headers: {
          Authorization: 'Bearer ' + auth.token,
        }
      });
    },
    onSuccess: () => {
      listResourceQuery.refetch();
    }
  })

  return mutation;
}