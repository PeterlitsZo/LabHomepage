import React from 'react';
import Head from 'next/head'
import Home from '../components/Home'
import Nav from '../components/Nav';

export default function Index() {
  return (
    <>
      <Head>
        <title>Lab Homepage</title>
        <meta name="description" content="The homepage of the lab." />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Nav />
      <Home />
    </>
  )
}
