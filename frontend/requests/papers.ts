import { useMutation, useQuery } from '@tanstack/react-query';
import { useDeleteWithToken, useGetWithToken, usePostWithToken } from './utils';

export interface Paper {
  id: string;
  title: string;
  content: string;
}

interface ListPaperResponse {
  papers: Paper[];
}

export const useListPaper = () => {
  const fn = useGetWithToken<ListPaperResponse>('/api/v1/papers');

  const query = useQuery({
    queryKey: ['papers'],
    queryFn: () => fn(''),
  })

  return query;
}

export const useGetPaper = (id: string) => {
  const fn = useGetWithToken<Paper>('/api/v1/papers/');

  const query = useQuery({
    queryKey: ['papers', id],
    queryFn: () => fn(id),
  })

  return query;
}

export const useCreatePaper = () => {
  type CreatePaperRequest = {
    title: string;
    content: string;
  };

  const fn = usePostWithToken('/api/v1/papers');

  const mutation = useMutation({
    mutationFn: (req: CreatePaperRequest) => fn('', req),
  })

  return mutation;
}

export const useDeletePaper = () => {
  const listPaperQuery = useListPaper();

  const fn = useDeleteWithToken('/api/v1/papers/');

  const mutation = useMutation({
    mutationFn: (id: string) => fn(id),
    onSuccess: () => {
      listPaperQuery.refetch();
    }
  })

  return mutation;
}