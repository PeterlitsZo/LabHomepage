import { useMutation, useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useContext } from 'react';
import { AuthContext } from '../contexts/auth';

export interface Person {
  id: string;
  name: string;
  content: string;
}

interface ListPersonResponse {
  people: Person[];
}

export const useListPerson = () => {
  const query = useQuery({
    queryKey: ['people'],
    queryFn: () => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/people';
      return axios.get<ListPersonResponse>(path);
    },
  })

  return query;
}

export const useGetPerson = (id: string) => {
  const query = useQuery({
    queryKey: ['people', id],
    queryFn: () => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/people/' + id;
      return axios.get<Person>(path);
    },
  })

  return query;
}

export const useCreatePerson = () => {
  type CreatePersonRequest = {
    name: string;
    content: string;
  };

  const auth = useContext(AuthContext);

  const mutation = useMutation({
    mutationFn: (req: CreatePersonRequest) => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/people';
      return axios.post(path, req, {
        headers: {
          Authorization: 'Bearer ' + auth.token,
        }
      });
    },
  })

  return mutation;
}

export const useDeletePerson = () => {
  const auth = useContext(AuthContext);
  const listPersonQuery = useListPerson();

  const mutation = useMutation({
    mutationFn: (id: string) => {
      const path = process.env.NEXT_PUBLIC_BACKEND_URL_PREFIX + '/api/v1/people/' + id;
      return axios.delete(path, {
        headers: {
          Authorization: 'Bearer ' + auth.token,
        }
      });
    },
    onSuccess: () => {
      listPersonQuery.refetch();
    }
  })

  return mutation;
}