import React from 'react';
import Head from 'next/head'
import Nav from '../../components/Nav';
import PersonList from '../../components/PersonList';

export default function Index() {
  return (
    <>
      <Head>
        <title>Lab Homepage | News</title>
        <meta name="description" content="The news of the lab." />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Nav />
      <PersonList />
    </>
  )
}
