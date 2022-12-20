import { useContext, useEffect } from "react";
import { PlusCircle, Trash2 } from 'lucide-react';

import { AuthContext } from "../../contexts/auth";
import { useDeletePaper, useListPaper } from "../../requests/papers"
import Link from "next/link";
import { Router, useRouter } from "next/router";

export const PaperList = () => {
  const auth = useContext(AuthContext);
  const router = useRouter();

  const listPaperQuery = useListPaper();
  const deletePaperMutation = useDeletePaper();

  useEffect(() => {
    console.log('list paper', listPaperQuery.data?.data);
  }, [listPaperQuery]);

  return (
    <main className="max-w-3xl m-auto py-24">
      <ul className="text-xl p-4 rounded border border-slate-300">
        <>
          {auth.isLoggedIn && (
            <Link href="/papers/create" className="flex content-center items-center gap-2 p-4 h-14 rounded text-slate-500 hover:text-sky-900 cursor-pointer hover:bg-slate-100">
              <PlusCircle />
              Create New Paper
            </Link>
          )}
          {listPaperQuery.isSuccess && (() => {
            const paper = listPaperQuery.data.data.papers;
            return [...paper].reverse().map(paper => (
              <li
                key={paper.id}
                onClick={() => {
                  router.push("/papers/" + paper.id)
                }}
                className="flex content-center items-center p-4 h-14 rounded hover:bg-slate-100 cursor-pointer"
              >
                <span>
                  {paper.title}
                </span>
                <span className="flex-1" />
                {auth.isLoggedIn && (
                  <Trash2
                    onClick={(e: MouseEvent) => {
                      e.stopPropagation();
                      deletePaperMutation.mutate(paper.id);
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