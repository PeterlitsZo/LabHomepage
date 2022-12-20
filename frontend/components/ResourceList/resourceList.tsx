import { useContext, useEffect } from "react";
import { PlusCircle, Trash2 } from 'lucide-react';

import { AuthContext } from "../../contexts/auth";
import { useDeleteResource, useListResource } from "../../requests/resources"
import Link from "next/link";
import { Router, useRouter } from "next/router";

export const ResourceList = () => {
  const auth = useContext(AuthContext);
  const router = useRouter();

  const listResourceQuery = useListResource();
  const deleteResourceMutation = useDeleteResource();

  console.log('list resource', listResourceQuery.isSuccess, listResourceQuery.data?.data);

  return (
    <main className="max-w-3xl m-auto py-24">
      <ul className="text-xl p-4 rounded border border-slate-300">
        <>
          {auth.isLoggedIn && (
            <Link href="/resources/create" className="flex content-center items-center gap-2 p-4 h-14 rounded text-slate-500 hover:text-sky-900 cursor-pointer hover:bg-slate-100">
              <PlusCircle />
              Create New Resource
            </Link>
          )}
          {listResourceQuery.isSuccess && (() => {
            const resource = listResourceQuery.data.data.resources;
            return [...resource].reverse().map(resource => (
              <li
                key={resource.id}
                onClick={() => {
                  router.push("/resources/" + resource.id)
                }}
                className="flex content-center items-center p-4 h-14 rounded hover:bg-slate-100 cursor-pointer"
              >
                <span>
                  {resource.title}
                </span>
                <span className="flex-1" />
                {auth.isLoggedIn && (
                  <Trash2
                    onClick={(e: MouseEvent) => {
                      e.stopPropagation();
                      deleteResourceMutation.mutate(resource.id);
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