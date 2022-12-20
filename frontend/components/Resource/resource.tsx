import { useContext, useEffect, useState } from "react";
import { PlusCircle } from 'lucide-react';

import { AuthContext } from "../../contexts/auth";
import { Resource as ResourceType, useGetResource, useListResource } from "../../requests/resources"
import Link from "next/link";
import { Router, useRouter } from "next/router";

interface ResourceProps {
  id: string;
}

export const Resource = (props: ResourceProps) => {
  const getResourceQuery = useGetResource(props.id);

  return (
    <main className="max-w-3xl m-auto py-24">
      <div className="border border-slate-300 p-8 rounded">
        {getResourceQuery.isSuccess && (function() {
          const resource = getResourceQuery.data.data;
          return (
            <>
              <h1 className="text-3xl">
                {resource.title}
              </h1>
              {resource.content !== "" && (
                <div className="mt-4 text-lg">
                  {resource.content}
                </div>
              )}
            </>
          )
        })()}
      </div>
    </main>
  )
}