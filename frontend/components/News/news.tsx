import { useContext, useEffect, useState } from "react";
import { PlusCircle } from 'lucide-react';

import { AuthContext } from "../../contexts/auth";
import { News as NewsType, useGetNews, useListNews } from "../../requests/news"
import Link from "next/link";
import { Router, useRouter } from "next/router";

interface NewsProps {
  id: string;
}

export const News = (props: NewsProps) => {
  const getNewsQuery = useGetNews(props.id);

  return (
    <main className="max-w-3xl m-auto py-24">
      <div className="border border-slate-300 p-8 rounded">
        {getNewsQuery.isSuccess && (function() {
          const news = getNewsQuery.data.data;
          return (
            <>
              <h1 className="text-3xl">
                {news.title}
              </h1>
              {news.content !== "" && (
                <div className="mt-4 text-lg">
                  {news.content}
                </div>
              )}
            </>
          )
        })()}
      </div>
    </main>
  )
}