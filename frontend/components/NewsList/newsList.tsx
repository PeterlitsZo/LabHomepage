import { useContext, useEffect } from "react";
import { PlusCircle, Trash2 } from 'lucide-react';

import { AuthContext } from "../../contexts/auth";
import { useDeleteNews, useListNews } from "../../requests/news"
import Link from "next/link";
import { Router, useRouter } from "next/router";

export const NewsList = () => {
  const auth = useContext(AuthContext);
  const router = useRouter();

  const listNewsQuery = useListNews();
  const deleteNewsMutation = useDeleteNews();

  useEffect(() => {
    console.log('list news', listNewsQuery.data?.data);
  }, [listNewsQuery]);

  return (
    <main className="max-w-3xl m-auto py-24">
      <ul className="text-xl p-4 rounded border border-slate-300">
        <>
          {auth.isLoggedIn && (
            <Link href="/news/create" className="flex content-center items-center gap-2 p-4 h-14 rounded text-slate-500 hover:text-sky-900 cursor-pointer hover:bg-slate-100">
              <PlusCircle />
              Create New News
            </Link>
          )}
          {listNewsQuery.isSuccess && (() => {
            const news = listNewsQuery.data.data.news;
            return [...news].reverse().map(news => (
              <li
                key={news.id}
                onClick={() => {
                  router.push("/news/" + news.id)
                }}
                className="flex content-center items-center p-4 h-14 rounded hover:bg-slate-100 cursor-pointer"
              >
                <span>
                  {news.title}
                </span>
                <span className="flex-1" />
                {auth.isLoggedIn && (
                  <Trash2
                    onClick={(e: MouseEvent) => {
                      e.stopPropagation();
                      deleteNewsMutation.mutate(news.id);
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