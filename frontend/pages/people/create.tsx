import React, { useEffect } from 'react';
import Head from 'next/head'
import Nav from '../../components/Nav';
import Editor from '../../components/Editor';
import { useCreatePerson } from '../../requests/people';
import { useRouter } from 'next/router';

export default function Index() {
  const createPerson = useCreatePerson();
  const router = useRouter();

  useEffect(() => {
    if (createPerson.isSuccess) {
      router.push("/people");
    }
  }, [createPerson, router]);

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
        submit={(titleContent) => {
          createPerson.mutate({ name: titleContent.title, content: titleContent.content })
        }}
        titlePlaceholder="The name of the person"
        contentPlaceholder="The details of the person"
      />
    </>
  )
}
