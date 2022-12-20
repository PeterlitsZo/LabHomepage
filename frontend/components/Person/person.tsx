import { useContext, useEffect, useState } from "react";
import { PlusCircle } from 'lucide-react';

import { AuthContext } from "../../contexts/auth";
import { Person as PersonType, useGetPerson, useListPerson } from "../../requests/people"
import Link from "next/link";
import { Router, useRouter } from "next/router";

interface PersonProps {
  id: string;
}

export const Person = (props: PersonProps) => {
  const getPersonQuery = useGetPerson(props.id);

  return (
    <main className="max-w-3xl m-auto py-24">
      <div className="border border-slate-300 p-8 rounded">
        {getPersonQuery.isSuccess && (function() {
          const person = getPersonQuery.data.data;
          return (
            <>
              <h1 className="text-3xl">
                {person.name}
              </h1>
              {person.content !== "" && (
                <div className="mt-4 text-lg">
                  {person.content}
                </div>
              )}
            </>
          )
        })()}
      </div>
    </main>
  )
}