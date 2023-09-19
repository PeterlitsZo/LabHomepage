import React, { useEffect } from 'react';
import Head from 'next/head'
import { useRouter } from "next/router";

import Nav from '../../components/Nav';
import Editor from "../../components/Editor";

import { useCreateNews } from "../../requests/news";

export default function Index() {
  const createNews = useCreateNews();
  const router = useRouter();

  useEffect(() => {
    if (createNews.isSuccess) {
      router.push("/news");
    }
  }, [createNews, router]);

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
        submit={(titleContent) => createNews.mutate(titleContent)}
        titlePlaceholder="The title of the news"
        contentPlaceholder="The content of the news"
      />
    </>
  )
}
