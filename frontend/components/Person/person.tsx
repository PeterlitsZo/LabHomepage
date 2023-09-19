import { useEffect, useState } from "react";
import { useGetPerson } from "../../requests/people"
import MarkdownCard from "../MarkdownCard";

interface PersonProps {
  id: string;
}

export const Person = (props: PersonProps) => {
  const getPersonQuery = useGetPerson(props.id);

  const [name, setName] = useState("");
  const [content, setContent] = useState("");

  useEffect(() => {
    if (getPersonQuery.isSuccess) {
      const person = getPersonQuery.data.data;
      setName(person.name);
      setContent(person.content);
    }
  }, [getPersonQuery.isSuccess]);

  return (
    <MarkdownCard
      title={name}
      content={content}
    />
  )
}