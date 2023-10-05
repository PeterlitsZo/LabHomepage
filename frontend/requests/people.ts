import { useMutation, useQuery } from '@tanstack/react-query';
import { useDeleteWithToken, useGetWithToken, usePostWithToken } from './utils';

export interface Person {
  id: string;
  name: string;
  content: string;
}

interface ListPersonResponse {
  people: Person[];
}

export const useListPerson = () => {
  const fn = useGetWithToken<ListPersonResponse>('/api/v1/people');

  const query = useQuery({
    queryKey: ['people'],
    queryFn: () => fn(''),
  });

  return query;
}

export const useGetPerson = (id: string) => {
  const fn = useGetWithToken<Person>('api/v1/people/');

  const query = useQuery({
    queryKey: ['people', id],
    queryFn: () => fn(id),
  })

  return query;
}

export const useCreatePerson = () => {
  type CreatePersonRequest = {
    name: string;
    content: string;
  };

  const fn = usePostWithToken('/api/v1/people');

  const mutation = useMutation({
    mutationFn: (req: CreatePersonRequest) => fn('', req),
  })

  return mutation;
}

export const useDeletePerson = () => {
  const listPersonQuery = useListPerson();

  const fn = useDeleteWithToken('/api/v1/people/');

  const mutation = useMutation({
    mutationFn: (id: string) => fn(id),
    onSuccess: () => {
      listPersonQuery.refetch();
    }
  })

  return mutation;
}