import { useMutation, useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useContext } from 'react';
import { AuthContext } from '../contexts/auth';

export interface Paper {
  id: string;
  title: string;
  content: string;
}

interface ListPaperResponse {
  papers: Paper[];
}

export const useListPaper = () => {
  const query = useQuery({
    queryKey: ['papers'],
    queryFn: () => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/papers';
      return axios.get<ListPaperResponse>(path);
    },
  })

  return query;
}

export const useGetPaper = (id: string) => {
  const query = useQuery({
    queryKey: ['papers', id],
    queryFn: () => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/papers/' + id;
      return axios.get<Paper>(path);
    },
  })

  return query;
}

export const useCreatePaper = () => {
  type CreatePaperRequest = {
    title: string;
    content: string;
  };

  const auth = useContext(AuthContext);

  const mutation = useMutation({
    mutationFn: (req: CreatePaperRequest) => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/papers';
      return axios.post(path, req, {
        headers: {
          Authorization: 'Bearer ' + auth.token,
        }
      });
    },
  })

  return mutation;
}

export const useDeletePaper = () => {
  const auth = useContext(AuthContext);
  const listPaperQuery = useListPaper();

  const mutation = useMutation({
    mutationFn: (id: string) => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/papers/' + id;
      return axios.delete(path, {
        headers: {
          Authorization: 'Bearer ' + auth.token,
        }
      });
    },
    onSuccess: () => {
      listPaperQuery.refetch();
    }
  })

  return mutation;
}