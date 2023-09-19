import React, { useEffect } from 'react';
import Head from 'next/head'
import Nav from '../../components/Nav';
import Editor from '../../components/Editor';
import { useCreateResource } from '../../requests/resources';
import { useRouter } from 'next/router';

export default function Index() {
  const createResource = useCreateResource();
  const router = useRouter();

  useEffect(() => {
    if (createResource.isSuccess) {
      router.push("/resources");
    }
  }, [createResource, router]);

  return (
    <>
      <Head>
        <title>Lab Homepage | News</title>
        <meta name="description" content="The news of the lab." />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Nav />
      <Editor
        submit={(titleContent) => createResource.mutate(titleContent)}
        titlePlaceholder="The title of the resource"
        contentPlaceholder="The content of the resource"
      />
    </>
  )
}
