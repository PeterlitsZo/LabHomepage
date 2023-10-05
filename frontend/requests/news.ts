import { useMutation, useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useDeleteWithToken, useGetWithToken, usePostWithToken } from './utils';

export interface News {
  id: string;
  title: string;
  content: string;
}

interface ListNewsResponse {
  news: News[];
}

export const useListNews = () => {
  const fn = useGetWithToken<ListNewsResponse>('/api/v1/news');

  const query = useQuery({
    queryKey: ['news'],
    queryFn: () => fn(''),
  })

  return query;
}

export const useGetNews = (id: string) => {
  const fn = useGetWithToken<News>('/api/v1/news/');

  const query = useQuery({
    queryKey: ['news', id],
    queryFn: () => fn(id),
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

  const fn = usePostWithToken('/api/v1/news');

  const mutation = useMutation({
    mutationFn: (req: CreateNewsRequest) => fn('', req),
  })

  return mutation;
}

export const useDeleteNews = () => {
  const listNewsQuery = useListNews();

  const fn = useDeleteWithToken('/api/v1/news/');

  const mutation = useMutation({
    mutationFn: (id: string) => fn(id),
    onSuccess: () => {
      listNewsQuery.refetch();
    }
  })

  return mutation;
}