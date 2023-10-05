import { useMutation, useQuery } from '@tanstack/react-query';
import { useDeleteWithToken, useGetWithToken, usePostWithToken } from './utils';

export interface Resource {
  id: string;
  title: string;
  content: string;
}

interface ListResourceResponse {
  resources: Resource[];
}

export const useListResource = () => {
  const fn = useGetWithToken<ListResourceResponse>();

  const query = useQuery({
    queryKey: ['resources'],
    queryFn: () => {
      return fn('/api/v1/resources');
    },
  })

  return query;
}

export const useGetResource = (id: string) => {
  const fn = useGetWithToken<Resource>('/api/v1/resources/');

  const query = useQuery({
    queryKey: ['resources', id],
    queryFn: () => fn(id),
  })

  return query;
}

export const useCreateResource = () => {
  type CreateResourceRequest = {
    title: string;
    content: string;
  };

  const fn = usePostWithToken('/api/v1/resources');

  const mutation = useMutation({
    mutationFn: (req: CreateResourceRequest) => fn('', req),
  })

  return mutation;
}

export const useDeleteResource = () => {
  const listResourceQuery = useListResource();

  const fn = useDeleteWithToken('/api/v1/resources/');

  const mutation = useMutation({
    mutationFn: (id: string) => fn(id),
    onSuccess: () => {
      listResourceQuery.refetch();
    }
  })

  return mutation;
}