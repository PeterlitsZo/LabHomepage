import React, { useEffect } from 'react';
import Head from 'next/head'
import { useRouter } from 'next/router';

import Nav from '../../components/Nav';
import Editor from '../../components/Editor';
import { useCreatePaper } from '../../requests/papers';

export default function Index() {
  const createPaper = useCreatePaper();
  const router = useRouter();

  useEffect(() => {
    if (createPaper.isSuccess) {
      router.push("/papers");
    }
  }, [createPaper, router]);

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
        submit={(titleContent) => createPaper.mutate(titleContent)}
        titlePlaceholder="The title of the paper"
        contentPlaceholder="The content of the paper"
      />
    </>
  )
}
