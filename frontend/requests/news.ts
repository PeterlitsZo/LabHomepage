import { useMutation, useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useContext } from 'react';
import { AuthContext } from '../contexts/auth';

export interface News {
  id: string;
  title: string;
  content: string;
}

interface ListNewsResponse {
  news: News[];
}

export const useListNews = () => {
  const query = useQuery({
    queryKey: ['news'],
    queryFn: () => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/news';
      return axios.get<ListNewsResponse>(path);
    },
  })

  return query;
}

export const useGetNews = (id: string) => {
  const query = useQuery({
    queryKey: ['news', id],
    queryFn: () => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/news/' + id;
      return axios.get<News>(path);
    },
  })

  return query;
}

/**
 * Get a `createNews` object. Use method `.mutate` to send request. See
 * TanStack/react-query's document to know more.
 */
export const useCreateNews = () => {
  type CreateNewsRequest = {
    title: string;
    content: string;
  };

  const auth = useContext(AuthContext);

  const mutation = useMutation({
    mutationFn: (req: CreateNewsRequest) => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/news';
      return axios.post(path, req, {
        headers: {
          Authorization: 'Bearer ' + auth.token,
        }
      });
    },
  })

  return mutation;
}

export const useDeleteNews = () => {
  const auth = useContext(AuthContext);
  const listNewsQuery = useListNews();

  const mutation = useMutation({
    mutationFn: (id: string) => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/news/' + id;
      return axios.delete(path, {
        headers: {
          Authorization: 'Bearer ' + auth.token,
        }
      });
    },
    onSuccess: () => {
      listNewsQuery.refetch();
    }
  })

  return mutation;
}