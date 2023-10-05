import { useContext, useEffect } from "react";
import { PlusCircle, Trash2 } from 'lucide-react';

import { useDeletePerson, useListPerson } from "../../requests/people"
import Link from "next/link";
import { Router, useRouter } from "next/router";
import List from "../List";

export const PersonList = () => {
  const listPersonQuery = useListPerson();
  const deletePersonMutation = useDeletePerson();

  return (
    <List
      createItem={{
        url: '/people/create',
        label: 'Create New Person',
      }}
      items={{
        whenRender: listPersonQuery.isSuccess,
        refPrefix: '/people/',
        deleteItem(id) {
          deletePersonMutation.mutate(id);
        },
        getTitle: d => d.name,
        data: listPersonQuery.data?.data.people,
      }}
    />
  )
}