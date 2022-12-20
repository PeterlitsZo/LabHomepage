import { useContext, useEffect } from "react";
import { PlusCircle, Trash2 } from 'lucide-react';

import { AuthContext } from "../../contexts/auth";
import { useDeletePerson, useListPerson } from "../../requests/people"
import Link from "next/link";
import { Router, useRouter } from "next/router";

export const PersonList = () => {
  const auth = useContext(AuthContext);
  const router = useRouter();

  const listPersonQuery = useListPerson();
  const deletePersonMutation = useDeletePerson();

  useEffect(() => {
    console.log('list person', listPersonQuery.data?.data);
  }, [listPersonQuery]);

  return (
    <main className="max-w-3xl m-auto py-24">
      <ul className="text-xl p-4 rounded border border-slate-300">
        <>
          {auth.isLoggedIn && (
            <Link href="/people/create" className="flex content-center items-center gap-2 p-4 h-14 rounded text-slate-500 hover:text-sky-900 cursor-pointer hover:bg-slate-100">
              <PlusCircle />
              Create New Person
            </Link>
          )}
          {listPersonQuery.isSuccess && (() => {
            const person = listPersonQuery.data.data.people;
            return [...person].reverse().map(person => (
              <li
                key={person.id}
                onClick={() => {
                  router.push("/people/" + person.id)
                }}
                className="flex content-center items-center p-4 h-14 rounded hover:bg-slate-100 cursor-pointer"
              >
                <span>
                  {person.name}
                </span>
                <span className="flex-1" />
                {auth.isLoggedIn && (
                  <Trash2
                    onClick={(e: MouseEvent) => {
                      e.stopPropagation();
                      deletePersonMutation.mutate(person.id);
                    }}
                    className="hover:bg-slate-200 p-2 h-10 w-10 rounded-full text-slate-500"
                  />
                )}
              </li>
            ))
          })()}
        </>
      </ul>
    </main>
  )
}