import { useEffect, useState } from "react";

import { useGetResource } from "../../requests/resources"
import MarkdownCard from "../MarkdownCard";

interface ResourceProps {
  id: string;
}

export const Resource = (props: ResourceProps) => {
  const getResourceQuery = useGetResource(props.id);

  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  useEffect(() => {
    if (getResourceQuery.isSuccess) {
      const resource = getResourceQuery.data.data;
      setTitle(resource.title);
      setContent(resource.content);
    }
  }, [getResourceQuery.isSuccess]);

  return (
    <MarkdownCard
      title={title}
      content={content}
    />
  )
}