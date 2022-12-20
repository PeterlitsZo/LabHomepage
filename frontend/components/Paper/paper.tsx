import { useContext, useEffect, useState } from "react";
import { PlusCircle } from 'lucide-react';

import { AuthContext } from "../../contexts/auth";
import { Paper as PaperType, useGetPaper, useListPaper } from "../../requests/papers"
import Link from "next/link";
import { Router, useRouter } from "next/router";

interface PaperProps {
  id: string;
}

export const Paper = (props: PaperProps) => {
  const getPaperQuery = useGetPaper(props.id);

  return (
    <main className="max-w-3xl m-auto py-24">
      <div className="border border-slate-300 p-8 rounded">
        {getPaperQuery.isSuccess && (function() {
          const paper = getPaperQuery.data.data;
          return (
            <>
              <h1 className="text-3xl">
                {paper.title}
              </h1>
              {paper.content !== "" && (
                <div className="mt-4 text-lg">
                  {paper.content}
                </div>
              )}
            </>
          )
        })()}
      </div>
    </main>
  )
}