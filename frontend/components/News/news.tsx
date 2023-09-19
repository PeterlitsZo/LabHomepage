import { useEffect, useState } from 'react';

import { useGetNews } from "../../requests/news"

import MarkdownCard from "../MarkdownCard";

interface NewsProps {
  id: string;
}

export const News = (props: NewsProps) => {
  const getNewsQuery = useGetNews(props.id);

  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  useEffect(() => {
    if (getNewsQuery.isSuccess) {
      const news = getNewsQuery.data.data;
      setTitle(news.title);
      setContent(news.content);
    }
  }, [getNewsQuery.isSuccess]);

  return (
    <MarkdownCard
      title={title}
      content={content}
    />
  )
}