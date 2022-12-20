import React from 'react';
import Head from 'next/head'
import Nav from '../../components/Nav';
import { useRouter } from 'next/router';
import Person from '../../components/Person';

export default function Index() {
  const router = useRouter();
  const { newsId } = router.query;

  return (
    <>
      <Head>
        <title>Lab Homepage | News</title>
        <meta name="description" content="The news of the lab." />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Nav />
      <Person id={newsId as string} />
    </>
  )
}
