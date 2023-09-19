import { useEffect, useState } from "react";

import { useGetPaper } from "../../requests/papers"
import MarkdownCard from "../MarkdownCard";

interface PaperProps {
  id: string;
}

export const Paper = (props: PaperProps) => {
  const getPaperQuery = useGetPaper(props.id);

  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  useEffect(() => {
    if (getPaperQuery.isSuccess) {
      const news = getPaperQuery.data.data;
      setTitle(news.title);
      setContent(news.content);
    }
  }, [getPaperQuery.isSuccess]);

  return (
    <MarkdownCard
      title={title}
      content={content}
    />
  )
}