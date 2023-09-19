import Image from 'next/image';
import { Router, useRouter } from 'next/router';
import React from 'react';
import { useListNews } from '../../requests/news';
import { useListPaper } from '../../requests/papers';
import { useListPerson } from '../../requests/people';
import styles from './home.module.css';

export const Home = () => {
  const listNewsQuery = useListNews();
  const listPaperQuery = useListPaper();
  const listPersonQuery = useListPerson();

  return (
    <div className="bg-slate-100 min-h-screen p-4">
      <div className="relative max-w-7xl m-auto">
        <h1 className="pt-60 text-7xl leading-none tracking-wide">
          Lab Homepage
        </h1>
        <p className="pt-2 text-xl text-slate-500">
          To be or not to be.
        </p>
        <div className="z-0 absolute -right-12 -top-12 w-200 h-200">
          <Image
            src="/homepage-pic.png"
            alt="Homepage Picture"
            className="object-contain"
            fill
          />
        </div>
      </div>
      <div className="relative z-10 mt-10 bg-white max-w-7xl m-auto p-8 pt-2 rounded">
        <Part
          title="Rencent News"
          dataFetchSuccess={listNewsQuery.isSuccess}
          dataRender={(render) => {
            let news = listNewsQuery.data!.data.news;
            return [...news].reverse().slice(0, 3).map(
              news => render(news.id, '/news/' + news.id, (
                <>
                  <h2 className="text-lg mb-2">{news.title}</h2>
                  <p className={`text-slate-500 ${styles.pInCard}`}>
                    {news.content}
                  </p>
                </>
              ))
            );
          }}
        />
        <Part
          title="Rencent Papers"
          dataFetchSuccess={listPaperQuery.isSuccess}
          dataRender={(render) => {
            let papers = listPaperQuery.data!.data.papers;
            return [...papers].reverse().slice(0, 3).map(
              paper => render(paper.id, '/papers/' + paper.id, (
                <>
                  <h2 className="text-lg mb-2">{paper.title}</h2>
                  <p className={`text-slate-500 ${styles.pInCard}`}>
                    {paper.content}
                  </p>
                </>
              ))
            );
          }}
        />
        <Part
          title="People in our Lab"
          dataFetchSuccess={listPersonQuery.isSuccess}
          dataRender={(render) => {
            let people = listPersonQuery.data!.data.people;
            return people.map(
              person => render(person.id, '/people/' + person.id, (
                <>
                  <h2 className="text-lg mb-2">{person.name}</h2>
                  <p className={`text-slate-500 ${styles.pInCard}`}>
                    {person.content}
                  </p>
                </>
              ))
            );
          }}
        />
      </div>
    </div>
  )
};

interface PartProps {
  title: string;
  dataFetchSuccess: boolean;
  dataRender: (render: (key: string, href: string, children: React.ReactNode) => React.ReactNode) => React.ReactNode;
}

const Part = (props: PartProps) => {
  const router = useRouter();

  return (
    <>
      <h2 className="text-2xl font-bold font-semibold mt-4 pt-4 mb-4">{props.title}</h2>
      <div className="grid grid-cols-3 gap-4">
        {props.dataFetchSuccess && props.dataRender((key: string, href: string, children: React.ReactNode) => (
          <div
            key={key}
            className="p-4 border border-slate-300 hover:border-slate-500 cursor-pointer rounded"
            onClick={() => {router.push(href)}}
          >
            {children}
          </div>
        ))}
      </div>
    </>
  );
}